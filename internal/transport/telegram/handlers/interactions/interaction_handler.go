package interactions

import (
	"github.com/arthurshafikov/tg-gladiator/internal/transport/telegram/handlers"
)

type Handler struct {
	*handlers.BaseHandler
}

func NewHandler(baseHandler *handlers.BaseHandler) *Handler {
	return &Handler{
		BaseHandler: baseHandler,
	}
}
