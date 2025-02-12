package enums

import (
	"fmt"

	"github.com/arthurshafikov/tg-gladiator/internal/core/errors"
)

type HeroClass string

const Swordsman HeroClass = "swordsman"

func GetHeroClasses() []HeroClass {
	return []HeroClass{
		Swordsman,
	}
}

func GetHeroClass(class string) (HeroClass, error) {
	heroClass := HeroClass(class)

	for _, heroClassOption := range GetHeroClasses() {
		if heroClassOption == heroClass {
			return heroClass, nil
		}
	}

	return heroClass, errors.ErrInvalidHeroClass
}

func (hc HeroClass) ToString() string {
	return string(hc)
}

type classCharacteristics struct {
	Attack                int
	Defense               int
	CriticalChancePercent int
	EvasionChancePercent  int
	StartHP               int
}

func (hc HeroClass) GetCharacteristics() *classCharacteristics {
	switch hc {
	case Swordsman:
		return &classCharacteristics{
			Attack:                6,
			Defense:               2,
			CriticalChancePercent: 5,
			EvasionChancePercent:  2,
			StartHP:               70,
		}
	default:
		panic(fmt.Errorf("undefined class for characteristics: '%s'", hc.ToString()))
	}
}
