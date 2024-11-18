package queries

import (
	"fmt"

	"github.com/arthurshafikov/tg-gladiator/internal/core/constants/enums"
	"github.com/arthurshafikov/tg-gladiator/internal/core/constants/queries"
	"github.com/arthurshafikov/tg-gladiator/internal/core/constants/telegram"
	"github.com/arthurshafikov/tg-gladiator/internal/core/errors"
	"github.com/arthurshafikov/tg-gladiator/internal/core/models"
	"github.com/arthurshafikov/tg-gladiator/internal/core/types"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (h *Handler) HandleStartTournamentFight(ctx *types.Context, query *tgbotapi.CallbackQuery, payload []string) error {
	if err := h.ValidatePayloadLength(payload, 1); err != nil {
		return err
	}

	heroID, err := h.GetIDFromString(payload[0])
	if err != nil {
		return err
	}

	fight, err := h.Services.TournamentFight.Create(ctx, heroID)
	if err != nil {
		return err
	}

	return h.sendFightOverviewMessage(ctx, fight, query)
}

func (h *Handler) HandleFightActionPunch(ctx *types.Context, query *tgbotapi.CallbackQuery, payload []string) error {
	if err := h.ValidatePayloadLength(payload, 1); err != nil {
		return err
	}

	heroID, err := h.GetIDFromString(payload[0])
	if err != nil {
		return err
	}

	fight, err := h.Services.TournamentFight.FindActiveByHeroID(ctx, heroID)
	if err != nil {
		return err
	}

	fight, fightEvents, err := h.Services.TournamentFight.MakeTurn(ctx, fight, enums.FightActionPunch)
	if err != nil {
		return err
	}

	fightEventsMsgText := ctx.Messages().FightEvents(fight, fightEvents)
	msg := h.Helper.NewEditMessage(ctx.GetChatID(), query.Message.MessageID, fightEventsMsgText)
	if err := h.Helper.Send(msg); err != nil {
		return err
	}

	if err := h.sendFightOverviewMessage(ctx, fight); err != nil {
		return err
	}

	if fight.HasEnded() {
		var msgText string
		if fight.IsHeroWon() {
			msgText = fmt.Sprintf(
				ctx.Messages().FightResultHeroWon,
				fight.GoldReward,
			)
		} else {
			msgText = ctx.Messages().FightResultOpponentWon
		}

		msg := h.Helper.NewMessage(ctx.GetChatID(), msgText)
		msg.ReplyMarkup = h.Helper.CreateKeyboard([]telegram.KeyboardButton{
			{
				CallbackQuery: queries.OpenMyHero.WithID(heroID),
				Text:          ctx.Messages().FightResultBackButton,
			},
		}, 1)

		return h.Helper.Send(msg)
	}

	return nil
}

func (h *Handler) sendFightOverviewMessage(ctx *types.Context, fight *models.Fight, query ...*tgbotapi.CallbackQuery) error {
	msgText, err := ctx.Messages().FightInfo(fight)
	if err != nil {
		h.Logger.Error(err)

		return errors.ErrServerError
	}

	keyboard := &tgbotapi.InlineKeyboardMarkup{
		InlineKeyboard: [][]tgbotapi.InlineKeyboardButton{},
	}
	if !fight.HasEnded() {
		buttons := []telegram.KeyboardButton{
			{
				CallbackQuery: queries.FightActionPunch.WithID(fight.HeroID),
				Text:          ctx.Messages().FightActionPunch,
			},
			{
				CallbackQuery: queries.OpenMyHero.WithID(fight.HeroID),
				Text:          ctx.Messages().FightActionRunAway, // @todo action should reduce energy
			},
		}
		keyboard = h.Helper.CreateKeyboard(buttons, 1)
	}

	if len(query) > 0 {
		fightOverviewMsg := h.Helper.NewEditMessage(ctx.GetChatID(), query[0].Message.MessageID, msgText)
		fightOverviewMsg.ReplyMarkup = keyboard

		return h.Helper.Send(fightOverviewMsg)
	}

	fightOverviewMsg := h.Helper.NewMessage(ctx.GetChatID(), msgText)
	fightOverviewMsg.ReplyMarkup = keyboard

	return h.Helper.Send(fightOverviewMsg)
}
