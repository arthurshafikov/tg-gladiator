package telegram

import (
	"context"
	"fmt"

	"github.com/arthurshafikov/tg-gladiator/internal/core/errors"
	"github.com/arthurshafikov/tg-gladiator/internal/core/models"
)

func (b *Bot) authorize(ctx context.Context, chatID int64) (*models.Chat, error) {
	chat, err := b.services.Chat.FindByChatID(ctx, chatID)
	if err != nil {
		if errors.Is(err, errors.ErrNotFound) {
			return nil, fmt.Errorf("not authorized")
		}

		return nil, err
	}

	return chat, nil
}
