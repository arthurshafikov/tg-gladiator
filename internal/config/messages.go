package config

import (
	"fmt"

	"github.com/sirupsen/logrus"
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

	HeroName string

	heroClasses map[string]string
	Errors      map[string]string
}

func (m *Messages) GetHeroClass(heroClass string) string {
	if _, ok := m.heroClasses[heroClass]; !ok {
		logrus.Error(fmt.Errorf("undefined hero class: '%s'", heroClass))

		return "undefined"
	}

	return m.heroClasses[heroClass]
}
