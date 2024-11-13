package handlers

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/arthurshafikov/tg-gladiator/internal/config"
	"github.com/arthurshafikov/tg-gladiator/internal/core/constants/queries"
	"github.com/arthurshafikov/tg-gladiator/internal/core/errors"
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
