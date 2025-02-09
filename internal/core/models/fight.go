package models

import (
	"math"
	"time"

	"github.com/arthurshafikov/tg-gladiator/internal/core/constants/enums"
)

const (
	FightFieldHeroID     = "hero_id"
	FightFieldStatus     = "status"
	FightFieldOpponentHP = "opponent_hp"
	FightFieldHeroHP     = "hero_hp"
	FightFieldGoldReward = "gold_reward"

	FightArmorDamageReductionModificator = 40 // the more - the less % of damage will get blocked
)

type Fight struct {
	ID           int64  `json:"id"`
	Status       string `json:"status"`
	HeroID       int64  `json:"hero_id"`
	HeroHP       int    `json:"hero_hp"`
	OpponentType string `json:"opponent_type"`
	OpponentID   int64  `json:"opponent_id"`
	OpponentHP   int    `json:"opponent_hp"`
	GoldReward   int    `json:"gold_reward"`

	CreatedAt time.Time `json:"created_at"`

	Hero     *Hero
	Opponent Fighter `gorm:"-"`
}

func (f *Fight) GetCurrentHPFor(fighter Fighter) int {
	if fighter.GetType() == enums.FighterTypeHero {
		return f.HeroHP
	}

	return f.OpponentHP
}

func (f *Fight) IsHeroWon() bool {
	return f.OpponentHP <= 0
}

func (f *Fight) HasEnded() bool {
	return f.Status == enums.FightStatusEnded
}

type Fighter interface {
	GetID() int64
	GetType() enums.FighterType
	GetName() string
	GetHP() int
	HasAttackRange() bool
	GetMinAttack() int
	GetMaxAttack() int
	GetDefense() int
	GetCriticalChancePercent() int
	GetEvasionChancePercent() int
	GetMinReward() int
	GetMaxReward() int
}

// @todo DTO?
type FightEvent struct {
	ActionType         enums.FightActionType
	DamageDealt        int
	DamageReceived     int
	DamageBlocked      int
	IsCritical         bool
	EvadeChancePercent int
	WasEvaded          bool
}

type FightEvents struct {
	HeroFightEvent     FightEvent
	OpponentFightEvent FightEvent
}

func CalculateArmorReductionMultiplier(defense int) float64 {
	if defense < 1 {
		return 0
	}

	return float64(defense) / float64(defense+FightArmorDamageReductionModificator)
}

func CalculateArmorReductionPercent(defense int) float64 {
	multiplier := CalculateArmorReductionMultiplier(defense)

	return math.Round(multiplier * 100)
}
