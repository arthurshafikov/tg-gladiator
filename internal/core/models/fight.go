package models

import (
	"time"
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
	if fighter.GetID() == f.HeroID {
		return f.HeroHP
	}

	return f.OpponentHP
}

type Fighter interface {
	GetID() int64
	GetName() string
	GetHP() int
	GetMinAttack() int
	GetMaxAttack() int
	GetDefense() int
	GetCriticalChancePercent() int
	GetEvasionChancePercent() int
}
