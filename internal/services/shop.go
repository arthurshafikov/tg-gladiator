package services

import (
	"slices"
	"time"

	"github.com/arthurshafikov/tg-gladiator/internal/core/errors"
	"github.com/arthurshafikov/tg-gladiator/internal/core/helpers"
	"github.com/arthurshafikov/tg-gladiator/internal/core/models"
	"github.com/arthurshafikov/tg-gladiator/internal/core/types"
	"github.com/arthurshafikov/tg-gladiator/internal/repository"
	"github.com/sirupsen/logrus"
)

const ShopItemsRefreshFrequency = time.Hour * 6 // 6 hours

type ShopService struct {
	itemsRepo         repository.Item
	heroShopRepo      repository.HeroShop
	heroShopItemsRepo repository.HeroShopItem
}

func newShopService(
	itemsRepo repository.Item,
	heroShopRepo repository.HeroShop,
	heroShopItemsRepo repository.HeroShopItem,
) *ShopService {
	return &ShopService{
		itemsRepo:         itemsRepo,
		heroShopRepo:      heroShopRepo,
		heroShopItemsRepo: heroShopItemsRepo,
	}
}

func (s *ShopService) GetShopItemsFor(ctx *types.Context, heroID int64) ([]models.HeroShopItem, error) {
	heroShop, err := s.heroShopRepo.FindBy(ctx.GetContext(), &models.HeroShop{
		HeroID: heroID,
	})
	if err != nil && !errors.Is(err, errors.ErrNotFound) {
		logrus.Error(err)

		return nil, errors.ErrServerError
	}

	if heroShop == nil {
		heroShop, err = s.heroShopRepo.Create(ctx.GetContext(), models.HeroShop{
			HeroID:    heroID,
			UpdatesAt: time.Time{},
			CreatedAt: time.Now(),
		})
		if err != nil {
			return nil, err
		}
	}

	if heroShop.UpdatesAt.Before(time.Now()) {
		if err := s.updateHeroShopItems(ctx, heroShop.ID); err != nil {
			return nil, err
		}
	}

	shopItems, err := s.heroShopItemsRepo.GetBy(ctx.GetContext(), &models.HeroShopItem{
		HeroShopID: heroShop.ID,
	})
	if err != nil {
		return nil, err
	}

	return shopItems, nil
}

func (s *ShopService) updateHeroShopItems(ctx *types.Context, heroShopID int64) error {
	// @todo transaction
	if err := s.heroShopItemsRepo.DeleteBy(ctx.GetContext(), &models.HeroShopItem{
		HeroShopID: heroShopID,
	}); err != nil {
		return err
	}

	allItems, err := s.itemsRepo.GetBy(ctx.GetContext(), nil)
	if err != nil {
		return err
	}

	alreadyTakenItemIndexes := []int{}
	for i := 0; i < 10; i++ {
		var randomIndex int
		for {
			randomIndex, err = helpers.GenerateRandomNumberInRange(0, len(allItems)-1) // @todo helper func?
			if err != nil {
				return err
			}

			if !slices.Contains(alreadyTakenItemIndexes, randomIndex) {
				break
			}
		}
		alreadyTakenItemIndexes = append(alreadyTakenItemIndexes, randomIndex)

		item := allItems[randomIndex]
		if _, err := s.heroShopItemsRepo.Create(ctx.GetContext(), models.HeroShopItem{
			ItemID:     item.ID,
			HeroShopID: heroShopID,
			Price:      item.BasePrice, // @todo randomize item price due to the priceCoefficient (discounts?)
		}); err != nil {
			return err
		}
	}

	if _, err := s.heroShopRepo.Update(ctx.GetContext(), heroShopID, &models.HeroShop{
		UpdatesAt: time.Now().Add(ShopItemsRefreshFrequency),
	}); err != nil {
		return err
	}

	return nil
}
