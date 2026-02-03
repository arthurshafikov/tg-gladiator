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
		// XP на 1 уровень: int(25*1^2.2) = 25
		{
			currentXP:      0,
			currentLevel:   1,
			expectedXPLeft: 25,
		},
		// XP на 1+2 уровень: int(25*1^2.2)+int(25*2^2.2) = 25+114=139
		{
			currentXP:      51,
			currentLevel:   2,
			expectedXPLeft: 88, // 139-51
		},
		// XP на 1+2+3 уровень: 25+114+282=421
		{
			currentXP:      250,
			currentLevel:   3,
			expectedXPLeft: 171, // 421-250
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
		// 1 уровень: 0-24 XP
		{
			xp:            0,
			expectedLevel: 1,
		},
		{
			xp:            24,
			expectedLevel: 1,
		},
		// 2 уровень: 25-138 XP
		{
			xp:            25,
			expectedLevel: 2,
		},
		{
			xp:            138,
			expectedLevel: 2,
		},
		// 3 уровень: 139-420 XP
		{
			xp:            139,
			expectedLevel: 3,
		},
		{
			xp:            420,
			expectedLevel: 3,
		},
		// 4 уровень: 421+
		{
			xp:            421,
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
	t.Skip("stats")

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
