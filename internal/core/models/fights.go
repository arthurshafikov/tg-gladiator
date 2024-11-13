package models

import "time"

type Fight struct {
	ID           int64     `json:"id"`
	Status       string    `json:"status"`
	HeroID       int64     `json:"hero_id"`
	OpponentType int64     `json:"opponent_type"`
	OpponentID   int64     `json:"opponent_id"`
	HeroHP       int       `json:"hero_hp"`
	OpponentHP   int       `json:"opponent_hp"`
	CreatedAt    time.Time `json:"created_at"`
}
