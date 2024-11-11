package handlers

import (
	"fmt"
	"strings"

	"github.com/arthurshafikov/tg-gladiator/internal/config"
	"github.com/arthurshafikov/tg-gladiator/internal/core/constants/queries"
	"github.com/arthurshafikov/tg-gladiator/internal/core/errors"
	"github.com/arthurshafikov/tg-gladiator/internal/services"
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

func (h *BaseHandler) ValidateDataLength(data []string, expectedMinLength int) error {
	if len(data) < expectedMinLength {
		h.Logger.Error(fmt.Sprintf(
			"Data Len is not %v, data: %s",
			expectedMinLength,
			strings.Join(data, queries.SpecialDelimeterInQueryCallback),
		))

		return errors.ErrServerError
	}

	return nil
}
