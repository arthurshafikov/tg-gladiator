package queries

import (
	"fmt"

	"github.com/arthurshafikov/tg-gladiator/internal/core/constants/enums"
	"github.com/arthurshafikov/tg-gladiator/internal/core/constants/queries"
	"github.com/arthurshafikov/tg-gladiator/internal/core/constants/telegram"
	"github.com/arthurshafikov/tg-gladiator/internal/core/errors"
	"github.com/arthurshafikov/tg-gladiator/internal/core/models"
	"github.com/arthurshafikov/tg-gladiator/internal/core/types"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (h *Handler) HandleOpenMyHero(ctx *types.Context, query *tgbotapi.CallbackQuery, payload []string) error {
	if err := h.ValidatePayloadLength(payload, 1); err != nil {
		return err
	}

	heroID, err := h.GetIDFromString(payload[0])
	if err != nil {
		return err
	}

	hero, err := h.Services.Hero.FindMy(ctx, heroID)
	if err != nil {
		return err
	}

	return h.openMyHero(
		ctx,
		query,
		hero,
		ctx.Messages().HeroOverview,
	)
}

func (h *Handler) HandleHeroCreationStart(ctx *types.Context, query *tgbotapi.CallbackQuery) error {
	msg := h.Helper.NewEditMessage(ctx.GetChatID(), query.Message.MessageID, ctx.Messages().HeroCreationSelectClass)
	keyboardButtons := make([]telegram.KeyboardButton, 0, len(enums.GetHeroClasses()))
	for _, heroClass := range enums.GetHeroClasses() {
		heroClassText, err := ctx.Messages().GetHeroClass(heroClass)
		if err != nil {
			h.Logger.Error(err)

			return errors.ErrServerError
		}

		keyboardButtons = append(keyboardButtons, telegram.KeyboardButton{
			CallbackQuery: queries.HeroCreationSelectClass.With(heroClass.ToString()),
			Text:          heroClassText,
		})
	}
	keyboardButtons = append(keyboardButtons, telegram.KeyboardButton{
		CallbackQuery: queries.OpenMyHeroes,
		Text:          ctx.Messages().DefaultBackBtn,
	})

	msg.ReplyMarkup = h.Helper.CreateKeyboard(keyboardButtons, 1)

	return h.Helper.Send(msg)
}

func (h *Handler) HandleHeroCreationSelectClass(
	ctx *types.Context,
	query *tgbotapi.CallbackQuery,
	payload []string,
) error {
	if err := h.ValidatePayloadLength(payload, 1); err != nil {
		return err
	}

	hero, err := h.Services.Hero.Create(ctx, payload[0])
	if err != nil {
		return err
	}

	return h.openMyHero(
		ctx,
		query,
		hero,
		ctx.Messages().HeroCreationSuccess,
	)
}

func (h *Handler) openMyHero(ctx *types.Context, query *tgbotapi.CallbackQuery, hero *models.Hero, startText string) error {
	heroInfo, err := ctx.Messages().GetHeroInfo(hero)
	if err != nil {
		h.Logger.Error(err)

		return errors.ErrServerError
	}

	msg := h.Helper.NewEditMessage(
		ctx.GetChatID(),
		query.Message.MessageID,
		fmt.Sprintf(
			"%s\n\n%s",
			startText,
			heroInfo,
		),
	)

	keyboardButtons := make([]telegram.KeyboardButton, 0, 1)

	keyboardButtons = append(
		keyboardButtons,
		telegram.KeyboardButton{
			CallbackQuery: queries.OpenMyHeroes,
			Text:          ctx.Messages().DefaultBackBtn,
		},
	)
	msg.ReplyMarkup = h.Helper.CreateKeyboard(keyboardButtons, 1)

	return h.Helper.Send(msg)
}
