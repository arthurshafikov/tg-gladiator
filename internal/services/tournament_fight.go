package services

import (
	"fmt"
	"time"

	"github.com/arthurshafikov/tg-gladiator/internal/core/constants/enums"
	"github.com/arthurshafikov/tg-gladiator/internal/core/errors"
	"github.com/arthurshafikov/tg-gladiator/internal/core/helpers"
	"github.com/arthurshafikov/tg-gladiator/internal/core/models"
	"github.com/arthurshafikov/tg-gladiator/internal/core/types"
)

type TournamentFightService struct {
	logger                       Logger
	fightService                 *FightService
	heroService                  *HeroService
	enemyService                 *EnemyService
	tournamentFightEventsService *TournamentFightEventsService
}

func newTournamentFightService(
	logger Logger,
	fightService *FightService,
	heroService *HeroService,
	enemyService *EnemyService,
	tournamentFightEventsService *TournamentFightEventsService,
) *TournamentFightService {
	return &TournamentFightService{
		logger:                       logger,
		fightService:                 fightService,
		heroService:                  heroService,
		enemyService:                 enemyService,
		tournamentFightEventsService: tournamentFightEventsService,
	}
}

func (s *TournamentFightService) Create(ctx *types.Context, heroID int64) (*models.Fight, error) {
	hero, err := s.heroService.FindMy(ctx, heroID)
	if err != nil {
		return nil, err
	}

	enemy, err := s.enemyService.getRandomEnemy(ctx, hero.Level)
	if err != nil {
		return nil, err
	}

	fight := &models.Fight{
		Status:       enums.FightStatusActive,
		HeroID:       hero.ID,
		OpponentType: enums.OpponentTypeMob,
		OpponentID:   enemy.ID,
		HeroHP:       hero.CurrentHP,
		OpponentHP:   enemy.HP,
		CreatedAt:    time.Now(),
	}

	fight, err = s.fightService.create(ctx, *fight)
	if err != nil {
		return nil, err
	}

	return fight, nil
}

func (s *TournamentFightService) RunAwayAsHero(ctx *types.Context, fightID int64) error {
	// @todo check access
	// @todo energy - 1
	if _, err := s.fightService.update(ctx, fightID, &models.Fight{
		Status: enums.FightStatusFleed,
	}); err != nil {
		s.logger.Error(err)

		return errors.ErrServerError
	}

	return nil
}

func (s *TournamentFightService) MakeTurn(
	ctx *types.Context,
	fight *models.Fight,
	actionType enums.FightActionType,
) (*models.Fight, error) {
	updateFields := map[string]interface{}{}

	events := &models.FightEvents{}

	var err error
	events.HeroFightEvent, err = s.tournamentFightEventsService.getFightEventFor(
		actionType,
		fight.Hero,
		fight.Opponent,
	)
	if err != nil {
		return nil, err
	}

	events.OpponentFightEvent, err = s.tournamentFightEventsService.getRandomFightEventFor(
		fight.Opponent,
		fight.Hero,
	)
	if err != nil {
		return nil, err
	}

	if events.HeroFightEvent.DamageReceived > 0 {
		opponentHP := fight.OpponentHP - events.HeroFightEvent.DamageReceived
		if opponentHP < 0 {
			opponentHP = 0
		}

		updateFields[models.FightFieldOpponentHP] = opponentHP
	}

	if events.OpponentFightEvent.DamageDealt > 0 {
		heroHP := fight.HeroHP - events.OpponentFightEvent.DamageReceived
		if heroHP < 0 {
			heroHP = 0
		}

		updateFields[models.FightFieldHeroHP] = heroHP
	}

	if updateFields[models.FightFieldOpponentHP] == 0 || updateFields[models.FightFieldHeroHP] == 0 {
		updateFields[models.FightFieldStatus] = enums.FightStatusEnded

		if updateFields[models.FightFieldOpponentHP] == 0 {
			goldReward, err := s.calculateGoldReward(fight.Opponent)
			if err != nil {
				return nil, err
			}

			updateFields[models.FightFieldGoldReward] = goldReward

			if err := s.heroService.RewardGold(ctx, fight.HeroID, goldReward); err != nil {
				return nil, err
			}

			xpReward, err := s.calculateXPReward(fight.Opponent)
			if err != nil {
				return nil, err
			}

			updateFields[models.FightFieldXPReward] = xpReward

			// @todo show msg about new level
			if err := s.heroService.RewardXP(ctx, fight.HeroID, xpReward); err != nil {
				return nil, err
			}
		}
	}

	fight, err = s.fightService.updateMap(ctx, fight.ID, updateFields)
	if err != nil {
		return nil, err
	}

	fight.Events = events

	return fight, nil
}

func (s *TournamentFightService) calculateGoldReward(opponent models.Fighter) (int, error) {
	if opponent.GetMaxGoldReward() < opponent.GetMinGoldReward() {
		s.logger.Error(fmt.Errorf(
			"gold reward max cannot be less than min, enemy id: %s - %v",
			opponent.GetType(),
			opponent.GetID(),
		))

		return 0, errors.ErrServerError
	}

	return s.generateRandomNumberInRange(opponent.GetMinGoldReward(), opponent.GetMaxGoldReward())
}

func (s *TournamentFightService) calculateXPReward(opponent models.Fighter) (int, error) {
	if opponent.GetMaxXPReward() < opponent.GetMinXPReward() {
		s.logger.Error(fmt.Errorf(
			"xp reward max cannot be less than min, enemy id: %s - %v",
			opponent.GetType(),
			opponent.GetID(),
		))

		return 0, errors.ErrServerError
	}

	return s.generateRandomNumberInRange(opponent.GetMinXPReward(), opponent.GetMaxXPReward())
}

func (s *TournamentFightService) generateRandomNumberInRange(min, max int) (int, error) {
	random, err := helpers.GenerateRandomNumberInRange(min, max)
	if err != nil {
		s.logger.Error(err)

		return 0, errors.ErrServerError
	}

	return random, nil
}
