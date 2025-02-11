package middlewares

import (
	"slices"

	"github.com/arthurshafikov/tg-gladiator/internal/core/constants/queries"
	"github.com/arthurshafikov/tg-gladiator/internal/core/errors"
	"github.com/arthurshafikov/tg-gladiator/internal/core/types"
	"github.com/arthurshafikov/tg-gladiator/internal/services"
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

type QueryMiddlewareChain struct {
	services *services.Services
}

func NewQueryMiddlewareChain(services *services.Services) *QueryMiddlewareChain {
	return &QueryMiddlewareChain{
		services: services,
	}
}

func (m *QueryMiddlewareChain) Chain(middlewares ...QueryMiddlewareFunc) QueryMiddlewareFunc {
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
				return m.Chain(rest...)(ctx, query, payloadQuery, payload, next)
			},
		)
	}
}

func (m *QueryMiddlewareChain) QueryCheckDontHaveActiveFight(
	ctx *types.Context,
	query *tgbotapi.CallbackQuery,
	payloadQuery queries.Query,
	payload []string,
	next QueryHandlerFunc,
) error {
	if !slices.Contains[[]queries.Query, queries.Query](queries.FightQueries(), payloadQuery) {
		activeFight, err := m.services.Fight.FindMyActive(ctx)
		if err != nil && !errors.Is(err, errors.ErrNotFound) {
			return err
		}

		if activeFight != nil {
			return errors.ErrHasActiveFight
		}
	}

	return next(ctx, query, payload)
}
