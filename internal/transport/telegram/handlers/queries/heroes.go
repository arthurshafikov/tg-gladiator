package queries

import (
	"github.com/arthurshafikov/tg-gladiator/internal/core/constants/queries"
	"github.com/arthurshafikov/tg-gladiator/internal/core/constants/telegram"
	"github.com/arthurshafikov/tg-gladiator/internal/core/errors"
	"github.com/arthurshafikov/tg-gladiator/internal/core/types"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (h *Handler) HandleOpenMyHeroes(ctx *types.Context, query *tgbotapi.CallbackQuery) error {
	heroes, err := h.Services.Heroes.GetMy(ctx)
	if err != nil {
		return err
	}
	if heroes == nil {
		return errors.ErrServerError
	}

	keyboardButtons := make([]telegram.KeyboardButton, 0, len(*heroes))

	for _, hero := range *heroes {
		keyboardButtons = append(keyboardButtons, telegram.KeyboardButton{
			CallbackQuery: queries.OpenMyHeroes, // @todo OpenMyHero
			Text:          hero.GetName(ctx.Messages()),
		})
	}
	keyboardButtons = append(keyboardButtons, telegram.KeyboardButton{
		CallbackQuery: queries.OpenMenu,
		Text:          ctx.Messages().DefaultBackBtn,
	})

	msg := h.Helper.NewEditMessageText(ctx.GetChatID(), query.Message.MessageID, ctx.Messages().MyHeroesList)
	msg.ReplyMarkup = h.Helper.CreateKeyboard(keyboardButtons, 1)

	return h.Helper.Send(msg)
}
