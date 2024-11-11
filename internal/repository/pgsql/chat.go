package pgsql

import (
	"context"

	"github.com/arthurshafikov/tg-gladiator/internal/core/errors"
	"github.com/arthurshafikov/tg-gladiator/internal/core/models"
	"gorm.io/gorm"
)

type Chat struct {
	BaseRepo
}

func NewChatRepository(baseRepo BaseRepo) *Chat {
	return &Chat{
		BaseRepo: baseRepo,
	}
}

func (r *Chat) Find(ctx context.Context, id int64) (*models.Chat, error) {
	fields := &models.Chat{
		ID: id,
	}

	return r.FindBy(ctx, fields)
}

func (r *Chat) FindBy(ctx context.Context, fields *models.Chat) (*models.Chat, error) {
	var chat models.Chat
	if err := r.getDBInstance(ctx).
		Where(fields).
		First(&chat).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.ErrNotFound
		}

		return nil, err
	}

	return &chat, nil
}

func (r *Chat) Create(ctx context.Context, chat models.Chat) (*models.Chat, error) {
	if chat.ChatID == 0 {
		return nil, errors.ErrServerError
	}

	if err := r.getDBInstance(ctx).
		Where("chat_id = ?", chat.ChatID).
		First(&models.Chat{}).
		Error; err == nil {
		return nil, errors.ErrAlreadyExists
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	if err := r.getDBInstance(ctx).Create(&chat).Error; err != nil {
		return nil, err
	}

	return &chat, nil
}

func (r *Chat) Update(ctx context.Context, id int64, fields *models.Chat) (*models.Chat, error) {
	if err := r.getDBInstance(ctx).Model(&models.Chat{
		ID: id,
	}).Updates(fields).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.ErrNotFound
		}

		return nil, err
	}
	chat, err := r.Find(ctx, id)
	if err != nil {
		return nil, err
	}

	return chat, nil
}

func (r *Chat) UpdateMap(ctx context.Context, id int64, fields map[string]interface{}) (*models.Chat, error) {
	if err := r.getDBInstance(ctx).Model(&models.Chat{
		ID: id,
	}).Updates(fields).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.ErrNotFound
		}

		return nil, err
	}
	chat, err := r.Find(ctx, id)
	if err != nil {
		return nil, err
	}

	return chat, nil
}
