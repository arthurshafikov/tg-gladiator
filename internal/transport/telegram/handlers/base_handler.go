package handlers

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/arthurshafikov/tg-gladiator/internal/config"
	"github.com/arthurshafikov/tg-gladiator/internal/core/constants/queries"
	"github.com/arthurshafikov/tg-gladiator/internal/core/constants/telegram"
	"github.com/arthurshafikov/tg-gladiator/internal/core/errors"
	"github.com/arthurshafikov/tg-gladiator/internal/core/models"
	"github.com/arthurshafikov/tg-gladiator/internal/core/types"
	"github.com/arthurshafikov/tg-gladiator/internal/services"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type BaseHandler struct {
	Services *services.Services
	Logger   services.Logger
	Helper   TelegramHandlerHelper
	Config   *config.Config
}

func NewBaseHandler(tgHandlerParams HandlerParams) *BaseHandler {
	return &BaseHandler{
		Services: tgHandlerParams.Services,
		Logger:   tgHandlerParams.Logger,
		Helper:   tgHandlerParams.Helper,
		Config:   tgHandlerParams.Config,
	}
}

func (h *BaseHandler) ValidatePayloadLength(payload []string, expectedMinLength int) error {
	if len(payload) < expectedMinLength {
		h.Logger.Error(fmt.Sprintf(
			"Payload Len is not %v, payload: %s",
			expectedMinLength,
			strings.Join(payload, queries.SpecialDelimeterInQueryCallback),
		))

		return errors.ErrServerError
	}

	return nil
}

func (h *BaseHandler) OpenNewMenu(ctx *types.Context) error {
	if err := h.Services.Interactions.Remove(ctx); err != nil {
		return err
	}

	msg := h.Helper.NewMessage(
		ctx.GetChatID(),
		ctx.Messages().OpenedMenu,
	)
	msg.ReplyMarkup = h.getMenuKeyboard(ctx)

	return h.Helper.Send(msg)
}

func (h *BaseHandler) EditMessageOpenMenu(ctx *types.Context, message *tgbotapi.Message) error {
	if err := h.Services.Interactions.Remove(ctx); err != nil {
		return err
	}

	msg := h.Helper.NewEditMessage(
		ctx.GetChatID(), message.MessageID,
		ctx.Messages().OpenedMenu,
	)
	msg.ReplyMarkup = h.getMenuKeyboard(ctx)

	return h.Helper.Send(msg)
}

func (h *BaseHandler) GetKeyboardWithBackButton(ctx *types.Context, backQuery ...queries.Query) *tgbotapi.InlineKeyboardMarkup {
	query := queries.OpenMenu
	if len(backQuery) > 0 {
		query = backQuery[0]
	}

	keyboard := tgbotapi.InlineKeyboardMarkup{}
	keyboard.InlineKeyboard = append(
		keyboard.InlineKeyboard,
		h.GetBackButtonKeyboardRow(ctx, query),
	)

	return &keyboard
}

func (h *BaseHandler) GetBackButtonKeyboardRow(
	ctx *types.Context,
	queryValues ...queries.Query,
) []tgbotapi.InlineKeyboardButton {
	query := queries.OpenMenu
	if len(queryValues) > 0 {
		query = queryValues[0]
	}

	return h.Helper.NewRow(ctx.Messages().DefaultBackBtn, query)
}

func (h *BaseHandler) getMenuKeyboard(ctx *types.Context) *tgbotapi.InlineKeyboardMarkup {
	keyboard := tgbotapi.InlineKeyboardMarkup{}

	keyboard.InlineKeyboard = append(
		keyboard.InlineKeyboard,
		h.Helper.NewRow(ctx.Messages().MenuItemMyHeroes, queries.OpenMyHeroes),
	)

	return &keyboard
}

func (h *BaseHandler) GetIDFromString(str string) (int64, error) {
	IDInt64, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		h.Logger.Error(err)

		return 0, errors.ErrServerError
	}

	return IDInt64, nil
}

func (h *BaseHandler) HandleShowActiveFightOverview(ctx *types.Context) error {
	if err := h.Helper.SendTextMessage(ctx.GetChatID(), ctx.Messages().PleaseFinishActiveFight); err != nil {
		return err
	}

	fight, err := h.Services.TournamentFight.FindMyActive(ctx)
	if err != nil {
		return err
	}

	return h.SendFightOverviewMessage(ctx, fight)
}

func (h *BaseHandler) SendFightOverviewMessage(ctx *types.Context, fight *models.Fight, query ...*tgbotapi.CallbackQuery) error {
	msgText, err := ctx.Messages().FightInfo(fight)
	if err != nil {
		h.Logger.Error(err)

		return errors.ErrServerError
	}

	keyboard := &tgbotapi.InlineKeyboardMarkup{
		InlineKeyboard: [][]tgbotapi.InlineKeyboardButton{},
	}
	if !fight.HasEnded() {
		buttons := []telegram.KeyboardButton{
			{
				CallbackQuery: queries.FightActionSimpleStrike.WithID(fight.HeroID),
				Text:          ctx.Messages().FightActionSimpleStrike,
			},
			{
				CallbackQuery: queries.FightActionStrongStrike.WithID(fight.HeroID),
				Text:          ctx.Messages().FightActionStrongStrike, // @todo should decrease hero's armor
			},
			{
				CallbackQuery: queries.FightActionPreciseStrike.WithID(fight.HeroID),
				Text:          ctx.Messages().FightActionPreciseStrike,
			},
			{
				CallbackQuery: queries.OpenMyHero.WithID(fight.HeroID),
				Text:          ctx.Messages().FightActionRunAway, // @todo action should reduce double energy
			},
		}
		keyboard = h.Helper.CreateKeyboard(buttons, 1)
	}

	if len(query) > 0 {
		fightOverviewMsg := h.Helper.NewEditMessage(ctx.GetChatID(), query[0].Message.MessageID, msgText)
		fightOverviewMsg.ReplyMarkup = keyboard

		return h.Helper.Send(fightOverviewMsg)
	}

	fightOverviewMsg := h.Helper.NewMessage(ctx.GetChatID(), msgText)
	fightOverviewMsg.ReplyMarkup = keyboard

	return h.Helper.Send(fightOverviewMsg)
}

func (h *BaseHandler) OpenMyHero(ctx *types.Context, hero *models.Hero, query ...*tgbotapi.CallbackQuery) error {
	ended, err := h.endExistingActiveFightIfExist(ctx, hero, query...)
	if err != nil {
		return err
	}
	if ended {
		query = nil
	}

	heroInfo, err := ctx.Messages().GetHeroInfo(hero)
	if err != nil {
		h.Logger.Error(err)

		return errors.ErrServerError
	}

	keyboardButtons := make([]telegram.KeyboardButton, 0, 2)

	keyboardButtons = append(
		keyboardButtons,
		telegram.KeyboardButton{
			CallbackQuery: queries.StartTournamentFight.WithID(hero.ID),
			Text:          ctx.Messages().MenuItemStartTournamentFight,
		},
		telegram.KeyboardButton{
			CallbackQuery: queries.HeroDelete.WithID(hero.ID),
			Text:          ctx.Messages().MenuItemDeleteHero,
		},
		telegram.KeyboardButton{
			CallbackQuery: queries.OpenMyHeroes,
			Text:          ctx.Messages().DefaultBackBtn,
		},
	)
	keyboard := h.Helper.CreateKeyboard(keyboardButtons, 1)

	if len(query) > 0 {
		msg := h.Helper.NewEditMessage(
			ctx.GetChatID(),
			query[0].Message.MessageID,
			heroInfo,
		)
		msg.ReplyMarkup = keyboard

		return h.Helper.Send(msg)
	} else {
		msg := h.Helper.NewMessage(
			ctx.GetChatID(),
			heroInfo,
		)
		msg.ReplyMarkup = keyboard

		return h.Helper.Send(msg)
	}
}

func (h *BaseHandler) endExistingActiveFightIfExist(ctx *types.Context, hero *models.Hero, query ...*tgbotapi.CallbackQuery) (bool, error) {
	fight, err := h.Services.TournamentFight.FindActiveByHeroID(ctx, hero.ID)
	if err != nil {
		if errors.Is(err, errors.ErrNotFound) {
			return false, nil
		}

		return false, err
	}

	if err := h.Services.TournamentFight.RunAwayAsHero(ctx, fight.ID); err != nil {
		return false, err
	}

	msgText := ctx.Messages().FightRunAwaySuccess

	if len(query) > 0 {
		msg := h.Helper.NewEditMessage(ctx.GetChatID(), query[0].Message.MessageID, msgText)

		if err := h.Helper.Send(msg); err != nil {
			return false, err
		}
	} else {
		if err := h.Helper.SendTextMessage(ctx.GetChatID(), msgText); err != nil {
			return false, err
		}
	}

	return true, nil
}
