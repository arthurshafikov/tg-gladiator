package listeners

import (
	"context"
	"fmt"

	"github.com/arthurshafikov/tg-gladiator/internal/config"
	"github.com/arthurshafikov/tg-gladiator/internal/core/models"
	"github.com/arthurshafikov/tg-gladiator/internal/core/types"
	"github.com/arthurshafikov/tg-gladiator/internal/services"
)

type HeroLevelUpListener struct {
	notificationsHandler NotificationsHandler
	chatService          services.Chat

	messages *config.MessagesBag
}

func NewHeroLevelUpListener(
	notificationsHandler NotificationsHandler,
	chatService services.Chat,
	messages *config.MessagesBag,
) *HeroLevelUpListener {
	return &HeroLevelUpListener{
		notificationsHandler: notificationsHandler,
		chatService:          chatService,
		messages:             messages,
	}
}

func (l *HeroLevelUpListener) Handle(ctx context.Context, params ...any) error {
	if len(params) < 1 {
		return fmt.Errorf("params len < 1 HeroLevelUpListener")
	}

	hero, ok := params[0].(models.Hero)
	if !ok {
		return fmt.Errorf("params[0] is not a hero HeroLevelUpListener: %#v", params)
	}

	chat, err := l.chatService.FindByID(ctx, hero.ChatID)
	if err != nil {
		return fmt.Errorf("HeroLevelUpListener chatService find err: %w", err)
	}

	newCtx := types.NewContext(ctx, chat, l.messages.GetMessages(chat.Language))

	return l.notificationsHandler.NotifyAboutHeroLevelUp(newCtx, &hero)
}
