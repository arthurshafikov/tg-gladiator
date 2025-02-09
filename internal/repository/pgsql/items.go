package pgsql

import (
	"context"

	"github.com/arthurshafikov/tg-gladiator/internal/core/errors"
	"github.com/arthurshafikov/tg-gladiator/internal/core/models"
	"gorm.io/gorm"
)

type Item struct {
	BaseRepo
}

func NewItemRepository(baseRepo BaseRepo) *Item {
	return &Item{
		BaseRepo: baseRepo,
	}
}

func (r *Item) GetBy(ctx context.Context, fields *models.Item) ([]models.Item, error) {
	query := r.getDBInstance(ctx)

	if fields != nil {
		query = query.Where(fields)
	}

	var items []models.Item
	if err := query.
		Find(&items).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.ErrNotFound
		}

		return nil, err
	}

	return items, nil
}
