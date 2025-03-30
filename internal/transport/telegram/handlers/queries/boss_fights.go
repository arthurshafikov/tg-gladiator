package queries

import (
	"github.com/arthurshafikov/tg-gladiator/internal/core/types"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (h *Handler) HandleChallengeBoss(ctx *types.Context, query *tgbotapi.CallbackQuery, payload []string) error {
	if err := h.ValidatePayloadLength(payload, 1); err != nil {
		return err
	}

	heroID, err := h.GetIDFromString(payload[0])
	if err != nil {
		return err
	}

	fight, err := h.Services.BossFight.Create(ctx, heroID)
	if err != nil {
		return err
	}

	return h.SendFightOverviewMessage(ctx, fight, query)
}
