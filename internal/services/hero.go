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

	validatorService *ValidatorService
}

func newHeroService(
	logger Logger,
	repo repository.Hero,
	validatorService *ValidatorService,
) *HeroService {
	return &HeroService{
		logger:           logger,
		repo:             repo,
		validatorService: validatorService,
	}
}

func (s *HeroService) Create(ctx *types.Context, name, class string) (*models.Hero, error) {
	name, err := s.validatorService.SanitizeHeroName(name)
	if err != nil {
		return nil, err
	}

	heroClass, err := enums.GetHeroClass(class)
	if err != nil {
		return nil, err
	}

	classCharacteristics, err := heroClass.GetCharacteristics()
	if err != nil {
		s.logger.Error(err)

		return nil, errors.ErrServerError
	}

	hero, err := s.repo.Create(ctx.GetContext(), models.Hero{
		ChatID:        ctx.GetChat().ID,
		Name:          name,
		Class:         heroClass,
		CurrentHP:     classCharacteristics.StartHP,
		CurrentEnergy: StartEnergy,
		CurrentGold:   StartGold,
	})
	if err != nil {
		s.logger.Error(err)

		return nil, errors.ErrServerError
	}

	return hero, nil
}

func (s *HeroService) DeleteMy(ctx *types.Context, id int64) error {
	hero, err := s.FindMy(ctx, id)
	if err != nil {
		return err
	}

	if err := s.repo.DeleteByID(ctx.GetContext(), hero.ID); err != nil {
		if errors.Is(err, errors.ErrNotFound) {
			return errors.ErrNotFound
		}

		s.logger.Error(err)

		return errors.ErrServerError
	}

	return nil
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
