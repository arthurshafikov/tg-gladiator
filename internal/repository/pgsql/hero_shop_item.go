package pgsql

import (
	"context"

	"github.com/arthurshafikov/tg-gladiator/internal/core/errors"
	"github.com/arthurshafikov/tg-gladiator/internal/core/models"
	"gorm.io/gorm"
)

type HeroShopItem struct {
	BaseRepo
}

func NewHeroShopItemRepository(baseRepo BaseRepo) *HeroShopItem {
	return &HeroShopItem{
		BaseRepo: baseRepo,
	}
}

func (r *HeroShopItem) Create(ctx context.Context, heroShopItem models.HeroShopItem) (*models.HeroShopItem, error) {
	if err := r.getDBInstance(ctx).Create(&heroShopItem).Error; err != nil {
		return nil, err
	}

	return &heroShopItem, nil
}

func (r *HeroShopItem) DeleteBy(ctx context.Context, fields *models.HeroShopItem) error {
	if err := r.getDBInstance(ctx).Where(fields).Delete(fields).Error; err != nil {
		return err
	}

	return nil
}

func (r *HeroShopItem) GetBy(ctx context.Context, fields *models.HeroShopItem) ([]models.HeroShopItem, error) {
	var heroShopItems []models.HeroShopItem
	if err := r.getDBInstance(ctx).
		Preload("Item").
		Where(fields).
		Find(&heroShopItems).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.ErrNotFound
		}

		return nil, err
	}

	return heroShopItems, nil
}
