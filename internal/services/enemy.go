package services

import (
	"github.com/arthurshafikov/tg-gladiator/internal/core/errors"
	"github.com/arthurshafikov/tg-gladiator/internal/core/models"
	"github.com/arthurshafikov/tg-gladiator/internal/core/types"
	"github.com/arthurshafikov/tg-gladiator/internal/repository"
)

type EnemyService struct {
	logger Logger
	repo   repository.Enemy
}

func newEnemyService(logger Logger, repo repository.Enemy) *EnemyService {
	return &EnemyService{
		logger: logger,
		repo:   repo,
	}
}

// @todo randomize the enemy
func (s *EnemyService) getRandomEnemy(ctx *types.Context) (*models.Enemy, error) {
	enemies, err := s.repo.GetBy(ctx.GetContext(), &models.Enemy{})
	if err != nil {
		s.logger.Error(err)

		return nil, errors.ErrServerError
	}

	return &(*enemies)[0], nil
}
