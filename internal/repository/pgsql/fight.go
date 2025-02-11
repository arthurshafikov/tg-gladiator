package pgsql

import (
	"context"

	"github.com/arthurshafikov/tg-gladiator/internal/core/constants/enums"
	"github.com/arthurshafikov/tg-gladiator/internal/core/errors"
	"github.com/arthurshafikov/tg-gladiator/internal/core/models"
	"gorm.io/gorm"
)

type Fight struct {
	BaseRepo
}

func NewFightRepository(baseRepo BaseRepo) *Fight {
	return &Fight{
		BaseRepo: baseRepo,
	}
}

func (r *Fight) FindBy(ctx context.Context, fields *models.Fight) (*models.Fight, error) {
	return r.findBy(ctx, fields)
}

func (r *Fight) FindByMap(ctx context.Context, fields map[string]interface{}) (*models.Fight, error) {
	return r.findBy(ctx, fields)
}

func (r *Fight) Create(ctx context.Context, fight models.Fight) (*models.Fight, error) {
	if err := r.getDBInstance(ctx).
		Where("hero_id = ?", fight.HeroID).
		Where("status = ?", enums.FightStatusActive).
		First(&models.Fight{}).
		Error; err == nil {
		return nil, errors.ErrAlreadyExists
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	if err := r.getDBInstance(ctx).Create(&fight).Error; err != nil {
		return nil, err
	}

	return &fight, nil
}

func (r *Fight) Update(ctx context.Context, id int64, fields *models.Fight) error {
	if err := r.getDBInstance(ctx).Model(&models.Fight{
		ID: id,
	}).Updates(fields).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.ErrNotFound
		}

		return err
	}

	return nil
}

func (r *Fight) UpdateMap(ctx context.Context, id int64, fields map[string]interface{}) error {
	if err := r.getDBInstance(ctx).Model(&models.Fight{
		ID: id,
	}).Updates(fields).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.ErrNotFound
		}

		return err
	}

	return nil
}

func (r *Fight) find(ctx context.Context, id int64) (*models.Fight, error) {
	fields := &models.Fight{
		ID: id,
	}

	return r.FindBy(ctx, fields)
}

func (r *Fight) findBy(ctx context.Context, fields interface{}) (*models.Fight, error) {
	var fight models.Fight
	if err := r.getDBInstance(ctx).
		Where(fields).
		First(&fight).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.ErrNotFound
		}

		return nil, err
	}

	return &fight, nil
}
