package services

import (
	"github.com/arthurshafikov/tg-gladiator/internal/core/constants/enums"
	"github.com/arthurshafikov/tg-gladiator/internal/core/errors"
	"github.com/arthurshafikov/tg-gladiator/internal/core/models"
	"github.com/arthurshafikov/tg-gladiator/internal/core/types"
	"github.com/arthurshafikov/tg-gladiator/internal/repository"
)

const (
	StartEnergy = 10
	StartGold   = 0
)

type HeroService struct {
	logger Logger
	repo   repository.Hero
}

func newHeroService(
	logger Logger,
	repo repository.Hero,
) *HeroService {
	return &HeroService{
		logger: logger,
		repo:   repo,
	}
}

func (s *HeroService) Create(ctx *types.Context, class string) (*models.Hero, error) {
	heroClass, err := enums.GetHeroClass(class)
	if err != nil {
		return nil, err
	}

	startHP, err := heroClass.GetStartHP()
	if err != nil {
		s.logger.Error(err)

		return nil, errors.ErrServerError
	}

	hero, err := s.repo.Create(ctx.GetContext(), models.Hero{
		ChatID:        ctx.GetChat().ID,
		Class:         heroClass,
		CurrentHP:     startHP,
		CurrentEnergy: StartEnergy,
		CurrentGold:   StartGold,
	})
	if err != nil {
		s.logger.Error(err)

		return nil, errors.ErrServerError
	}

	return hero, nil
}

func (s *HeroService) FindMy(ctx *types.Context, id int64) (*models.Hero, error) {
	hero, err := s.repo.Find(ctx.GetContext(), id)
	if err != nil {
		if errors.Is(errors.ErrNotFound, err) {
			return nil, errors.ErrNotFound
		}

		s.logger.Error(err)

		return nil, errors.ErrServerError
	}

	if hero.ChatID != ctx.GetChat().ID {
		return nil, errors.ErrForbidden
	}

	return hero, nil
}
