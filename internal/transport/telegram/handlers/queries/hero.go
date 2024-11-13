package queries

import (
	"fmt"

	"github.com/arthurshafikov/tg-gladiator/internal/core/constants/enums"
	"github.com/arthurshafikov/tg-gladiator/internal/core/constants/interactions"
	"github.com/arthurshafikov/tg-gladiator/internal/core/constants/queries"
	"github.com/arthurshafikov/tg-gladiator/internal/core/constants/telegram"
	"github.com/arthurshafikov/tg-gladiator/internal/core/errors"
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

	return h.OpenMyHero(
		ctx,
		hero,
		query,
	)
}

func (h *Handler) HandleHeroCreationStart(ctx *types.Context, query *tgbotapi.CallbackQuery) error {
	msgText := ctx.Messages().HeroCreationSelectClass

	keyboardButtons := make([]telegram.KeyboardButton, 0, len(enums.GetHeroClasses()))
	for _, heroClass := range enums.GetHeroClasses() {
		heroClassText, err := ctx.Messages().GetHeroClass(heroClass)
		if err != nil {
			h.Logger.Error(err)

			return errors.ErrServerError
		}

		classInfo, err := ctx.Messages().ClassInfo(heroClass)
		if err != nil {
			h.Logger.Error(err)

			return errors.ErrServerError
		}

		msgText += fmt.Sprintf("\n\n%s", classInfo)

		keyboardButtons = append(keyboardButtons, telegram.KeyboardButton{
			CallbackQuery: queries.HeroCreationSelectClass.With(heroClass.ToString()),
			Text:          heroClassText,
		})
	}
	keyboardButtons = append(keyboardButtons, telegram.KeyboardButton{
		CallbackQuery: queries.OpenMyHeroes,
		Text:          ctx.Messages().DefaultBackBtn,
	})
	msg := h.Helper.NewEditMessage(ctx.GetChatID(), query.Message.MessageID, msgText)
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

	if err := h.Services.Interactions.Set(ctx, interactions.HeroCreationEnterName.With(payload[0])); err != nil {
		return err
	}

	msg := h.Helper.NewEditMessage(ctx.GetChatID(), query.Message.MessageID, ctx.Messages().HeroCreationEnterNamePrompt)
	msg.ReplyMarkup = h.GetKeyboardWithBackButton(ctx, queries.OpenMyHeroes)

	return h.Helper.Send(msg)
}

func (h *Handler) HandleHeroDelete(ctx *types.Context, query *tgbotapi.CallbackQuery, payload []string) error {
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

	if err := h.Services.Interactions.Set(ctx, interactions.HeroDelete.WithID(hero.ID)); err != nil {
		return err
	}

	heroName, err := ctx.Messages().GetHeroName(hero)
	if err != nil {
		h.Logger.Error(err)

		return errors.ErrServerError
	}

	msg := h.Helper.NewEditMessage(
		ctx.GetChatID(),
		query.Message.MessageID,
		fmt.Sprintf(
			ctx.Messages().HeroDeleteConfirmationPrompt,
			heroName, // @todo name field in heroes table
			ctx.Messages().DeleteConfirmation,
		),
	)
	msg.ReplyMarkup = h.GetKeyboardWithBackButton(ctx, queries.OpenMyHero.WithID(hero.ID))

	return h.Helper.Send(msg)
}
