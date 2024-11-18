package models

import (
	"time"

	"github.com/arthurshafikov/tg-gladiator/internal/core/constants/enums"
)

type Fight struct {
	ID           int64     `json:"id"`
	Status       string    `json:"status"`
	HeroID       int64     `json:"hero_id"`
	OpponentType string    `json:"opponent_type"`
	OpponentID   int64     `json:"opponent_id"`
	HeroHP       int       `json:"hero_hp"`
	OpponentHP   int       `json:"opponent_hp"`
	CreatedAt    time.Time `json:"created_at"`

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
	GetType() enums.FighterType
	GetName() string
	GetHP() int
	GetMinAttack() int
	GetMaxAttack() int
	GetDefense() int
	GetCriticalChancePercent() int
	GetEvasionChancePercent() int
}

type FightEvent struct {
	ActionType     enums.FightActionType
	DamageDealt    int
	DamageReceived int
	IsCritical     bool
	WasEvaded      bool
}

type FightEvents struct {
	HeroFightEvent     FightEvent
	OpponentFightEvent FightEvent
}
