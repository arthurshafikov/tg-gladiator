package services

import (
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
