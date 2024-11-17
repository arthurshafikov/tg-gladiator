package queries

import (
	"github.com/arthurshafikov/tg-gladiator/internal/core/errors"
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

	fight, err := h.Services.TournamentFight.Create(ctx, heroID)
	if err != nil {
		return err
	}

	msgText, err := ctx.Messages().FightInfo(fight)
	if err != nil {
		h.Logger.Error(err)

		return errors.ErrServerError
	}
	msg := h.Helper.NewEditMessage(ctx.GetChatID(), query.Message.MessageID, msgText)

	// @todo
	// show buttons (punch, run away)
	// msg.ReplyMarkup =

	return h.Helper.Send(msg)
}
