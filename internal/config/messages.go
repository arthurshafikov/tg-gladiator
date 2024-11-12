package config

import (
	"fmt"

	"github.com/arthurshafikov/tg-gladiator/internal/core/constants/enums"
	"github.com/arthurshafikov/tg-gladiator/internal/core/models"
)

type MessagesBag struct {
	RU Messages
	EN Messages
}

type Messages struct {
	StartSuccess string
	Help         string

	OpenedMenu     string
	DefaultBackBtn string

	MenuItemMyHeroes string
	MyHeroesList     string

	HeroCreationStart       string
	HeroCreationSelectClass string

	HeroCreationSuccess string

	HeroClass   string
	HeroAttack  string
	HeroDefense string
	HeroHP      string
	HeroEnergy  string
	HeroGold    string

	HeroName string

	HeroClasses map[string]string
	Errors      map[string]string
}

func (m *Messages) GetHeroName(hero *models.Hero) (string, error) {
	heroClass, err := m.GetHeroClass(hero.Class)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf(
		m.HeroName,
		heroClass,
		hero.CurrentHP,
		hero.CurrentGold,
	), nil
}

func (m *Messages) GetHeroClass(heroClass enums.HeroClass) (string, error) {
	fmt.Printf("%#v\n", m.HeroClasses)
	if _, ok := m.HeroClasses[heroClass.ToString()]; !ok {
		return "", fmt.Errorf("undefined hero class: '%s'", heroClass.ToString())
	}

	return m.HeroClasses[heroClass.ToString()], nil
}

func (m *Messages) GetHeroInfo(hero *models.Hero) (string, error) {
	heroClass, err := m.GetHeroClass(hero.Class)
	if err != nil {
		return "", err
	}

	msg := fmt.Sprintf(
		"%s: %s\n",
		m.HeroClass,
		heroClass,
	)

	attack, err := hero.Class.GetAttack()
	if err != nil {
		return "", err
	}

	msg += fmt.Sprintf(
		"%s: %v\n",
		m.HeroAttack,
		attack,
	)

	defense, err := hero.Class.GetDefense()
	if err != nil {
		return "", err
	}

	msg += fmt.Sprintf(
		"%s: %v\n",
		m.HeroDefense,
		defense,
	)

	msg += fmt.Sprintf(
		"%s: %v\n\n",
		m.HeroHP,
		hero.CurrentHP,
	)

	msg += fmt.Sprintf(
		"%s: %v\n\n",
		m.HeroGold,
		hero.CurrentGold,
	)

	msg += fmt.Sprintf(
		"%s: %v\n",
		m.HeroEnergy,
		hero.CurrentEnergy,
	)

	return msg, nil
}
