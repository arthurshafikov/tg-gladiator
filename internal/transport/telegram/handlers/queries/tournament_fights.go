package queries

import (
	"fmt"

	"github.com/arthurshafikov/tg-gladiator/internal/core/constants/enums"
	"github.com/arthurshafikov/tg-gladiator/internal/core/constants/queries"
	"github.com/arthurshafikov/tg-gladiator/internal/core/constants/telegram"
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

	return h.SendFightOverviewMessage(ctx, fight, query)
}

func (h *Handler) HandleFightActionSimpleStrike(ctx *types.Context, query *tgbotapi.CallbackQuery, payload []string) error {
	return h.tournamentTurn(ctx, query, payload, enums.FightActionSimpleStrike)
}

func (h *Handler) HandleFightActionStrongStrike(ctx *types.Context, query *tgbotapi.CallbackQuery, payload []string) error {
	return h.tournamentTurn(ctx, query, payload, enums.FightActionStrongStrike)
}

func (h *Handler) HandleFightActionPreciseStrike(ctx *types.Context, query *tgbotapi.CallbackQuery, payload []string) error {
	return h.tournamentTurn(ctx, query, payload, enums.FightActionPreciseStrike)
}

func (h *Handler) tournamentTurn(
	ctx *types.Context,
	query *tgbotapi.CallbackQuery,
	payload []string,
	fightAction enums.FightActionType,
) error {
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

	fight, fightEvents, err := h.Services.TournamentFight.MakeTurn(ctx, fight, fightAction)
	if err != nil {
		return err
	}

	fightEventsMsgText := ctx.Messages().FightEvents(fight, fightEvents)
	msg := h.Helper.NewEditMessage(ctx.GetChatID(), query.Message.MessageID, fightEventsMsgText)
	if err := h.Helper.Send(msg); err != nil {
		return err
	}

	if err := h.SendFightOverviewMessage(ctx, fight); err != nil {
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
