package pgsql

import (
	"context"

	"github.com/arthurshafikov/tg-gladiator/internal/core/models"
)

type HeroItem struct {
	BaseRepo
}

func NewHeroItemRepository(baseRepo BaseRepo) *HeroItem {
	return &HeroItem{
		BaseRepo: baseRepo,
	}
}

func (r *HeroItem) Create(ctx context.Context, heroItem models.HeroItem) (*models.HeroItem, error) {
	if err := r.getDBInstance(ctx).Create(&heroItem).Error; err != nil {
		return nil, err
	}

	return &heroItem, nil
}

func (r *HeroItem) DeleteBy(ctx context.Context, fields *models.HeroItem) error {
	if err := r.getDBInstance(ctx).Where(fields).Delete(fields).Error; err != nil {
		return err
	}

	return nil
}

// func (r *HeroItem) FindBy(ctx context.Context, fields *models.HeroItem) (*models.HeroItem, error) {
// 	var heroItem models.HeroItem
// 	if err := r.getDBInstance(ctx).
// 		Where(fields).
// 		First(&heroItem).Error; err != nil {
// 		if errors.Is(err, gorm.ErrRecordNotFound) {
// 			return nil, errors.ErrNotFound
// 		}

// 		return nil, err
// 	}

// 	return &heroItem, nil
// }

// func (r *HeroItem) GetBy(ctx context.Context, fields *models.HeroItem) ([]models.HeroItem, error) {
// 	var heroItems []models.HeroItem
// 	if err := r.getDBInstance(ctx).
// 		Preload("Item").
// 		Where(fields).
// 		Find(&heroItems).Error; err != nil {
// 		if errors.Is(err, gorm.ErrRecordNotFound) {
// 			return nil, errors.ErrNotFound
// 		}

// 		return nil, err
// 	}

// 	return heroItems, nil
// }
