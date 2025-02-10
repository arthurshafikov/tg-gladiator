package services

import (
	"github.com/arthurshafikov/tg-gladiator/internal/core/errors"
	"github.com/arthurshafikov/tg-gladiator/internal/core/models"
	"github.com/arthurshafikov/tg-gladiator/internal/core/types"
	"github.com/arthurshafikov/tg-gladiator/internal/repository"
	"github.com/sirupsen/logrus"
)

const HeroItemsPerPage = 10

type HeroItemService struct {
	repo repository.HeroItem

	heroService *HeroService
}

func newHeroItemService(repo repository.HeroItem, heroService *HeroService) *HeroItemService {
	return &HeroItemService{
		repo:        repo,
		heroService: heroService,
	}
}

func (s *HeroItemService) Create(ctx *types.Context, heroID int64, itemID int64) (*models.HeroItem, error) {
	heroItem, err := s.repo.Create(ctx.GetContext(), models.HeroItem{
		HeroID: heroID,
		ItemID: itemID,
	})
	if err != nil {
		logrus.Error(err)

		return nil, errors.ErrServerError
	}

	return heroItem, nil
}

func (s *HeroItemService) FindMy(ctx *types.Context, heroID, itemID int64) (*models.HeroItem, error) {
	hero, err := s.heroService.FindMy(ctx, heroID)
	if err != nil {
		return nil, err
	}
	if hero.ChatID != ctx.GetChat().ID {
		return nil, errors.ErrForbidden
	} // @todo helper function verifyMyHero?

	heroItem, err := s.repo.FindBy(ctx.GetContext(), &models.HeroItem{
		HeroID: heroID,
		ItemID: itemID,
	})
	if err != nil {
		if errors.Is(err, errors.ErrNotFound) {
			return nil, err
		}

		logrus.Error(err)

		return nil, errors.ErrServerError
	}

	return heroItem, nil
}

func (s *HeroItemService) GetPaginatedForHero(
	ctx *types.Context,
	heroID int64,
	page int,
) (*models.PaginatedHeroItem, error) {
	heroItemsPaginated, err := s.repo.GetBy(ctx.GetContext(), &models.HeroItem{
		HeroID: heroID,
	}, &types.Pagination{
		Page:    page,
		PerPage: HeroItemsPerPage,
	})
	if err != nil {
		logrus.Error(err)

		return nil, err
	}

	return heroItemsPaginated, nil
}

func (s *HeroItemService) Equip(ctx *types.Context, heroID, itemID int64) error {
	heroItem, err := s.FindMy(ctx, heroID, itemID)
	if err != nil {
		return err
	}

	if _, err := s.repo.UpdateMap(ctx.GetContext(), heroItem, map[string]interface{}{
		"is_equipped": true,
	}); err != nil {
		logrus.Error(err)

		return errors.ErrServerError
	}

	return nil
}

func (s *HeroItemService) Unequip(ctx *types.Context, heroID, itemID int64) error {
	heroItem, err := s.FindMy(ctx, heroID, itemID)
	if err != nil {
		return err
	}

	if _, err := s.repo.UpdateMap(ctx.GetContext(), heroItem, map[string]interface{}{
		"is_equipped": false,
	}); err != nil {
		logrus.Error(err)

		return errors.ErrServerError
	}

	return nil
}
