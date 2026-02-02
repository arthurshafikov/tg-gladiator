package models

import (
	"fmt"
	"testing"
)

func TestGetXPForNextLevelLeft(t *testing.T) {
	tests := []struct {
		currentXP      int
		currentLevel   int
		expectedXPLeft int
	}{
		{
			currentXP:      0,
			currentLevel:   1,
			expectedXPLeft: 20,
		},
		{
			currentXP:      51,
			currentLevel:   2,
			expectedXPLeft: 49,
		},
		{
			currentXP:      250,
			currentLevel:   3,
			expectedXPLeft: 30,
		},
	}
	for _, tt := range tests {
		t.Run("case", func(t *testing.T) {
			h := &Hero{
				XP:    tt.currentXP,
				Level: tt.currentLevel,
			}
			if got := h.GetXPForNextLevelLeft(); got != tt.expectedXPLeft {
				t.Errorf("Hero.GetXPForNextLevelLeft() = %v, want %v", got, tt.expectedXPLeft)
			}
		})
	}
}

func TestCalculateLevelForXP(t *testing.T) {
	tests := []struct {
		xp            int
		expectedLevel int
	}{
		{
			xp:            19,
			expectedLevel: 1,
		},
		{
			xp:            0,
			expectedLevel: 1,
		},
		{
			xp:            20,
			expectedLevel: 2,
		},
		{
			xp:            99,
			expectedLevel: 2,
		},
		{
			xp:            100,
			expectedLevel: 3,
		},
		{
			xp:            279,
			expectedLevel: 3,
		},
		{
			xp:            280,
			expectedLevel: 4,
		},
	}
	for _, tt := range tests {
		t.Run("case", func(t *testing.T) {
			h := &Hero{}
			if got := h.CalculateLevelForXP(tt.xp); got != tt.expectedLevel {
				t.Errorf("Hero.CalculateLevelForXP() = %v, want %v", got, tt.expectedLevel)
			}
		})
	}
}

func TestTotalGetXPLevelInfo(t *testing.T) {

	h := Hero{
		Level: 1,
		XP:    0,
	}

	for ; h.Level < 25; h.Level++ {
		xpLeft := h.GetXPForNextLevelLeft()

		h.XP += xpLeft

		opponent := Enemy{
			Level: h.Level,
		}

		enemyKills := xpLeft / opponent.GetXPReward()

		fmt.Printf("For level: %v, required XP - %v, total XP - %v, need to kill %v enemies", h.Level+1, xpLeft, h.XP, enemyKills)

		fmt.Println()
	}
}
