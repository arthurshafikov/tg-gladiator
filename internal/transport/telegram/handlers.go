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
	"github.com/sirupsen/logrus"
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

	middleware := middlewares.NewQueryMiddlewareChain(
		b.services,
	)

	middlewareChain := middleware.Chain(
		middleware.QueryCheckDontHaveActiveFight,
	)

	var err error
	switch payloadQuery {
	case queries.OpenMenu:
		err = middlewareChain(ctx, query, payloadQuery, payload, b.queryHandler.HandleOpenMenu)
	case queries.OpenMyHeroes:
		err = middlewareChain(ctx, query, payloadQuery, payload, b.queryHandler.HandleOpenMyHeroes)
	case queries.OpenMyHero:
		err = middlewareChain(ctx, query, payloadQuery, payload, b.queryHandler.HandleOpenMyHero)

	case queries.HeroCreationStart:
		err = middlewareChain(ctx, query, payloadQuery, payload, b.queryHandler.HandleHeroCreationStart)
	case queries.HeroCreationSelectClass:
		err = middlewareChain(ctx, query, payloadQuery, payload, b.queryHandler.HandleHeroCreationSelectClass)

	case queries.StartTournamentFight:
		err = middlewareChain(ctx, query, payloadQuery, payload, b.queryHandler.HandleStartTournamentFight)
	case queries.FightActionSimpleStrike:
		err = middlewareChain(ctx, query, payloadQuery, payload, b.queryHandler.HandleFightActionSimpleStrike)
	case queries.FightActionStrongStrike:
		err = middlewareChain(ctx, query, payloadQuery, payload, b.queryHandler.HandleFightActionStrongStrike)
	case queries.FightActionPreciseStrike:
		err = middlewareChain(ctx, query, payloadQuery, payload, b.queryHandler.HandleFightActionPreciseStrike)
	case queries.FightActionRunAway:
		err = middlewareChain(ctx, query, payloadQuery, payload, b.queryHandler.HandleFightActionRunAway)

	case queries.OpenShop:
		err = middlewareChain(ctx, query, payloadQuery, payload, b.queryHandler.HandleOpenShop)

	case queries.HeroDelete:
		err = middlewareChain(ctx, query, payloadQuery, payload, b.queryHandler.HandleHeroDelete)

	default:
		err = errors.ErrUndefinedCallbackQuery
	}

	if errors.Is(err, errors.ErrHasActiveFight) {
		if err := b.helper.AnswerCallbackQueryNil(query.ID); err != nil {
			logrus.Error(err)
		}

		err = b.queryHandler.HandleShowActiveFightOverview(ctx)
	}

	return err
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

	var err error
	switch payloadInteraction {
	case interactions.HeroCreationEnterName:
		err = middlewareChain(ctx, message, payloadInteraction, payload, b.interactionHandler.HandleHeroCreationEnterName)

	case interactions.HeroDelete:
		err = middlewareChain(ctx, message, payloadInteraction, payload, b.interactionHandler.HandleHeroDelete)
	default:
		err = errors.ErrUndefinedInteraction
	}

	if errors.Is(err, errors.ErrHasActiveFight) {
		err = b.interactionHandler.HandleShowActiveFightOverview(ctx)
	}

	return err
}

func (b *Bot) handleMessage(ctx *types.Context) error {
	return nil
}

func (b *Bot) handleChangedUsername(ctx *types.Context, username string) {
	if err := b.services.Chat.UpdateUsername(ctx, username); err != nil {
		b.logger.Error(err)
	}
}
