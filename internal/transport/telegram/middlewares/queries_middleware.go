package middlewares

import (
	"github.com/arthurshafikov/tg-gladiator/internal/core/constants/queries"
	"github.com/arthurshafikov/tg-gladiator/internal/core/types"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type QueryHandlerFunc func(ctx *types.Context, query *tgbotapi.CallbackQuery, payload []string) error

type QueryMiddlewareFunc func(
	ctx *types.Context,
	query *tgbotapi.CallbackQuery,
	payloadQuery queries.Query,
	payload []string,
	next QueryHandlerFunc,
) error

func NewQueryMiddlewareChain(payloadQuery queries.Query, middlewares ...QueryMiddlewareFunc) QueryMiddlewareFunc {
	return func(
		ctx *types.Context,
		query *tgbotapi.CallbackQuery,
		payloadQuery queries.Query,
		payload []string,
		next QueryHandlerFunc,
	) error {
		if len(middlewares) == 0 {
			return next(ctx,
				query, payload)
		}

		current := middlewares[0]
		rest := middlewares[1:]

		return current(
			ctx,
			query,
			payloadQuery,
			payload,
			func(ctx *types.Context, query *tgbotapi.CallbackQuery, payload []string) error {
				return NewQueryMiddlewareChain(payloadQuery, rest...)(ctx, query, payloadQuery, payload, next)
			},
		)
	}
}

func QueryCheckIsActiveAccount(
	ctx *types.Context,
	query *tgbotapi.CallbackQuery,
	payloadQuery queries.Query,
	payload []string,
	next QueryHandlerFunc,
) error {
	return next(ctx, query, payload)
}
