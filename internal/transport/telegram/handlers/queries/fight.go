package queries

import (
	"fmt"

	"github.com/arthurshafikov/tg-gladiator/internal/core/constants/enums"
	"github.com/arthurshafikov/tg-gladiator/internal/core/constants/queries"
	"github.com/arthurshafikov/tg-gladiator/internal/core/constants/telegram"
	"github.com/arthurshafikov/tg-gladiator/internal/core/helpers"
	"github.com/arthurshafikov/tg-gladiator/internal/core/types"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (h *Handler) HandleOpenFightOptions(ctx *types.Context, query *tgbotapi.CallbackQuery, payload []string) error {
	if err := h.ValidatePayloadLength(payload, 1); err != nil {
		return err
	}

	heroID, err := h.GetIDFromString(payload[0])
	if err != nil {
		return err
	}

	keyboardButtons := make([]telegram.KeyboardButton, 0, 10)

	for _, location := range enums.AllFightLocations() {
		keyboardButtons = append(
			keyboardButtons,
			telegram.KeyboardButton{
				CallbackQuery: queries.StartFight.WithID(heroID).With(string(location)),
				Text:          ctx.Messages().FightLocations[string(location)],
			},
		)
	}
	keyboard := h.Helper.CreateKeyboard(keyboardButtons, 2)

	bossesLeft, err := h.Services.BossFight.GetBossesAmountLeft(ctx, heroID)
	if err != nil {
		return err
	}

	keyboard.InlineKeyboard = append(keyboard.InlineKeyboard, tgbotapi.NewInlineKeyboardRow(
		tgbotapi.InlineKeyboardButton{
			CallbackData: helpers.GetPointer(string(queries.ChallengeBoss.WithID(heroID))),
			Text:         fmt.Sprintf(ctx.Messages().MenuItemChallengeBoss, bossesLeft),
		},
	))
	keyboard.InlineKeyboard = append(keyboard.InlineKeyboard, h.GetBackButtonKeyboardRow(ctx, queries.OpenMyHero.WithID(heroID)))

	msgText := fmt.Sprintf("%s\n", ctx.Messages().FightSelectLocation)

	for _, location := range enums.AllFightLocations() {
		minLevel, maxLevel := enums.FightLocationLevelRange(location)

		msgText += fmt.Sprintf(
			"\n- %s (%v - %v %s)",
			ctx.Messages().FightLocations[string(location)],
			minLevel,
			maxLevel,
			ctx.Messages().Level,
		)
	}

	msg := h.Helper.NewEditMessage(
		ctx.GetChatID(),
		query.Message.MessageID,
		msgText,
	)
	msg.ReplyMarkup = keyboard

	return h.Helper.Send(msg)
}
