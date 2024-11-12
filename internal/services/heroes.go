package services

import (
	"github.com/arthurshafikov/tg-gladiator/internal/core/errors"
	"github.com/arthurshafikov/tg-gladiator/internal/core/models"
	"github.com/arthurshafikov/tg-gladiator/internal/core/types"
	"github.com/arthurshafikov/tg-gladiator/internal/repository"
)

type HeroesService struct {
	logger Logger
	repo   repository.Hero
}

func newHeroesService(
	logger Logger,
	repo repository.Hero,
) *HeroesService {
	return &HeroesService{
		logger: logger,
		repo:   repo,
	}
}

func (s *HeroesService) GetMy(ctx *types.Context) (*[]models.Hero, error) {
	heroes, err := s.repo.GetBy(ctx.GetContext(), &models.Hero{
		ChatID: ctx.GetChat().ID,
	})
	if err != nil {
		s.logger.Error(err)

		return nil, errors.ErrServerError
	}

	return heroes, nil
}
