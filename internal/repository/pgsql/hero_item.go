package pgsql

import (
	"context"

	"github.com/arthurshafikov/tg-gladiator/internal/core/errors"
	"github.com/arthurshafikov/tg-gladiator/internal/core/models"
	"github.com/arthurshafikov/tg-gladiator/internal/core/types"
	"gorm.io/gorm"
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

func (r *HeroItem) FindBy(ctx context.Context, fields *models.HeroItem) (*models.HeroItem, error) {
	var heroItem models.HeroItem
	if err := r.getDBInstance(ctx).
		Preload("Item").
		Where(fields).
		First(&heroItem).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.ErrNotFound
		}

		return nil, err
	}

	return &heroItem, nil
}

func (r *HeroItem) GetBy(
	ctx context.Context,
	fields *models.HeroItem,
	pagination ...*types.Pagination,
) (*models.PaginatedHeroItem, error) {
	var paginationRequest *types.Pagination
	if len(pagination) > 0 {
		paginationRequest = pagination[0]
	}

	var paginatedResult models.PaginatedHeroItem
	if paginationRequest != nil {
		paginatedResult.Page = pagination[0].Page
		paginatedResult.PerPage = pagination[0].PerPage
	}
	query := r.getDBInstance(ctx).Table("hero_items")

	var heroItems []models.HeroItem
	if err := query.Session(&gorm.Session{}).
		Preload("Item").
		Where(fields).
		Scopes(types.ApplyPagination(paginationRequest)).
		Find(&heroItems).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.ErrNotFound
		}

		return nil, err
	}
	paginatedResult.Rows = heroItems

	if paginationRequest != nil {
		var totalCount int64
		if err := query.Where(fields).Count(&totalCount).Error; err != nil {
			return nil, err
		}
		paginatedResult.TotalCount = totalCount
	}

	return &paginatedResult, nil
}

func (r *HeroItem) UpdateMap(
	ctx context.Context,
	heroItem *models.HeroItem,
	fields map[string]interface{},
) (*models.HeroItem, error) {
	if err := r.getDBInstance(ctx).Model(heroItem).Where(&models.HeroItem{
		HeroID: heroItem.HeroID,
		ItemID: heroItem.ItemID,
	}).Updates(fields).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.ErrNotFound
		}

		return nil, err
	}
	heroItem, err := r.FindBy(ctx, heroItem)
	if err != nil {
		return nil, err
	}

	return heroItem, nil
}
