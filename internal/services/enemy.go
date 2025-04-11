package services

import (
	"crypto/rand"
	"math/big"

	"github.com/arthurshafikov/tg-gladiator/internal/core/constants/enums"
	"github.com/arthurshafikov/tg-gladiator/internal/core/errors"
	"github.com/arthurshafikov/tg-gladiator/internal/core/models"
	"github.com/arthurshafikov/tg-gladiator/internal/core/types"
	"github.com/arthurshafikov/tg-gladiator/internal/repository"
	"github.com/sirupsen/logrus"
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

func (s *EnemyService) find(ctx *types.Context, id int64) (*models.Enemy, error) {
	enemy, err := s.repo.Find(ctx.GetContext(), id)
	if err != nil {
		if errors.Is(err, errors.ErrNotFound) {
			return nil, errors.ErrNotFound
		}

		logrus.Error(err)

		return nil, errors.ErrServerError
	}

	return enemy, nil
}

func (s *EnemyService) getRandomNonBossEnemyFromLocation(ctx *types.Context, fightLocation enums.FightLocation) (*models.Enemy, error) {
	minLevel, maxLevel := enums.FightLocationLevelRange(fightLocation)

	enemies, err := s.repo.GetByMap(ctx.GetContext(), &types.WhereConditions{
		Between: map[string]types.BetweenClause{
			models.EnemyFieldLevel: {
				Min:    minLevel,
				Max:    maxLevel,
				Strict: false,
			},
		},
		Where: map[string]interface{}{
			models.EnemyFieldBossType: nil,
		},
	})

	if err != nil {
		s.logger.Error(err)

		return nil, errors.ErrServerError
	}

	if enemies == nil || len(*enemies) < 1 {
		return nil, errors.ErrEmpty
	}

	randomEnemyIndex, err := rand.Int(rand.Reader, big.NewInt(int64(len(*enemies))))
	if err != nil {
		s.logger.Error(err)

		return nil, errors.ErrServerError
	}

	return &(*enemies)[randomEnemyIndex.Int64()], nil
}

func (s *EnemyService) getNextBossEnemy(ctx *types.Context, heroID int64) (*models.Enemy, error) {
	boss, err := s.repo.GetNextBossForHero(ctx.GetContext(), heroID)
	if err != nil {
		if errors.Is(err, errors.ErrNotFound) {
			return nil, err
		}

		s.logger.Error(err)

		return nil, errors.ErrServerError
	}

	return boss, nil
}

func (s *EnemyService) getBossesAmountLeft(ctx *types.Context, heroID int64) (int, error) {
	amount, err := s.repo.GetBossesAmountLeft(ctx.GetContext(), heroID)
	if err != nil {
		if errors.Is(err, errors.ErrNotFound) {
			return 0, err
		}

		s.logger.Error(err)

		return 0, errors.ErrServerError
	}

	return amount, nil
}
