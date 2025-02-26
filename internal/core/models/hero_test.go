package models

import (
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
