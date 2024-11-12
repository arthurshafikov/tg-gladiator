package queries

import (
	"github.com/arthurshafikov/tg-gladiator/internal/core/constants/enums"
	"github.com/arthurshafikov/tg-gladiator/internal/core/constants/queries"
	"github.com/arthurshafikov/tg-gladiator/internal/core/constants/telegram"
	"github.com/arthurshafikov/tg-gladiator/internal/core/types"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (h *Handler) HandleStartHeroCreation(ctx *types.Context, query *tgbotapi.CallbackQuery) error {
	msg := h.Helper.NewEditMessage(ctx.GetChatID(), query.Message.MessageID, ctx.Messages().HeroCreationSelectClass)
	keyboardButtons := make([]telegram.KeyboardButton, 0, len(enums.GetHeroClasses()))
	for _, heroClass := range enums.GetHeroClasses() {
		keyboardButtons = append(keyboardButtons, telegram.KeyboardButton{
			CallbackQuery: queries.HeroCreationSelectClass,
			Text: ctx.Messages().GetHeroClass(heroClass),
		})
	}

	msg.ReplyMarkup = h.Helper.CreateKeyboard(keyboardButtons, 1)

	return h.Helper.Send(msg)
}
