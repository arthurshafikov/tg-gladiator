package pgsql

import (
	"context"

	"github.com/arthurshafikov/tg-gladiator/internal/core/errors"
	"github.com/arthurshafikov/tg-gladiator/internal/core/models"
	"gorm.io/gorm"
)

type Hero struct {
	BaseRepo
}

func NewHeroRepository(baseRepo BaseRepo) *Hero {
	return &Hero{
		BaseRepo: baseRepo,
	}
}

func (r *Hero) Create(ctx context.Context, hero models.Hero) (*models.Hero, error) {
	if err := r.getDBInstance(ctx).Create(&hero).Error; err != nil {
		return nil, err
	}

	return &hero, nil
}

func (r *Hero) DeleteByID(ctx context.Context, id int64) error {
	if err := r.getDBInstance(ctx).Delete(&models.Hero{
		ID: id,
	}).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.ErrNotFound
		}

		return err
	}

	return nil
}

func (r *Hero) Find(ctx context.Context, id int64) (*models.Hero, error) {
	fields := &models.Hero{
		ID: id,
	}

	return r.FindBy(ctx, fields)
}

func (r *Hero) FindBy(ctx context.Context, fields *models.Hero) (*models.Hero, error) {
	var hero models.Hero
	if err := r.getDBInstance(ctx).
		Where(fields).
		First(&hero).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.ErrNotFound
		}

		return nil, err
	}

	return &hero, nil
}

func (r *Hero) GetBy(ctx context.Context, fields *models.Hero) (*[]models.Hero, error) {
	var heroes []models.Hero
	if err := r.getDBInstance(ctx).
		Where(fields).
		Find(&heroes).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.ErrNotFound
		}

		return nil, err
	}

	return &heroes, nil
}

func (r *Hero) Update(ctx context.Context, id int64, fields *models.Hero) (*models.Hero, error) {
	if err := r.getDBInstance(ctx).Model(&models.Hero{
		ID: id,
	}).Updates(fields).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.ErrNotFound
		}

		return nil, err
	}
	Hero, err := r.Find(ctx, id)
	if err != nil {
		return nil, err
	}

	return Hero, nil
}

func (r *Hero) UpdateMap(ctx context.Context, id int64, fields map[string]interface{}) (*models.Hero, error) {
	if err := r.getDBInstance(ctx).Model(&models.Hero{
		ID: id,
	}).Updates(fields).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.ErrNotFound
		}

		return nil, err
	}
	Hero, err := r.Find(ctx, id)
	if err != nil {
		return nil, err
	}

	return Hero, nil
}
