package services

import (
	"fmt"

	"github.com/arthurshafikov/tg-gladiator/internal/core/constants"
	"github.com/arthurshafikov/tg-gladiator/internal/core/constants/enums"
	"github.com/arthurshafikov/tg-gladiator/internal/core/errors"
	"github.com/arthurshafikov/tg-gladiator/internal/core/models"
	"github.com/arthurshafikov/tg-gladiator/internal/core/types"
	"github.com/arthurshafikov/tg-gladiator/internal/repository"
)

const (
	StartEnergy = 10
	StartGold   = 0

	LevelUpStatsUpgradeAvailable = 5
)

type HeroService struct {
	logger Logger
	repo   repository.Hero

	validatorService *ValidatorService
	heroItemsService *HeroItemService

	eventsHandler EventsHandler
}

func newHeroService(
	logger Logger,
	repo repository.Hero,
	validatorService *ValidatorService,
	heroItemsService *HeroItemService,
	eventsHandler EventsHandler,
) *HeroService {
	return &HeroService{
		logger:           logger,
		repo:             repo,
		validatorService: validatorService,
		heroItemsService: heroItemsService,
		eventsHandler:    eventsHandler,
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

		AttackBonus:                0,
		DefenseBonus:               0,
		HpBonus:                    0,
		CriticalChancePercentBonus: 0,
		EvasionChancePercentBonus:  0,
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

func (s *HeroService) RewardXP(ctx *types.Context, id int64, amount int) error {
	hero, err := s.FindMy(ctx, id)
	if err != nil {
		return err
	}

	fields := map[string]interface{}{
		models.HeroFieldXP: hero.XP + amount,
	}

	var newLevel = false
	if hero.Level < hero.CalculateLevelForXP(hero.XP+amount) {
		hero.Level = hero.Level + 1
		fields[models.HeroFieldLevel] = hero.Level

		fields[models.HeroFieldLevelUpBonusesLeft] = hero.LevelUpBonusesLeft + LevelUpStatsUpgradeAvailable

		newLevel = true
	}

	if _, err := s.repo.UpdateMap(ctx.GetContext(), id, fields); err != nil {
		return err
	}

	if newLevel {
		s.eventsHandler.Dispatch(constants.EventHeroLevelUp, *hero)
	}

	return nil
}

func (s *HeroService) SpendLevelUpBonus(ctx *types.Context, id int64, statName string) (*models.Hero, error) {
	hero, err := s.FindMy(ctx, id)
	if err != nil {
		return nil, err
	}

	if hero.LevelUpBonusesLeft < 1 {
		return nil, errors.ErrForbidden // @todo clearer error?
	}

	fields := make(map[string]any, 1)

	switch statName {
	case enums.StatsUpgradeHealth:
		fields[models.HeroFieldHealthBonus] = hero.HpBonus + 10 // @todo all these values from config
	case enums.StatsUpgradeAttack:
		fields[models.HeroFieldAttackBonus] = hero.AttackBonus + 1
	case enums.StatsUpgradeDefence:
		fields[models.HeroFieldDefenseBonus] = hero.DefenseBonus + 1
	case enums.StatsUpgradeCriticalChance:
		fields[models.HeroFieldCriticalChancePercentBonus] = hero.CriticalChancePercentBonus + 1
	case enums.StatsUpgradeEvasionChance:
		fields[models.HeroFieldEvasionChancePercentBonus] = hero.EvasionChancePercentBonus + 1
	default:
		s.logger.Error(fmt.Errorf("statName unknown error: %s", statName))

		return nil, errors.ErrServerError
	}

	fields[models.HeroFieldLevelUpBonusesLeft] = hero.LevelUpBonusesLeft - 1

	hero, err = s.repo.UpdateMap(ctx.GetContext(), id, fields)
	if err != nil {
		s.logger.Error(err)

		return nil, errors.ErrServerError
	}

	return hero, nil
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
