package notifications

import (
	"fmt"

	"github.com/arthurshafikov/tg-gladiator/internal/core/models"
	"github.com/arthurshafikov/tg-gladiator/internal/core/types"
)

func (h *Handler) NotifyAboutHeroLevelUp(ctx *types.Context, hero *models.Hero) error {
	return h.Helper.SendTextMessage(ctx.GetChatID(), fmt.Sprintf(
		ctx.Messages().HeroLevelUpNotification,
		hero.Name,
		hero.Level,
	))
}
