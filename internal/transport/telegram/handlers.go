package telegram

import (
	"strings"

	"github.com/arthurshafikov/tg-gladiator/internal/core/constants/commands"
	"github.com/arthurshafikov/tg-gladiator/internal/core/constants/interactions"
	"github.com/arthurshafikov/tg-gladiator/internal/core/constants/queries"
	"github.com/arthurshafikov/tg-gladiator/internal/core/errors"
	"github.com/arthurshafikov/tg-gladiator/internal/core/types"
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

	payload := splittedData[1:]

	switch queries.Query(splittedData[0]) {
	case queries.OpenMenu:
		return b.queryHandler.HandleOpenMenu(ctx, query)
	case queries.OpenMyHeroes:
		return b.queryHandler.HandleOpenMyHeroes(ctx, query)
	case queries.OpenMyHero:
		return b.queryHandler.HandleOpenMyHero(ctx, query, payload)

	case queries.HeroCreationStart:
		return b.queryHandler.HandleHeroCreationStart(ctx, query)
	case queries.HeroCreationSelectClass:
		return b.queryHandler.HandleHeroCreationSelectClass(ctx, query, payload)

	case queries.StartTournamentFight:
		return b.queryHandler.HandleStartTournamentFight(ctx, query, payload)
	case queries.FightActionSimpleStrike:
		return b.queryHandler.HandleFightActionSimpleStrike(ctx, query, payload)
	case queries.FightActionStrongStrike:
		return b.queryHandler.HandleFightActionStrongStrike(ctx, query, payload)
	case queries.FightActionPreciseStrike:
		return b.queryHandler.HandleFightActionPreciseStrike(ctx, query, payload)

	case queries.HeroDelete:
		return b.queryHandler.HandleHeroDelete(ctx, query, payload)

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

	payload := splittedData[1:]

	switch interactions.Interaction(splittedData[0]) {
	case interactions.HeroCreationEnterName:
		return b.interactionHandler.HandleHeroCreationEnterName(ctx, message, payload)

	case interactions.HeroDelete:
		return b.interactionHandler.HandleHeroDelete(ctx, message, payload)
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
