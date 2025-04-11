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
			expectedXPLeft: 50,
		},
		{
			currentXP:      51,
			currentLevel:   2,
			expectedXPLeft: 199,
		},
		{
			currentXP:      250,
			currentLevel:   3,
			expectedXPLeft: 450,
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
			xp:            49,
			expectedLevel: 1,
		},
		{
			xp:            0,
			expectedLevel: 1,
		},
		{
			xp:            51,
			expectedLevel: 2,
		},
		{
			xp:            249,
			expectedLevel: 2,
		},
		{
			xp:            250,
			expectedLevel: 3,
		},
		{
			xp:            451,
			expectedLevel: 3,
		},
		{
			xp:            699,
			expectedLevel: 3,
		},
		{
			xp:            700,
			expectedLevel: 4,
		},
		{
			xp:            12345678,
			expectedLevel: 90,
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
	t.Skip("uncomment to see stats")

	h := Hero{
		Level: 1,
		XP:    0,
	}

	enemiesAvgXPForLevel := map[int]int{
		1: 15,
		2: 20,
		3: 40,
		4: 50,
		5: 65,
		6: 75,
		7: 100,
		8: 120,
	}

	for ; h.Level < 25; h.Level++ {
		xpLeft := h.GetXPForNextLevelLeft()

		h.XP += xpLeft

		enemyKills := 0
		if avgXP, ok := enemiesAvgXPForLevel[h.Level]; ok {
			enemyKills = xpLeft / avgXP
		}

		fmt.Printf("For level: %v, required XP - %v, total XP - %v, need to kill %v enemies", h.Level+1, xpLeft, h.XP, enemyKills)

		fmt.Println()
	}
}
