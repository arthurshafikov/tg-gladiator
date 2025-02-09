package middlewares

import (
	"github.com/arthurshafikov/tg-gladiator/internal/core/constants/interactions"
	"github.com/arthurshafikov/tg-gladiator/internal/core/types"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type InteractionHandlerFunc func(ctx *types.Context, message *tgbotapi.Message, payload []string) error

type InteractionMiddlewareFunc func(
	ctx *types.Context,
	message *tgbotapi.Message,
	payloadInteraction interactions.Interaction,
	payload []string,
	next InteractionHandlerFunc,
) error

func NewInteractionMiddlewareChain(middlewares ...InteractionMiddlewareFunc) InteractionMiddlewareFunc {
	return func(
		ctx *types.Context,
		message *tgbotapi.Message,
		payloadInteraction interactions.Interaction,
		payload []string,
		next InteractionHandlerFunc,
	) error {
		if len(middlewares) == 0 {
			return next(ctx, message, payload)
		}

		current := middlewares[0]
		rest := middlewares[1:]

		return current(
			ctx,
			message,
			payloadInteraction,
			payload,
			func(
				ctx *types.Context,
				message *tgbotapi.Message,
				payload []string,
			) error {
				return NewInteractionMiddlewareChain(rest...)(ctx, message, payloadInteraction, payload, next)
			},
		)
	}
}

func InteractionLogSomething(
	ctx *types.Context,
	message *tgbotapi.Message,
	payloadInteraction interactions.Interaction,
	payload []string,
	next InteractionHandlerFunc,
) error {
	return next(ctx, message, payload)
}
