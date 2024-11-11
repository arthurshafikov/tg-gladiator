package handlers

import (
	"github.com/arthurshafikov/tg-gladiator/internal/config"
	"github.com/arthurshafikov/tg-gladiator/internal/core/constants/commands"
	"github.com/arthurshafikov/tg-gladiator/internal/core/constants/queries"
	"github.com/arthurshafikov/tg-gladiator/internal/core/constants/telegram"
	"github.com/arthurshafikov/tg-gladiator/internal/core/errors"
	"github.com/arthurshafikov/tg-gladiator/internal/services"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

const DefaultTelegramParseMode = telegram.TelegramHTMLParseMode

type TelegramHandlerHelper interface {
	NewCallback(id string, text string) tgbotapi.CallbackConfig
	NewMessage(chatID int64, text string) tgbotapi.MessageConfig
	NewEditMessageText(chatID int64, messageID int, text string) tgbotapi.EditMessageTextConfig
	NewEditMessageReplyMarkup(
		chatID int64,
		messageID int,
		replyMarkup *tgbotapi.InlineKeyboardMarkup,
	) tgbotapi.EditMessageReplyMarkupConfig
	Send(msg tgbotapi.Chattable) error
	SendMessage(msg tgbotapi.MessageConfig) error
	SendTextMessage(chatID int64, text string) error
	SendMessageWithKeyboard(msg tgbotapi.MessageConfig, replyCommands ...string) error
	SendMessageWithoutKeyboard(msg tgbotapi.MessageConfig) error
	AnswerCallbackQuery(callback tgbotapi.CallbackConfig) error
	AnswerCallbackQueryNil(queryID string) error
	SendInvoice(msg tgbotapi.Chattable) error
	AnswerPreCheckoutQuery(config tgbotapi.PreCheckoutConfig) error
	NewUpdateChannel(offset int) tgbotapi.UpdateConfig
	GetUpdatesChan(config tgbotapi.UpdateConfig) (tgbotapi.UpdatesChannel, error)
	NewRow(btnName string, query queries.Query) []tgbotapi.InlineKeyboardButton
	NewRowButton(btnName string, query queries.Query) tgbotapi.InlineKeyboardButton
	NewDeleteMessage(chatID int64, messageID int) tgbotapi.DeleteMessageConfig
	CreateKeyboard(buttons []telegram.KeyboardButton, perRow int) *tgbotapi.InlineKeyboardMarkup
}

type HandlerParams struct {
	Services *services.Services
	Logger   services.Logger
	Config   *config.Config

	Helper TelegramHandlerHelper
}

type Helper struct {
	logger services.Logger
	bot    *tgbotapi.BotAPI
	config *config.Config
}

func NewHelper(logger services.Logger, bot *tgbotapi.BotAPI, config *config.Config) *Helper {
	return &Helper{
		logger: logger,
		bot:    bot,
		config: config,
	}
}

func (h *Helper) NewMessage(chatID int64, text string) tgbotapi.MessageConfig {
	return tgbotapi.NewMessage(chatID, text)
}

func (h *Helper) NewCallback(id string, text string) tgbotapi.CallbackConfig {
	return tgbotapi.NewCallback(id, text)
}

func (h *Helper) NewEditMessageText(chatID int64, messageID int, text string) tgbotapi.EditMessageTextConfig {
	return tgbotapi.NewEditMessageText(chatID, messageID, text)
}

func (h *Helper) NewEditMessageReplyMarkup(
	chatID int64,
	messageID int,
	replyMarkup *tgbotapi.InlineKeyboardMarkup,
) tgbotapi.EditMessageReplyMarkupConfig {
	return tgbotapi.NewEditMessageReplyMarkup(chatID, messageID, *replyMarkup)
}

func (h *Helper) Send(msg tgbotapi.Chattable) error {
	if newMsg, ok := msg.(tgbotapi.EditMessageTextConfig); ok {
		newMsg.ParseMode = DefaultTelegramParseMode
		msg = newMsg
	}
	if newMsg, ok := msg.(tgbotapi.MessageConfig); ok {
		newMsg.ParseMode = DefaultTelegramParseMode
		msg = newMsg
	}

	if _, err := h.bot.Send(msg); err != nil {
		h.logger.Error(err)

		return errors.ErrServerError
	}

	return nil
}

func (h *Helper) AnswerCallbackQuery(callback tgbotapi.CallbackConfig) error {
	if _, err := h.bot.AnswerCallbackQuery(callback); err != nil {
		h.logger.Error(err)

		return errors.ErrServerError
	}

	return nil
}

func (h *Helper) AnswerCallbackQueryNil(queryID string) error {
	_, err := h.bot.AnswerCallbackQuery(h.NewCallback(queryID, ""))

	return err
}

func (h *Helper) SendMessage(msg tgbotapi.MessageConfig) error {
	if msg.ParseMode == "" {
		msg.ParseMode = DefaultTelegramParseMode
	}
	if msg.ReplyMarkup == nil {
		msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(false)
	}
	msg.DisableWebPagePreview = true
	_, err := h.bot.Send(msg)

	return err
}

func (h *Helper) SendTextMessage(chatID int64, text string) error {
	msg := h.NewMessage(chatID, text)
	msg.ReplyMarkup = nil

	return h.SendMessage(msg)
}

func (h *Helper) SendMessageWithKeyboard(msg tgbotapi.MessageConfig, replyCommands ...string) error {
	if msg.ReplyMarkup == nil {
		var replyKeyboard [][]tgbotapi.KeyboardButton
		if len(replyCommands) == 0 {
			replyKeyboard = append(
				replyKeyboard,
				tgbotapi.NewKeyboardButtonRow(
					tgbotapi.NewKeyboardButton("/"+commands.Help),
				),
			)
		} else {
			var buttons []tgbotapi.KeyboardButton
			for _, replyCommand := range replyCommands {
				buttons = append(buttons, tgbotapi.NewKeyboardButton("/"+replyCommand))
			}

			replyKeyboard = append(
				replyKeyboard,
				tgbotapi.NewKeyboardButtonRow(buttons...),
			)
		}
		msg.ReplyMarkup = tgbotapi.ReplyKeyboardMarkup{
			Keyboard:        replyKeyboard,
			ResizeKeyboard:  true,
			OneTimeKeyboard: true,
		}
	}

	return h.SendMessage(msg)
}

func (h *Helper) SendMessageWithoutKeyboard(msg tgbotapi.MessageConfig) error {
	msg.ReplyMarkup = tgbotapi.ReplyKeyboardRemove{
		RemoveKeyboard: true,
	}

	return h.SendMessage(msg)
}

func (h *Helper) SendInvoice(msg tgbotapi.Chattable) error {
	if _, err := h.bot.Send(msg); err != nil {
		h.logger.Error(err)

		return errors.ErrServerError
	}

	return nil
}

func (h *Helper) AnswerPreCheckoutQuery(config tgbotapi.PreCheckoutConfig) error {
	if _, err := h.bot.AnswerPreCheckoutQuery(config); err != nil {
		h.logger.Error(err)

		return errors.ErrServerError
	}

	return nil
}

func (h *Helper) NewUpdateChannel(offset int) tgbotapi.UpdateConfig {
	return tgbotapi.NewUpdate(offset)
}

func (h *Helper) GetUpdatesChan(config tgbotapi.UpdateConfig) (tgbotapi.UpdatesChannel, error) {
	return h.bot.GetUpdatesChan(config)
}

func (h *Helper) NewRow(btnTitle string, query queries.Query) []tgbotapi.InlineKeyboardButton {
	return []tgbotapi.InlineKeyboardButton{
		h.NewRowButton(btnTitle, query),
	}
}

func (h *Helper) NewRowButton(btnTitle string, query queries.Query) tgbotapi.InlineKeyboardButton {
	return tgbotapi.NewInlineKeyboardButtonData(btnTitle, string(query))
}

func (h *Helper) CreateKeyboard(buttons []telegram.KeyboardButton, perRow int) *tgbotapi.InlineKeyboardMarkup {
	keyboard := tgbotapi.InlineKeyboardMarkup{
		InlineKeyboard: [][]tgbotapi.InlineKeyboardButton{},
	}

	activeRow := []tgbotapi.InlineKeyboardButton{}
	for i, button := range buttons {
		if perRow != -1 && i%perRow == 0 {
			keyboard.InlineKeyboard = append(keyboard.InlineKeyboard, activeRow)

			activeRow = []tgbotapi.InlineKeyboardButton{}
		}

		activeRow = append(activeRow, h.NewRowButton(
			button.Text,
			button.CallbackQuery,
		))
	}
	keyboard.InlineKeyboard = append(keyboard.InlineKeyboard, activeRow)

	return &keyboard
}

func (h *Helper) NewDeleteMessage(chatID int64, messageID int) tgbotapi.DeleteMessageConfig {
	return tgbotapi.NewDeleteMessage(chatID, messageID)
}
