package models

import (
	"fmt"
	"testing"
	"time"
)

func TestHeroShop_GetUpdatesInText(t *testing.T) {
	tests := []struct {
		updatesAt time.Time
		want      string
	}{
		{
			updatesAt: time.Now().Add(time.Minute * -5),
			want:      "1 минут",
		},
		{
			updatesAt: time.Now(),
			want:      "1 минут",
		},
		{
			updatesAt: time.Now().Add(time.Minute * 27),
			want:      "26 минут",
		},
		{
			updatesAt: time.Now().Add(time.Hour * 15),
			want:      "14 час(а) 59 минут",
		},
		{
			updatesAt: time.Now().Add(time.Hour*2 + time.Minute*15),
			want:      "2 час(а) 14 минут",
		},
		{
			updatesAt: time.Now().Add(time.Hour*50 + time.Minute*31),
			want:      "50 час(а) 30 минут",
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("case expects %s", tt.want), func(t *testing.T) {
			hs := &HeroShop{
				UpdatesAt: tt.updatesAt,
			}
			if got := hs.GetUpdatesInText(); got != tt.want {
				t.Errorf("HeroShop.GetUpdatesInText() = %v, want %v", got, tt.want)
			}
		})
	}
}
