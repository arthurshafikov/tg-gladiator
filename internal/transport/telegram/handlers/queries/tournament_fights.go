package queries

import (
	"github.com/arthurshafikov/tg-gladiator/internal/core/types"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (h *Handler) HandleStartTournamentFIght(ctx *types.Context, query *tgbotapi.CallbackQuery, payload []string) error {
	if err := h.ValidatePayloadLength(payload, 1); err != nil {
		return err
	}

	heroID, err := h.GetIDFromString(payload[0])
	if err != nil {
		return err
	}

	_, err = h.Services.TournamentFight.Create(ctx, heroID)
	if err != nil {
		return err
	}

	// @todo show a fight window (stats for both sides)
	msg := h.Helper.NewEditMessage(ctx.GetChatID(), query.Message.MessageID, "") //ctx.Messages().FightInfo(fight))

	// @todo
	// show buttons (punch, run away)
	// msg.ReplyMarkup =

	return h.Helper.Send(msg)
}
