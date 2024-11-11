package commands

import (
	"github.com/arthurshafikov/tg-gladiator/internal/core/models"
	"github.com/arthurshafikov/tg-gladiator/internal/core/types"
	"github.com/arthurshafikov/tg-gladiator/internal/transport/telegram/handlers"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type Handler struct {
	*handlers.BaseHandler
}

func NewHandler(baseHandler *handlers.BaseHandler) *Handler {
	return &Handler{
		BaseHandler: baseHandler,
	}
}

func (h *Handler) HandleStart(ctx *types.Context, message *tgbotapi.Message) error {
	data := models.Chat{
		ChatID: message.Chat.ID,
	}

	if message.From != nil {
		data.Name = message.From.FirstName + " " + message.From.LastName
		data.Username = message.From.UserName
	}

	chat, _, err := h.BaseHandler.Services.Chat.FirstOrCreate(ctx.GetContext(), data)
	if err != nil {
		return err
	}
	ctx.SetChat(chat)

	msg := h.Helper.NewMessage(
		ctx.GetChatID(),
		ctx.Messages().StartSuccess,
	)

	return h.Helper.Send(msg)
}

func (h *Handler) HandleHelp(ctx *types.Context) error {
	msg := h.Helper.NewMessage(
		ctx.GetChatID(),
		ctx.Messages().Help,
	)

	return h.Helper.Send(msg)
}
