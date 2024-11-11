package queries

import (
	"github.com/arthurshafikov/tg-gladiator/internal/transport/telegram/handlers"
)

type QueryHandler struct {
	*handlers.BaseHandler
}

func NewQueryHandler(baseHandler *handlers.BaseHandler) *QueryHandler {
	return &QueryHandler{
		BaseHandler: baseHandler,
	}
}
