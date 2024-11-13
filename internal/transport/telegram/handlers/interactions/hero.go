package interactions

import (
	"github.com/arthurshafikov/tg-gladiator/internal/core/types"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (h *Handler) HandleHeroDelete(ctx *types.Context, message *tgbotapi.Message, payload []string) error {
	if err := h.ValidatePayloadLength(payload, 1); err != nil {
		return err
	}

	heroID, err := h.GetIDFromString(payload[0])
	if err != nil {
		return err
	}

	var msgText string
	if message.Text == ctx.Messages().DeleteConfirmation {
		if err := h.Services.Hero.DeleteMy(ctx, heroID); err != nil {
			return err
		}

		msgText = ctx.Messages().HeroDeleteSuccessful
	} else {
		msgText = ctx.Messages().HeroDeleteCancelled
	}

	if err := h.Helper.SendTextMessage(ctx.GetChatID(), msgText); err != nil {
		return err
	}

	return h.OpenNewMenu(ctx)
}

func (h *Handler) HandleHeroCreationEnterName(ctx *types.Context, message *tgbotapi.Message, payload []string) error {
	hero, err := h.Services.Hero.Create(ctx, message.Text, payload[0])
	if err != nil {
		return err
	}

	if err := h.Helper.SendTextMessage(ctx.GetChatID(), ctx.Messages().HeroCreationSuccess); err != nil {
		return err
	}

	return h.OpenMyHero(
		ctx,
		hero,
	)
}
