package pgsql

import (
	"context"

	"github.com/arthurshafikov/tg-gladiator/internal/core/errors"
	"github.com/arthurshafikov/tg-gladiator/internal/core/models"
	"gorm.io/gorm"
)

type HeroShop struct {
	BaseRepo
}

func NewHeroShopRepository(baseRepo BaseRepo) *HeroShop {
	return &HeroShop{
		BaseRepo: baseRepo,
	}
}

func (r *HeroShop) Find(ctx context.Context, id int64) (*models.HeroShop, error) {
	fields := &models.HeroShop{
		ID: id,
	}

	return r.FindBy(ctx, fields)
}

func (r *HeroShop) FindBy(ctx context.Context, fields *models.HeroShop) (*models.HeroShop, error) {
	var heroShop models.HeroShop
	if err := r.getDBInstance(ctx).
		Where(fields).
		First(&heroShop).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.ErrNotFound
		}

		return nil, err
	}

	return &heroShop, nil
}

func (r *HeroShop) Create(ctx context.Context, heroShop models.HeroShop) (*models.HeroShop, error) {
	if err := r.getDBInstance(ctx).
		Where("hero_id = ?", heroShop.HeroID).
		First(&models.HeroShop{}).
		Error; err == nil {
		return nil, errors.ErrAlreadyExists
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	if err := r.getDBInstance(ctx).Create(&heroShop).Error; err != nil {
		return nil, err
	}

	heroShopWithAllData, err := r.Find(ctx, heroShop.ID)
	if err != nil {
		return nil, err
	}

	return heroShopWithAllData, nil
}

func (r *HeroShop) Update(ctx context.Context, id int64, fields *models.HeroShop) (*models.HeroShop, error) {
	if err := r.getDBInstance(ctx).Model(&models.HeroShop{
		ID: id,
	}).Updates(fields).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.ErrNotFound
		}

		return nil, err
	}
	heroShop, err := r.Find(ctx, id)
	if err != nil {
		return nil, err
	}

	return heroShop, nil
}

func (r *HeroShop) UpdateMap(ctx context.Context, id int64, fields map[string]interface{}) (*models.HeroShop, error) {
	if err := r.getDBInstance(ctx).Model(&models.HeroShop{
		ID: id,
	}).Updates(fields).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.ErrNotFound
		}

		return nil, err
	}
	heroShop, err := r.Find(ctx, id)
	if err != nil {
		return nil, err
	}

	return heroShop, nil
}
