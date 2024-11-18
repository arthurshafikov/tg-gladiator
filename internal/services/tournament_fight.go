package services

import (
	"crypto/rand"
	"math/big"
	"time"

	"github.com/arthurshafikov/tg-gladiator/internal/core/constants/enums"
	"github.com/arthurshafikov/tg-gladiator/internal/core/errors"
	"github.com/arthurshafikov/tg-gladiator/internal/core/models"
	"github.com/arthurshafikov/tg-gladiator/internal/core/types"
	"github.com/arthurshafikov/tg-gladiator/internal/repository"
)

type TournamentFightService struct {
	logger       Logger
	repo         repository.Fight
	heroService  *HeroService
	enemyService *EnemyService
}

func newTournamentFightService(logger Logger, repo repository.Fight, heroService *HeroService, enemyService *EnemyService) *TournamentFightService {
	return &TournamentFightService{
		logger:       logger,
		repo:         repo,
		heroService:  heroService,
		enemyService: enemyService,
	}
}

func (s *TournamentFightService) Create(ctx *types.Context, heroID int64) (*models.Fight, error) {
	enemy, err := s.enemyService.getRandomEnemy(ctx)
	if err != nil {
		return nil, err
	}

	hero, err := s.heroService.FindMy(ctx, heroID)
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

	fight, err = s.repo.Create(ctx.GetContext(), *fight)
	if err != nil {
		if errors.Is(err, errors.ErrAlreadyExists) {
			return nil, errors.ErrAlreadyExists
		}

		s.logger.Error(err)

		return nil, err
	}

	return fight, nil
}

func (s *TournamentFightService) FindActiveByHeroID(ctx *types.Context, heroID int64) (*models.Fight, error) {
	// @todo check access
	fight, err := s.repo.FindBy(ctx.GetContext(), &models.Fight{
		HeroID: heroID,
		Status: enums.FightStatusActive,
	})
	if err != nil {
		if errors.Is(err, errors.ErrNotFound) {
			return nil, errors.ErrNotFound
		}

		s.logger.Error(err)

		return nil, errors.ErrServerError
	}

	return fight, nil
}

func (s *TournamentFightService) MakeTurn(
	ctx *types.Context,
	fight *models.Fight,
	actionType enums.FightActionType, // @todo temp unused
) (*models.Fight, *models.FightEvents, error) {
	fields := map[string]interface{}{}

	events := &models.FightEvents{}

	var err error
	events.HeroFightEvent, err = s.getFighterFightEvent(fight.Hero, fight.Opponent)
	if err != nil {
		return nil, nil, err
	}

	events.OpponentFightEvent, err = s.getFighterFightEvent(fight.Opponent, fight.Hero)
	if err != nil {
		return nil, nil, err
	}

	if events.HeroFightEvent.DamageReceived > 0 {
		opponentHP := fight.OpponentHP - events.HeroFightEvent.DamageReceived
		if opponentHP < 0 {
			opponentHP = 0
		}

		fields[models.FightFieldOpponentHP] = opponentHP
	}

	if events.OpponentFightEvent.DamageDealt > 0 {
		heroHP := fight.HeroHP - events.OpponentFightEvent.DamageReceived
		if heroHP < 0 {
			heroHP = 0
		}

		fields[models.FightFieldHeroHP] = heroHP
	}

	if fields[models.FightFieldOpponentHP] == 0 || fields[models.FightFieldHeroHP] == 0 {
		fields[models.FightFieldStatus] = enums.FightStatusEnded
	}

	fight, err = s.repo.UpdateMap(ctx.GetContext(), fight.ID, fields)
	if err != nil {
		if errors.Is(err, errors.ErrNotFound) {
			return nil, nil, errors.ErrNotFound
		}

		s.logger.Error(err)

		return nil, nil, errors.ErrServerError
	}

	return fight, events, nil
}

func (s *TournamentFightService) RunAwayAsHero(ctx *types.Context, fightID int64) error {
	// @todo check access
	if _, err := s.repo.Update(ctx.GetContext(), fightID, &models.Fight{
		Status: enums.FightStatusFleed,
	}); err != nil {
		s.logger.Error(err)

		return errors.ErrServerError
	}

	return nil
}

func (s *TournamentFightService) getFighterFightEvent(
	fighter models.Fighter,
	opponent models.Fighter,
) (models.FightEvent, error) {
	fighterEvent := models.FightEvent{}

	isCritical, err := s.calculateRandomChance(fighter.GetCriticalChancePercent())
	if err != nil {
		return fighterEvent, err
	}
	fighterEvent.IsCritical = isCritical

	// various damage
	if fighter.GetMaxAttack() != fighter.GetMinAttack() {
		damageDealt, err := rand.Int(
			rand.Reader,
			big.NewInt(
				int64(fighter.GetMaxAttack()-fighter.GetMinAttack()),
			),
		)
		if err != nil {
			s.logger.Error(err)

			return fighterEvent, errors.ErrServerError
		}
		damageDealtInt := int(damageDealt.Int64())
		damageDealtInt += 1 + fighter.GetMinAttack()
		fighterEvent.DamageDealt = damageDealtInt
	} else {
		fighterEvent.DamageDealt = fighter.GetMinAttack()
	}

	if fighterEvent.IsCritical {
		fighterEvent.DamageDealt *= 2
	}

	wasEvaded, err := s.calculateRandomChance(opponent.GetEvasionChancePercent())
	if err != nil {
		return fighterEvent, err
	}
	fighterEvent.WasEvaded = wasEvaded

	if fighterEvent.WasEvaded {
		fighterEvent.DamageReceived = 0
	} else {
		fighterEvent.DamageReceived = fighterEvent.DamageDealt // @todo armor
	}

	return fighterEvent, nil
}

func (s *TournamentFightService) calculateRandomChance(desiredChance int) (bool, error) {
	random, err := rand.Int(rand.Reader, big.NewInt(100))
	if err != nil {
		s.logger.Error(err)

		return false, errors.ErrServerError
	}

	return (random.Int64() + 1) <= int64(desiredChance), nil
}
