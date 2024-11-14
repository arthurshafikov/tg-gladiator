package pgsql

import (
	"context"

	"github.com/arthurshafikov/tg-gladiator/internal/core/errors"
	"github.com/arthurshafikov/tg-gladiator/internal/core/models"
	"gorm.io/gorm"
)

type Enemy struct {
	BaseRepo
}

func NewEnemyRepository(baseRepo BaseRepo) *Enemy {
	return &Enemy{
		BaseRepo: baseRepo,
	}
}

func (r *Enemy) Find(ctx context.Context, id int64) (*models.Enemy, error) {
	fields := &models.Enemy{
		ID: id,
	}

	return r.FindBy(ctx, fields)
}

func (r *Enemy) FindBy(ctx context.Context, fields *models.Enemy) (*models.Enemy, error) {
	var enemy models.Enemy
	if err := r.getDBInstance(ctx).
		Where(fields).
		First(&enemy).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.ErrNotFound
		}

		return nil, err
	}

	return &enemy, nil
}

func (r *Enemy) GetBy(ctx context.Context, fields *models.Enemy) (*[]models.Enemy, error) {
	var enemies []models.Enemy
	if err := r.getDBInstance(ctx).
		Where(fields).
		Find(&enemies).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.ErrNotFound
		}

		return nil, err
	}

	return &enemies, nil
}

// func (r *Enemy) Create(ctx context.Context, enemy models.Enemy) (*models.Enemy, error) {
// 	if enemy.EnemyID == 0 {
// 		return nil, errors.ErrServerError
// 	}

// 	if err := r.getDBInstance(ctx).
// 		Where("enemy_id = ?", enemy.EnemyID).
// 		First(&models.Enemy{}).
// 		Error; err == nil {
// 		return nil, errors.ErrAlreadyExists
// 	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
// 		return nil, err
// 	}

// 	if err := r.getDBInstance(ctx).Create(&enemy).Error; err != nil {
// 		return nil, err
// 	}

// 	return &enemy, nil
// }

// func (r *Enemy) Update(ctx context.Context, id int64, fields *models.Enemy) (*models.Enemy, error) {
// 	if err := r.getDBInstance(ctx).Model(&models.Enemy{
// 		ID: id,
// 	}).Updates(fields).Error; err != nil {
// 		if errors.Is(err, gorm.ErrRecordNotFound) {
// 			return nil, errors.ErrNotFound
// 		}

// 		return nil, err
// 	}
// 	enemy, err := r.Find(ctx, id)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return enemy, nil
// }

// func (r *Enemy) UpdateMap(ctx context.Context, id int64, fields map[string]interface{}) (*models.Enemy, error) {
// 	if err := r.getDBInstance(ctx).Model(&models.Enemy{
// 		ID: id,
// 	}).Updates(fields).Error; err != nil {
// 		if errors.Is(err, gorm.ErrRecordNotFound) {
// 			return nil, errors.ErrNotFound
// 		}

// 		return nil, err
// 	}
// 	enemy, err := r.Find(ctx, id)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return enemy, nil
// }
