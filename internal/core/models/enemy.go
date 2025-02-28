package models

import "github.com/arthurshafikov/tg-gladiator/internal/core/constants/enums"

type Enemy struct {
	ID                    int64  `gorm:"->" json:"id"`
	Level                 int    `json:"level"`
	Name                  string `json:"name"`
	HP                    int    `json:"hp"`
	MinAttack             int    `json:"min_attack"`
	MaxAttack             int    `json:"max_attack"`
	Defense               int    `json:"defense"`
	CriticalChancePercent int    `json:"critical_chance_percent"`
	EvasionChancePercent  int    `json:"evasion_chance_percent"`
	GoldRewardMin         int    `json:"gold_reward_min"`
	GoldRewardMax         int    `json:"gold_reward_max"`
	XpRewardMin           int    `json:"xp_reward_min"`
	XpRewardMax           int    `json:"xp_reward_max"`
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

func (e *Enemy) GetName() string {
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

func (e *Enemy) GetMinGoldReward() int {
	return e.GoldRewardMin
}

func (e *Enemy) GetMaxGoldReward() int {
	return e.GoldRewardMax
}

// @todo calculate reward depending on the enemy level (with koefficient)
// xpReward := int(math.Round(math.Pow(float64(enemy.Level), 1.5) * 10))
// goldReward := int(math.Round(float64(enemy.Level) * (5 + rand.Float64()*5)))
// dropChance := math.Min(0.1+0.02*float64(enemy.Level), 0.5) // до 50% шанс
func (e *Enemy) GetMinXPReward() int {
	return e.XpRewardMin
}

func (e *Enemy) GetMaxXPReward() int {
	return e.XpRewardMax
}
