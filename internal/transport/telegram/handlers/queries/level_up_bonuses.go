package queries

import (
	"fmt"

	"github.com/arthurshafikov/tg-gladiator/internal/core/constants/enums"
	"github.com/arthurshafikov/tg-gladiator/internal/core/constants/queries"
	"github.com/arthurshafikov/tg-gladiator/internal/core/constants/telegram"
	"github.com/arthurshafikov/tg-gladiator/internal/core/types"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (h *Handler) HandleOpenLevelUpBonusesOverview(ctx *types.Context, query *tgbotapi.CallbackQuery, payload []string) error {
	if err := h.ValidatePayloadLength(payload, 1); err != nil {
		return err
	}

	heroID, err := h.GetIDFromString(payload[0])
	if err != nil {
		return err
	}

	hero, err := h.Services.Hero.FindMy(ctx, heroID)
	if err != nil {
		return err
	}

	// h.Config.App @todo take level up values from config

	keyboard := h.Helper.CreateKeyboard([]telegram.KeyboardButton{
		{
			CallbackQuery: queries.SpendLevelUpBonus.WithID(hero.ID).With(enums.StatsUpgradeHealth),
			Text: fmt.Sprintf(
				"%s +%v (%v)",
				ctx.Messages().HP,
				5,
				hero.GetHP(),
			),
		},
		{
			CallbackQuery: queries.SpendLevelUpBonus.WithID(hero.ID).With(enums.StatsUpgradeAttack),
			Text: fmt.Sprintf(
				"%s +%v (%v)",
				ctx.Messages().Attack,
				1,
				hero.GetMinAttack(),
			),
		},
		{
			CallbackQuery: queries.SpendLevelUpBonus.WithID(hero.ID).With(enums.StatsUpgradeDefence),
			Text: fmt.Sprintf(
				"%s +%v (%v)",
				ctx.Messages().Defense,
				1,
				hero.GetDefense(),
			),
		},
		{
			CallbackQuery: queries.SpendLevelUpBonus.WithID(hero.ID).With(enums.StatsUpgradeEvasionChance),
			Text: fmt.Sprintf(
				"%s +%v%% (%v)",
				ctx.Messages().EvasionChancePercent,
				1,
				hero.GetEvasionChancePercent(),
			),
		},
		{
			CallbackQuery: queries.SpendLevelUpBonus.WithID(hero.ID).With(enums.StatsUpgradeCriticalChance),
			Text: fmt.Sprintf(
				"%s +%v%% (%v)",
				ctx.Messages().CriticalChancePercent,
				1,
				hero.GetCriticalChancePercent(),
			),
		},
	}, 1)

	keyboard.InlineKeyboard = append(
		keyboard.InlineKeyboard,
		h.GetBackButtonKeyboardRow(ctx, queries.OpenMyHero.WithID(hero.ID)),
	)

	msg := h.Helper.NewEditMessage(
		ctx.GetChatID(),
		query.Message.MessageID,
		fmt.Sprintf(ctx.Messages().HeroChooseWhereSpendLevelUpBonus, hero.LevelUpBonusesLeft),
	)
	msg.ReplyMarkup = keyboard

	return h.Helper.Send(msg)
}

func (h *Handler) HandleSpendLevelUpBonus(ctx *types.Context, query *tgbotapi.CallbackQuery, payload []string) error {
	if err := h.ValidatePayloadLength(payload, 2); err != nil {
		return err
	}

	heroID, err := h.GetIDFromString(payload[0])
	if err != nil {
		return err
	}

	hero, err := h.Services.Hero.FindMy(ctx, heroID)
	if err != nil {
		return err
	}

	statToUpgrade := payload[1] // @todo enum type

	if hero, err = h.Services.Hero.SpendLevelUpBonus(ctx, hero.ID, statToUpgrade); err != nil {
		return err
	}

	if hero.LevelUpBonusesLeft < 1 {
		return h.OpenMyHero(ctx, hero, query)
	}

	return h.HandleOpenLevelUpBonusesOverview(ctx, query, payload)
}
