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
		heroName, err := ctx.Messages().GetHeroName(&hero)
		if err != nil {
			h.Logger.Error(err)

			return errors.ErrServerError
		}

		keyboardButtons = append(keyboardButtons, telegram.KeyboardButton{
			CallbackQuery: queries.OpenMyHero.WithID(hero.ID),
			Text:          heroName,
		})
	}
	keyboardButtons = append(keyboardButtons, telegram.KeyboardButton{
		CallbackQuery: queries.HeroCreationStart,
		Text:          ctx.Messages().HeroCreationStart,
	})
	keyboardButtons = append(keyboardButtons, telegram.KeyboardButton{
		CallbackQuery: queries.OpenMenu,
		Text:          ctx.Messages().DefaultBackBtn,
	})

	msg := h.Helper.NewEditMessage(ctx.GetChatID(), query.Message.MessageID, ctx.Messages().MyHeroesList)
	msg.ReplyMarkup = h.Helper.CreateKeyboard(keyboardButtons, 1)

	return h.Helper.Send(msg)
}
