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

	DeleteConfirmation string

	OpenedMenu     string
	DefaultBackBtn string

	MenuItemMyHeroes string
	MyHeroesList     string

	MenuItemDeleteHero           string
	HeroDeleteConfirmationPrompt string
	HeroDeleteCancelled          string
	HeroDeleteSuccessful         string

	HeroCreationStart           string
	HeroCreationSelectClass     string
	HeroCreationEnterNamePrompt string

	HeroCreationSuccess string

	HeroOverview string
	HeroName     string
	HeroClass    string
	HeroAttack   string
	HeroDefense  string
	HeroHP       string
	HeroEnergy   string
	HeroGold     string

	HeroNameWithStats string

	HeroClasses map[string]string
	Errors      map[string]string
}

func (m *Messages) GetHeroNameWithStats(hero *models.Hero) string {
	return fmt.Sprintf(
		m.HeroNameWithStats,
		hero.Name,
		hero.CurrentHP,
		hero.CurrentGold,
	)
}

func (m *Messages) GetHeroClass(heroClass enums.HeroClass) (string, error) {
	if _, ok := m.HeroClasses[heroClass.ToString()]; !ok {
		return "", fmt.Errorf("undefined hero class: '%s'", heroClass.ToString())
	}

	return m.HeroClasses[heroClass.ToString()], nil
}

func (m *Messages) ClassInfo(heroClass enums.HeroClass) (string, error) {
	if _, ok := m.HeroClasses[heroClass.ToString()]; !ok {
		return "", fmt.Errorf("undefined hero class: '%s'", heroClass.ToString())
	}

	msg := fmt.Sprintf(
		"%s: %s\n",
		m.HeroClass,
		m.HeroClasses[heroClass.ToString()],
	)

	attack, err := heroClass.GetAttack()
	if err != nil {
		return "", err
	}
	msg += fmt.Sprintf(
		"%s: %v\n",
		m.HeroAttack,
		attack,
	)

	defense, err := heroClass.GetDefense()
	if err != nil {
		return "", err
	}
	msg += fmt.Sprintf(
		"%s: %v\n",
		m.HeroDefense,
		defense,
	)

	startHP, err := heroClass.GetStartHP()
	if err != nil {
		return "", err
	}
	msg += fmt.Sprintf(
		"%s: %v\n\n",
		m.HeroHP,
		startHP,
	)

	return msg, nil
}

func (m *Messages) GetHeroInfo(hero *models.Hero) (string, error) {
	msg := fmt.Sprintf(
		"%s\n\n%s: %s\n",
		m.HeroOverview,
		m.HeroName,
		hero.Name,
	)

	heroClass, err := m.GetHeroClass(hero.Class)
	if err != nil {
		return "", err
	}
	msg += fmt.Sprintf(
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
