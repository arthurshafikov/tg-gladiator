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
	heroItemsService *HeroItemService
}

func newHeroService(
	logger Logger,
	repo repository.Hero,
	validatorService *ValidatorService,
	heroItemsService *HeroItemService,
) *HeroService {
	return &HeroService{
		logger:           logger,
		repo:             repo,
		validatorService: validatorService,
		heroItemsService: heroItemsService,
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

	classCharacteristics := heroClass.GetCharacteristics()

	hero, err := s.repo.Create(ctx.GetContext(), models.Hero{
		ChatID:        ctx.GetChat().ID,
		Name:          name,
		Class:         heroClass,
		XP:            0,
		Level:         1,
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

	heroEquipment, err := s.heroItemsService.GetHeroEquipment(ctx, hero.ID)
	if err != nil {
		return nil, err
	}
	hero.Equipment = heroEquipment.GetAllItems()

	return hero, nil
}

func (s *HeroService) GetMy(ctx *types.Context) (*[]models.Hero, error) {
	heroes, err := s.repo.GetBy(ctx.GetContext(), &models.Hero{
		ChatID: ctx.GetChat().ID,
	})
	if err != nil {
		if errors.Is(errors.ErrNotFound, err) {
			return nil, errors.ErrNotFound
		}

		s.logger.Error(err)

		return nil, errors.ErrServerError
	}

	return heroes, nil
}

func (s *HeroService) RewardGold(ctx *types.Context, id int64, amount int) error {
	hero, err := s.FindMy(ctx, id)
	if err != nil {
		return err
	}

	fields := map[string]interface{}{
		models.HeroFieldCurrentGold: hero.CurrentGold + amount,
	}

	if _, err := s.repo.UpdateMap(ctx.GetContext(), id, fields); err != nil {
		return err
	}

	return nil
}

func (s *HeroService) decreaseGold(ctx *types.Context, heroID int64, amount int) error {
	hero, err := s.FindMy(ctx, heroID)
	if err != nil {
		return err
	}

	if _, err := s.repo.UpdateMap(ctx.GetContext(), heroID, map[string]interface{}{
		models.HeroFieldCurrentGold: hero.CurrentGold - amount,
	}); err != nil {
		return err
	}

	return nil
}
