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
	default:
		return errors.ErrUndefinedCommand
	}

	return nil
}

//nolint:funlen,gocyclo
func (b *Bot) handleCallbackQuery(ctx *types.Context, query *tgbotapi.CallbackQuery) error {
	splittedData := strings.Split(query.Data, queries.SpecialDelimeterInQueryCallback)

	switch queries.Query(splittedData[0]) {
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

	switch interactions.Interaction(splittedData[0]) {
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
