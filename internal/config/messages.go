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

	MenuItemStartTournamentFight string
	MenuItemDeleteHero           string
	HeroDeleteConfirmationPrompt string
	HeroDeleteCancelled          string
	HeroDeleteSuccessful         string

	HeroCreationStart           string
	HeroCreationSelectClass     string
	HeroCreationEnterNamePrompt string

	HeroCreationSuccess string

	HeroOverview              string
	HeroName                  string
	HeroClass                 string
	HeroAttack                string
	HeroDefense               string
	HeroCriticalChancePercent string
	HeroEvasionChancePercent  string
	HP                        string
	HeroEnergy                string
	HeroGold                  string

	HeroNameWithStats string

	FightOverview       string
	FightStatistics     string
	FightChooseAction   string
	FightActionPunch    string
	FightActionRunAway  string
	FightRunAwaySuccess string

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

	classCharacteristics := heroClass.GetCharacteristics()

	msg += fmt.Sprintf(
		"%s: %v\n",
		m.HeroAttack,
		classCharacteristics.Attack,
	)

	msg += fmt.Sprintf(
		"%s: %v\n",
		m.HeroDefense,
		classCharacteristics.Defense,
	)

	msg += fmt.Sprintf(
		"%s: %v\n",
		m.HP,
		classCharacteristics.StartHP,
	)

	msg += fmt.Sprintf(
		"%s: %v%%\n",
		m.HeroCriticalChancePercent,
		classCharacteristics.CriticalChancePercent,
	)

	msg += fmt.Sprintf(
		"%s: %v%%\n\n",
		m.HeroEvasionChancePercent,
		classCharacteristics.EvasionChancePercent,
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
	classCharacteristics := hero.Class.GetCharacteristics()

	msg += fmt.Sprintf(
		"%s: %s\n",
		m.HeroClass,
		heroClass,
	)

	msg += fmt.Sprintf(
		"%s: %v\n",
		m.HeroAttack,
		classCharacteristics.Attack,
	)

	msg += fmt.Sprintf(
		"%s: %v\n",
		m.HeroDefense,
		classCharacteristics.Defense,
	)

	msg += fmt.Sprintf(
		"%s: %v\n",
		m.HP,
		hero.CurrentHP,
	)

	msg += fmt.Sprintf(
		"%s: %v%%\n",
		m.HeroCriticalChancePercent,
		classCharacteristics.CriticalChancePercent,
	)

	msg += fmt.Sprintf(
		"%s: %v%%\n",
		m.HeroEvasionChancePercent,
		classCharacteristics.EvasionChancePercent,
	)

	msg += fmt.Sprintf(
		"\n%s: %v\n\n",
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

func (m *Messages) FightInfo(fight *models.Fight) (string, error) {
	if fight.Hero == nil || fight.Opponent == nil {
		return "", fmt.Errorf("hero or enemy is null in FightInfo(), fight id: %v", fight.ID)
	}

	msg := fmt.Sprintf(
		m.FightOverview,
		fight.Hero.Name,
		fight.Opponent.GetName(),
	)

	msg += fmt.Sprintf(
		"%s:\n\n%s 👤:\n%s\n\n%s:\n%s",
		m.FightStatistics,
		fight.Hero.Name,
		m.GetFightStatistics(fight, fight.Hero),
		fight.Opponent.GetName(),
		m.GetFightStatistics(fight, fight.Opponent),
	)

	msg += fmt.Sprintf("\n%s", m.FightChooseAction)

	return msg, nil
}

func (m *Messages) GetFightStatistics(fight *models.Fight, fighter models.Fighter) string {
	msg := ""

	msg += fmt.Sprintf(
		"%s: %v/%v\n",
		m.HP,
		fight.GetCurrentHPFor(fighter),
		fighter.GetHP(),
	)

	msg += fmt.Sprintf(
		"%s: %v-%v\n",
		m.HeroAttack, // @todo rename
		fighter.GetMinAttack(),
		fighter.GetMaxAttack(),
	)

	msg += fmt.Sprintf(
		"%s: %v\n",
		m.HeroDefense, // @todo rename
		fighter.GetDefense(),
	)

	msg += fmt.Sprintf(
		"%s: %v%%\n",
		m.HeroCriticalChancePercent, // @todo rename
		fighter.GetCriticalChancePercent(),
	)

	msg += fmt.Sprintf(
		"%s: %v%%\n",
		m.HeroEvasionChancePercent, // @todo rename
		fighter.GetEvasionChancePercent(),
	)

	return msg
}
