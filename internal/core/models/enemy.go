package models

import "github.com/arthurshafikov/tg-gladiator/internal/core/constants/enums"

type Enemy struct {
	ID                    int64  `gorm:"->" json:"id"`
	Name                  string `json:"name"`
	HP                    int    `json:"hp"`
	MinAttack             int    `json:"min_attack"`
	MaxAttack             int    `json:"max_attack"`
	Defense               int    `json:"defense"`
	CriticalChancePercent int    `json:"critical_chance_percent"`
	EvasionChancePercent  int    `json:"evasion_chance_percent"`
	GoldRewardMin         int    `json:"gold_reward_min"`
	GoldRewardMax         int    `json:"gold_reward_max"`
}

func (Enemy) TableName() string {
	return "enemies"
}

func (h *Enemy) GetType() enums.FighterType {
	return enums.FighterTypeOpponent
}

func (h *Enemy) GetName() string {
	return h.Name
}

func (h *Enemy) GetHP() int {
	return h.HP
}

func (h *Enemy) GetMinAttack() int {
	return h.MinAttack
}

func (h *Enemy) GetMaxAttack() int {
	return h.MaxAttack
}

func (h *Enemy) GetDefense() int {
	return h.Defense
}
func (h *Enemy) GetCriticalChancePercent() int {
	return h.CriticalChancePercent
}

func (h *Enemy) GetEvasionChancePercent() int {
	return h.EvasionChancePercent
}
