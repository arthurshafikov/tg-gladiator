package services

import (
	"fmt"
	"time"

	"github.com/arthurshafikov/tg-gladiator/internal/core/constants/events"
	"github.com/arthurshafikov/tg-gladiator/internal/core/errors"
	"github.com/arthurshafikov/tg-gladiator/internal/core/models"
	"github.com/arthurshafikov/tg-gladiator/internal/core/types"
	"github.com/arthurshafikov/tg-gladiator/internal/repository"
	"github.com/sirupsen/logrus"
)

type HeroShopItemService struct {
	repo            repository.HeroShopItem
	heroShopRepo    repository.HeroShop
	heroService     *HeroService
	heroItemService *HeroItemService
	eventsHandler   EventsHandler
}

func newHeroShopItemService(
	repo repository.HeroShopItem,
	heroShopRepo repository.HeroShop,
	heroService *HeroService,
	heroItemService *HeroItemService,
	eventsHandler EventsHandler,
) *HeroShopItemService {
	return &HeroShopItemService{
		repo:            repo,
		heroShopRepo:    heroShopRepo,
		heroService:     heroService,
		heroItemService: heroItemService,
		eventsHandler:   eventsHandler,
	}
}

func (s *HeroShopItemService) BuyItem(ctx *types.Context, heroID, itemID int64) (*models.HeroShopItem, error) {
	shopItem, err := s.FindMy(ctx, heroID, itemID)
	if err != nil {
		return nil, err
	}

	hero, err := s.heroService.FindMy(ctx, heroID)
	if err != nil {
		return nil, err
	}

	if hero.CurrentGold < shopItem.Price {
		return nil, errors.ErrInsufficientGold
	}

	// @todo transaction
	if err := s.heroService.decreaseGold(ctx, heroID, shopItem.Price); err != nil {
		return nil, err
	}

	if _, err := s.heroItemService.PutIntoInventory(ctx, heroID, itemID); err != nil {
		return nil, err
	}

	if !shopItem.Item.IsPotion() {
		if err := s.repo.DeleteBy(ctx.GetContext(), &models.HeroShopItem{
			HeroShopID: shopItem.HeroShopID,
			ItemID:     shopItem.ItemID,
		}); err != nil {
			return nil, err
		}
	}

	s.eventsHandler.Dispatch(
		events.AnalyticEvent,
		models.CreateAnalyticEventDTO{
			ChatID:    hero.ChatID,
			Type:      models.AnalyticEventTypeItemPurchased,
			Timestamp: time.Now(),
			Payload: map[string]any{
				"item_id": itemID,
			},
		},
	)

	return shopItem, nil
}

func (s *HeroShopItemService) FindMy(ctx *types.Context, heroID, itemID int64) (*models.HeroShopItem, error) {
	hero, err := s.heroService.FindMy(ctx, heroID)
	if err != nil {
		return nil, err
	}

	heroShop, err := s.heroShopRepo.FindBy(ctx.GetContext(), &models.HeroShop{
		HeroID: hero.ID,
	})
	if err != nil {
		if errors.Is(err, errors.ErrNotFound) {
			logrus.Error(fmt.Errorf("shop not exists error for hero id: %v", hero.ID))

			return nil, errors.ErrShopNotExists
		}

		logrus.Error(err)

		return nil, errors.ErrServerError
	}

	if heroShop.UpdatesAt.Before(time.Now()) { // @todo helper
		return nil, errors.ErrHeroShopItemNotExists
	}

	shopItem, err := s.repo.FindBy(ctx.GetContext(), &models.HeroShopItem{
		HeroShopID: heroShop.ID,
		ItemID:     itemID,
	})
	if err != nil {
		if errors.Is(err, errors.ErrNotFound) {
			return nil, errors.ErrHeroShopItemNotExists
		}

		logrus.Error(err)

		return nil, errors.ErrServerError
	}

	return shopItem, nil
}
