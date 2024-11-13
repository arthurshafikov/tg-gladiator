package models

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
