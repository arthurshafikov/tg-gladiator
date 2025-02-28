package services

import (
	"context"
	"fmt"
	"time"

	"github.com/arthurshafikov/tg-gladiator/internal/core/errors"
	"github.com/arthurshafikov/tg-gladiator/internal/core/models"
	"github.com/arthurshafikov/tg-gladiator/internal/core/types"
	"github.com/arthurshafikov/tg-gladiator/internal/repository"
)

type ChatService struct {
	logger Logger
	repo   repository.Chat
}

func newChatService(
	logger Logger,
	repo repository.Chat,
) *ChatService {
	return &ChatService{
		logger: logger,
		repo:   repo,
	}
}

func (s *ChatService) FindByID(ctx context.Context, id int64) (*models.Chat, error) {
	chat, err := s.repo.FindBy(ctx, &models.Chat{
		ID: id,
	})
	if err != nil {
		if !errors.Is(err, errors.ErrNotFound) {
			s.logger.Error(err)

			return nil, errors.ErrServerError
		}

		return nil, errors.ErrNotFound
	}

	return chat, nil
}

func (s *ChatService) FindByChatID(ctx context.Context, chatID int64) (*models.Chat, error) {
	chat, err := s.repo.FindBy(ctx, &models.Chat{
		ChatID: chatID,
	})
	if err != nil {
		if !errors.Is(err, errors.ErrNotFound) {
			s.logger.Error(err)

			return nil, errors.ErrServerError
		}

		return nil, errors.ErrNotFound
	}

	return chat, nil
}

func (s *ChatService) FirstOrCreate(ctx context.Context, data models.Chat) (*models.Chat, bool, error) {
	chat, err := s.repo.FindBy(ctx, &models.Chat{
		ChatID: data.ChatID,
	})
	if err == nil {
		return chat, true, nil
	}
	if !errors.Is(err, errors.ErrNotFound) {
		return nil, false, err
	}
	data.Language = "RU"
	data.LatestActiveAt = time.Now()

	chat, err = s.repo.Create(ctx, data)
	if err != nil {
		s.logger.Error(err)

		return nil, false, errors.ErrServerError
	}

	return chat, false, nil
}

func (s *ChatService) UpdateLatestActiveAt(ctx *types.Context) {
	if _, err := s.update(ctx.GetContext(), ctx.GetChat().ID, &models.Chat{
		LatestActiveAt: time.Now(),
	}); err != nil {
		s.logger.Error(fmt.Errorf("update latest active at error: %w", err))
	}
}

func (s *ChatService) UpdateUsername(ctx *types.Context, username string) error {
	if len(username) < 1 {
		return nil
	}

	fields := &models.Chat{
		Username: username,
	}
	if _, err := s.update(ctx.GetContext(), ctx.GetChat().ID, fields); err != nil {
		return err
	}

	return nil
}

func (s *ChatService) update(ctx context.Context, id int64, fields *models.Chat) (*models.Chat, error) {
	chat, err := s.repo.Update(ctx, id, fields)
	if err != nil {
		if !errors.Is(err, errors.ErrNotFound) {
			s.logger.Error(err)

			return nil, errors.ErrServerError
		}

		return nil, errors.ErrNotFound
	}

	return chat, nil
}
