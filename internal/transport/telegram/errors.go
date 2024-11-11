package telegram

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/arthurshafikov/tg-gladiator/internal/config"
	"github.com/arthurshafikov/tg-gladiator/internal/core/constants/telegram"
	"github.com/arthurshafikov/tg-gladiator/internal/core/errors"
)

func (b *Bot) handleError(chatID int64, err error, messages *config.Messages) {
	msg := b.helper.NewMessage(chatID, b.getErrorText(err, messages))
	msg.ParseMode = telegram.TelegramMarkDownParseMode
	if err := b.helper.SendMessageWithoutKeyboard(msg); err != nil {
		b.logger.Error(err)
	}
}

func (b *Bot) handleCallbackQueryError(queryID string, chatID int64, err error, messages *config.Messages) {
	if err := b.helper.AnswerCallbackQuery(b.helper.NewCallback(queryID, b.getErrorText(err, messages))); err != nil {
		b.logger.Error(fmt.Errorf("error in AnswerCallbackQuery chat %v: %w", chatID, err))
	}
}

func (b *Bot) getErrorText(err error, messages *config.Messages) string {
	var text string
	if strings.Contains(err.Error(), "too_long") {
		_, after, found := strings.Cut(err.Error(), ":")
		if !found {
			b.logger.Error(fmt.Errorf("symbol ':' not found in too_long error message: %w", err))

			return messages.Errors[errors.ErrServerError.Error()]
		}

		maximumSymbols, err := strconv.Atoi(after)
		if err != nil {
			b.logger.Error(fmt.Errorf("could not parse int from too_long error message: %w", err))

			return messages.Errors[errors.ErrServerError.Error()]
		}

		text = fmt.Sprintf(messages.Errors["too_long"], maximumSymbols)
	} else {
		text = messages.Errors[err.Error()]
	}

	return text
}
