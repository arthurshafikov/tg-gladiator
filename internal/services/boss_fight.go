package services

import (
	"time"

	"github.com/arthurshafikov/tg-gladiator/internal/core/constants/enums"
	"github.com/arthurshafikov/tg-gladiator/internal/core/models"
	"github.com/arthurshafikov/tg-gladiator/internal/core/types"
)

type BossFightService struct {
	logger                 Logger
	fightService           *FightService
	heroService            *HeroService
	enemyService           *EnemyService
	tournamentFightService *TournamentFightService
}

func newBossFightService(
	logger Logger,
	fightService *FightService,
	heroService *HeroService,
	enemyService *EnemyService,
	tournamentFightService *TournamentFightService,
) *BossFightService {
	return &BossFightService{
		logger:                 logger,
		fightService:           fightService,
		heroService:            heroService,
		enemyService:           enemyService,
		tournamentFightService: tournamentFightService,
	}
}

func (s *BossFightService) Create(ctx *types.Context, heroID int64) (*models.Fight, error) {
	hero, err := s.heroService.FindMy(ctx, heroID)
	if err != nil {
		return nil, err
	}

	if err := s.heroService.DepleteEnergy(ctx, heroID); err != nil {
		return nil, err
	}

	enemy, err := s.enemyService.getNextBossEnemy(ctx, hero.ID)
	if err != nil {
		return nil, err
	}

	fight := &models.Fight{
		Status:       enums.FightStatusActive,
		HeroID:       hero.ID,
		OpponentType: enums.OpponentTypeMob, // OpponentTypeBoss?
		OpponentID:   enemy.ID,
		HeroHP:       hero.GetHP(),
		OpponentHP:   enemy.HP,
		CreatedAt:    time.Now(),
	}

	fight, err = s.fightService.create(ctx, *fight)
	if err != nil {
		return nil, err
	}

	return fight, nil
}

func (s *BossFightService) GetBossesAmountLeft(ctx *types.Context, heroID int64) (int, error) {
	hero, err := s.heroService.FindMy(ctx, heroID)
	if err != nil {
		return 0, err
	}

	return s.enemyService.getBossesAmountLeft(ctx, hero.ID)
}

func (s *BossFightService) RunAwayAsHero(ctx *types.Context, fightID int64) error {
	return s.tournamentFightService.RunAwayAsHero(ctx, fightID)
}

func (s *BossFightService) MakeTurn(
	ctx *types.Context,
	fight *models.Fight,
	actionType enums.FightActionType,
) (*models.Fight, error) {
	return s.tournamentFightService.MakeTurn(ctx, fight, actionType)
}
