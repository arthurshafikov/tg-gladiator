package telegram

import "github.com/arthurshafikov/tg-gladiator/internal/core/constants/queries"

const (
	TelegramMarkDownParseMode = "markdown"
	TelegramHTMLParseMode     = "html"

	TelegramChatTypePrivate    = "private"
	TelegramChatTypeGroup      = "group"
	TelegramChatTypeSupergroup = "supergroup"
	TelegramChatTypeChannel    = "channel"
)

type KeyboardButton struct {
	CallbackQuery queries.Query
	Text          string
}
