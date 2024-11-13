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

func (hc HeroClass) GetCharacteristics() (*classCharacteristics, error) {
	switch hc {
	case Swordsman:
		return &classCharacteristics{
			Attack:                10,
			Defense:               5,
			CriticalChancePercent: 10,
			EvasionChancePercent:  5,
			StartHP:               100,
		}, nil
	default:
		return nil, fmt.Errorf("undefined class for characteristics: '%s'", hc.ToString())
	}
}
