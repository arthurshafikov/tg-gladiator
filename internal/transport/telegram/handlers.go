package telegram

import (
	"strings"

	"github.com/arthurshafikov/tg-gladiator/internal/core/constants/commands"
	"github.com/arthurshafikov/tg-gladiator/internal/core/constants/interactions"
	"github.com/arthurshafikov/tg-gladiator/internal/core/constants/queries"
	"github.com/arthurshafikov/tg-gladiator/internal/core/errors"
	"github.com/arthurshafikov/tg-gladiator/internal/core/types"
	"github.com/arthurshafikov/tg-gladiator/internal/transport/telegram/middlewares"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (b *Bot) handleCommand(ctx *types.Context, message *tgbotapi.Message) error {
	switch message.Command() {
	case commands.Start:
		return b.commandHandler.HandleStart(ctx, message)
	case commands.Help:
		return b.commandHandler.HandleHelp(ctx)
	default:
		return errors.ErrUndefinedCommand
	}
}

//nolint:funlen,gocyclo
func (b *Bot) handleCallbackQuery(ctx *types.Context, query *tgbotapi.CallbackQuery) error {
	splittedData := strings.Split(query.Data, queries.SpecialDelimeterInQueryCallback)

	payloadQuery := queries.Query(splittedData[0])
	payload := []string{}
	if len(splittedData) > 1 {
		payload = splittedData[1:]
	}

	middlewareChain := middlewares.NewQueryMiddlewareChain(
		payloadQuery,
		middlewares.QueryCheckIsActiveAccount,
	)

	switch payloadQuery {
	case queries.OpenMenu:
		return middlewareChain(ctx, query, payloadQuery, payload, b.queryHandler.HandleOpenMenu)
	case queries.OpenMyHeroes:
		return middlewareChain(ctx, query, payloadQuery, payload, b.queryHandler.HandleOpenMyHeroes)
	case queries.OpenMyHero:
		return middlewareChain(ctx, query, payloadQuery, payload, b.queryHandler.HandleOpenMyHero)

	case queries.HeroCreationStart:
		return middlewareChain(ctx, query, payloadQuery, payload, b.queryHandler.HandleHeroCreationStart)
	case queries.HeroCreationSelectClass:
		return middlewareChain(ctx, query, payloadQuery, payload, b.queryHandler.HandleHeroCreationSelectClass)

	case queries.StartTournamentFight:
		return middlewareChain(ctx, query, payloadQuery, payload, b.queryHandler.HandleStartTournamentFight)
	case queries.FightActionSimpleStrike:
		return middlewareChain(ctx, query, payloadQuery, payload, b.queryHandler.HandleFightActionSimpleStrike)
	case queries.FightActionStrongStrike:
		return middlewareChain(ctx, query, payloadQuery, payload, b.queryHandler.HandleFightActionStrongStrike)
	case queries.FightActionPreciseStrike:
		return middlewareChain(ctx, query, payloadQuery, payload, b.queryHandler.HandleFightActionPreciseStrike)

	case queries.HeroDelete:
		return middlewareChain(ctx, query, payloadQuery, payload, b.queryHandler.HandleHeroDelete)

	default:
		return errors.ErrUndefinedCallbackQuery
	}
}

func (b *Bot) handleInteraction(
	ctx *types.Context,
	interaction *interactions.Interaction,
	message *tgbotapi.Message,
) error {
	splittedData := strings.Split(string(*interaction), interactions.SpecialDelimeterForInteractions)

	payloadInteraction := interactions.Interaction(splittedData[0])
	payload := []string{}
	if len(splittedData) > 1 {
		payload = splittedData[1:]
	}

	middlewareChain := middlewares.NewInteractionMiddlewareChain(
		middlewares.InteractionLogSomething,
	)

	switch payloadInteraction {
	case interactions.HeroCreationEnterName:
		return middlewareChain(ctx, message, payloadInteraction, payload, b.interactionHandler.HandleHeroCreationEnterName)

	case interactions.HeroDelete:
		return middlewareChain(ctx, message, payloadInteraction, payload, b.interactionHandler.HandleHeroDelete)
	default:
		return errors.ErrUndefinedInteraction
	}
}

func (b *Bot) handleMessage(ctx *types.Context) error {
	return nil
}

func (b *Bot) handleChangedUsername(ctx *types.Context, username string) {
	if err := b.services.Chat.UpdateUsername(ctx, username); err != nil {
		b.logger.Error(err)
	}
}
