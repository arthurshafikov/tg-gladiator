package types

import (
	"context"

	"github.com/arthurshafikov/tg-gladiator/internal/config"
	"github.com/arthurshafikov/tg-gladiator/internal/core/models"
)

type Context struct {
	ctx      context.Context
	chat     *models.Chat
	messages *config.Messages
}

func NewContext(ctx context.Context, chat *models.Chat, messages *config.Messages) *Context {
	return &Context{
		ctx:      ctx,
		chat:     chat,
		messages: messages,
	}
}

func (c *Context) GetContext() context.Context {
	return c.ctx
}

func (c *Context) GetChat() *models.Chat {
	return c.chat
}

func (c *Context) GetChatID() int64 {
	return c.chat.ChatID
}

func (c *Context) SetChat(chat *models.Chat) *Context {
	c.chat = chat

	return c
}

func (c *Context) Messages() *config.Messages {
	return c.messages
}
