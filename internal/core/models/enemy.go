package models

import (
	"math/rand"

	"github.com/arthurshafikov/tg-gladiator/internal/core/constants/enums"
	"gorm.io/gorm"
)

const (
	EnemyFieldLevel    = "level"
	EnemyFieldBossType = "boss_type"
)

type Enemy struct {
	ID                    int64          `gorm:"->" json:"id"`
	Level                 int            `json:"level"`
	Name                  string         `json:"name"`
	HP                    int            `json:"hp"`
	MinAttack             int            `json:"min_attack"`
	MaxAttack             int            `json:"max_attack"`
	Defense               int            `json:"defense"`
	CriticalChancePercent int            `json:"critical_chance_percent"`
	EvasionChancePercent  int            `json:"evasion_chance_percent"`
	BossType              string         `json:"boss_type"`
	DeletedAt             gorm.DeletedAt `gorm:"index"`
}

func (Enemy) TableName() string {
	return "enemies"
}

func (e *Enemy) GetID() int64 {
	return e.ID
}

func (e *Enemy) GetType() enums.FighterType {
	return enums.FighterTypeOpponent
}

func (e *Enemy) GetName(enemyNames ...map[string]string) string {
	if len(enemyNames) > 0 {
		if name, ok := enemyNames[0][e.Name]; ok {
			return name
		}
	}

	return e.Name
}

func (e *Enemy) GetLevel() int {
	return e.Level
}

func (e *Enemy) GetHP() int {
	return e.HP
}

func (e *Enemy) HasAttackRange() bool {
	return e.GetMaxAttack() != e.GetMinAttack()
}

func (e *Enemy) GetMinAttack() int {
	return e.MinAttack
}

func (e *Enemy) GetMaxAttack() int {
	return e.MaxAttack
}

func (e *Enemy) GetDefense() int {
	return e.Defense
}
func (e *Enemy) GetCriticalChancePercent() int {
	return e.CriticalChancePercent
}

func (e *Enemy) GetEvasionChancePercent() int {
	return e.EvasionChancePercent
}

func (h *Enemy) GetXPReward() int {
	return h.GetLevel() * 10
}

func (h *Enemy) GetGoldReward() int {
	minGold := h.GetLevel() * 5
	maxGold := h.GetLevel() * 15
	return rand.Intn(maxGold-minGold+1) + minGold
}
