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

func (r *Fight) Find(ctx context.Context, id int64) (*models.Fight, error) {
	fields := &models.Fight{
		ID: id,
	}

	return r.FindBy(ctx, fields)
}

func (r *Fight) FindBy(ctx context.Context, fields *models.Fight) (*models.Fight, error) {
	var fight models.Fight
	if err := r.getDBInstance(ctx).
		Preload("Hero").
		Where(fields).
		First(&fight).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.ErrNotFound
		}

		return nil, err
	}

	if fight.OpponentType == enums.OpponentTypeMob {
		enemy := models.Enemy{
			ID: fight.OpponentID,
		}
		if err := r.getDBInstance(ctx).Find(&enemy).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, errors.ErrNotFound
			}

			return nil, err
		}

		fight.Opponent = &enemy
	}

	return &fight, nil
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

	fightWithAllData, err := r.Find(ctx, fight.ID)
	if err != nil {
		return nil, err
	}

	return fightWithAllData, nil
}

func (r *Fight) Update(ctx context.Context, id int64, fields *models.Fight) (*models.Fight, error) {
	if err := r.getDBInstance(ctx).Model(&models.Fight{
		ID: id,
	}).Updates(fields).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.ErrNotFound
		}

		return nil, err
	}
	fight, err := r.Find(ctx, id)
	if err != nil {
		return nil, err
	}

	return fight, nil
}

func (r *Fight) UpdateMap(ctx context.Context, id int64, fields map[string]interface{}) (*models.Fight, error) {
	if err := r.getDBInstance(ctx).Model(&models.Fight{
		ID: id,
	}).Updates(fields).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.ErrNotFound
		}

		return nil, err
	}
	fight, err := r.Find(ctx, id)
	if err != nil {
		return nil, err
	}

	return fight, nil
}
