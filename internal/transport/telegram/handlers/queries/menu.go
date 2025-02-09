package queries

import (
	"github.com/arthurshafikov/tg-gladiator/internal/core/types"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (h *Handler) HandleOpenMenu(ctx *types.Context, query *tgbotapi.CallbackQuery, payload []string) error {
	return h.EditMessageOpenMenu(ctx, query.Message)
}
