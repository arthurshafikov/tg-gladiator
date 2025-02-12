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
	MenuItemOpenHeroEquipment    string
	MenuItemShop                 string
	MenuItemDeleteHero           string
	HeroDeleteConfirmationPrompt string
	HeroDeleteCancelled          string
	HeroDeleteSuccessful         string

	HeroCreationStart           string
	HeroCreationSelectClass     string
	HeroCreationEnterNamePrompt string

	HeroCreationSuccess string

	HeroOverview          string
	HeroName              string
	HeroClass             string
	Attack                string
	Defense               string
	CriticalChancePercent string
	EvasionChancePercent  string
	HP                    string
	HeroEnergy            string
	HeroGold              string

	HeroNameWithStats string

	FightOverview            string
	FightChooseAction        string
	FightActionSimpleStrike  string
	FightActionStrongStrike  string
	FightActionPreciseStrike string
	FightActionRunAway       string
	FightRunAwaySuccess      string
	PleaseFinishActiveFight  string

	FightTurnOverview                    string
	FightActionInfoDealtSimpleStrike     string
	FightActionInfoDealtStrongStrike     string
	FightActionInfoDealtPreciseStrike    string
	FightActionInfoDealtSimpleAttack     string
	FightActionInfoCritical              string
	FightActionInfoOpponentHPLost        string
	FightActionInfoOpponentDamageBlocked string
	FightActionInfoOpponentEvaded        string

	FightResultHeroWon     string
	FightResultOpponentWon string
	FightResultBackButton  string

	ShopIntro               string
	ShopYourBalance         string
	ShopBuyItemConfirmation string
	ShopBuyItemConfirm      string
	ShopBackButton          string
	ShopBuyItemSuccess      string

	HeroEquipmentOverview    string
	HeroEquipmentChooseItem  string
	HeroEquipmentEquipItem   string
	HeroEquipmentUnequipItem string

	ItemName         string
	ItemCategory     string
	ItemEquipsOn     string
	ItemAttackBonus  string
	ItemDefenseBonus string
	ItemCritBonus    string
	ItemEvasionBonus string

	PreviousPage string
	NextPage     string

	HeroClasses    map[string]string
	ItemCategories map[enums.ItemCategory]string
	ItemEquipsOns  map[enums.ItemEquipsOn]string
	ItemNames      map[string]string
	Errors         map[string]string
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
		m.Attack,
		classCharacteristics.Attack,
	)

	msg += fmt.Sprintf(
		"%s: %v\n",
		m.Defense,
		classCharacteristics.Defense,
	)

	msg += fmt.Sprintf(
		"%s: %v\n",
		m.HP,
		classCharacteristics.StartHP,
	)

	msg += fmt.Sprintf(
		"%s: %v%%\n",
		m.CriticalChancePercent,
		classCharacteristics.CriticalChancePercent,
	)

	msg += fmt.Sprintf(
		"%s: %v%%\n\n",
		m.EvasionChancePercent,
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

	msg += fmt.Sprintf(
		"%s: %s\n",
		m.HeroClass,
		heroClass,
	)

	msg += fmt.Sprintf(
		"%s: %v\n",
		m.HP,
		hero.CurrentHP,
	)

	msg += fmt.Sprintf(
		"%s: %v\n",
		m.Attack,
		hero.GetMinAttack(),
	)

	msg += fmt.Sprintf(
		"%s: %v%%\n",
		m.CriticalChancePercent,
		hero.GetCriticalChancePercent(),
	)

	msg += fmt.Sprintf(
		"%s: %v\n",
		m.Defense,
		hero.GetDefense(),
	)

	msg += fmt.Sprintf(
		"%s: %v%%\n",
		m.EvasionChancePercent,
		hero.GetEvasionChancePercent(),
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
		"%s\n\n",
		m.FightOverview,
	)

	msg += fmt.Sprintf(
		"%s 👤:\n%s\n\n%s:\n%s",
		fight.Hero.GetName(),
		m.GetFightStatistics(fight, fight.Hero),
		fight.Opponent.GetName(),
		m.GetFightStatistics(fight, fight.Opponent),
	)

	if !fight.HasEnded() {
		msg += fmt.Sprintf("\n\n%s", m.FightChooseAction)
	}

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

	if fighter.GetMinAttack() != fighter.GetMaxAttack() {
		msg += fmt.Sprintf(
			"%s: %v-%v\n",
			m.Attack,
			fighter.GetMinAttack(),
			fighter.GetMaxAttack(),
		)
	} else {
		msg += fmt.Sprintf(
			"%s: %v\n",
			m.Attack,
			fighter.GetMinAttack(),
		)
	}

	msg += fmt.Sprintf(
		"%s: %v%%\n",
		m.CriticalChancePercent,
		fighter.GetCriticalChancePercent(),
	)

	msg += fmt.Sprintf(
		"%s: %v (%v%%)\n",
		m.Defense,
		fighter.GetDefense(),
		models.CalculateArmorReductionPercent(fighter.GetDefense()),
	)

	msg += fmt.Sprintf(
		"%s: %v%%",
		m.EvasionChancePercent,
		fighter.GetEvasionChancePercent(),
	)

	return msg
}

func (m *Messages) FightEvents(fight *models.Fight, fightEvents *models.FightEvents) string {
	msg := m.FightTurnOverview

	msg += "\n\n"

	msg += fmt.Sprintf("%s 👤\n", fight.Hero.GetName())
	msg += fmt.Sprintf("%s\n\n", m.FightActionInfo(fightEvents.HeroFightEvent, fight.Opponent))

	msg += fmt.Sprintf("%s\n", fight.Opponent.GetName())
	msg += fmt.Sprintf("%s\n\n", m.FightActionInfo(fightEvents.OpponentFightEvent, fight.Hero))

	return msg
}

func (m *Messages) FightActionInfo(fightEvent models.FightEvent, opponent models.Fighter) string {
	var actionInfo string
	switch fightEvent.ActionType {
	case enums.FightActionSimpleStrike:
		actionInfo = m.FightActionInfoDealtSimpleStrike // @todo strike
	case enums.FightActionStrongStrike:
		actionInfo = m.FightActionInfoDealtStrongStrike
	case enums.FightActionPreciseStrike:
		actionInfo = m.FightActionInfoDealtPreciseStrike
	}

	var criticalInfo string
	if fightEvent.IsCritical {
		criticalInfo = fmt.Sprintf(" (%s)", m.FightActionInfoCritical)
	}

	var opponentReactionInfo string
	if fightEvent.WasEvaded {
		opponentReactionInfo = fmt.Sprintf(
			m.FightActionInfoOpponentEvaded,
			opponent.GetName(),
			fightEvent.EvadeChancePercent,
		)
	} else {
		var damageBlockedText string
		if fightEvent.DamageBlocked > 0 {
			damageBlockedText = fmt.Sprintf(
				" %s",
				fmt.Sprintf(
					m.FightActionInfoOpponentDamageBlocked,
					fightEvent.DamageBlocked,
				),
			)
		}

		opponentReactionInfo = fmt.Sprintf(
			"%s%s",
			fmt.Sprintf(
				m.FightActionInfoOpponentHPLost,
				opponent.GetName(),
				fightEvent.DamageReceived,
			),
			damageBlockedText,
		)
	}

	return fmt.Sprintf(
		"%s%s %s",
		fmt.Sprintf(
			actionInfo,
			fightEvent.DamageDealt,
		),
		criticalInfo,
		opponentReactionInfo,
	)
}

func (m *Messages) ItemDescription(item models.Item) string {
	msg := "\n"

	msg += fmt.Sprintf("%s: %s\n", m.ItemName, m.ItemNames[item.Name])
	msg += fmt.Sprintf("%s: %s\n", m.ItemCategory, m.ItemCategories[item.Category])
	msg += fmt.Sprintf("%s: %s\n", m.ItemEquipsOn, m.ItemEquipsOns[item.EquipsOn])

	if item.AttackBonus != 0 {
		if item.AttackBonus < 0 {
			msg += fmt.Sprintf("%s: %v💪\n", m.ItemAttackBonus, item.AttackBonus)
		} else {
			msg += fmt.Sprintf("%s: +%v💪\n", m.ItemAttackBonus, item.AttackBonus)
		}
	}
	if item.DefenseBonus != 0 {
		if item.DefenseBonus < 0 {
			msg += fmt.Sprintf("%s: %v🛡\n", m.ItemDefenseBonus, item.DefenseBonus)
		} else {
			msg += fmt.Sprintf("%s: +%v🛡\n", m.ItemDefenseBonus, item.DefenseBonus)
		}
	}
	if item.CriticalChancePercentBonus != 0 {
		if item.CriticalChancePercentBonus < 0 {
			msg += fmt.Sprintf("%s: %v%%💥\n", m.ItemCritBonus, item.CriticalChancePercentBonus)
		} else {
			msg += fmt.Sprintf("%s: +%v%%💥\n", m.ItemCritBonus, item.CriticalChancePercentBonus)
		}
	}
	if item.EvasionPercentBonus != 0 {
		if item.EvasionPercentBonus < 0 {
			msg += fmt.Sprintf("%s: %v%%🍀\n", m.ItemEvasionBonus, item.EvasionPercentBonus)
		} else {
			msg += fmt.Sprintf("%s: +%v%%🍀\n", m.ItemEvasionBonus, item.EvasionPercentBonus)
		}
	}

	return msg
}

func (m *Messages) GetItemNameWithIcon(item models.Item) string {
	icon := ""
	switch item.Category {
	case enums.ItemCategoryWeapon:
		icon = "🗡"
	case enums.ItemCategoryArmor:
		icon = "🛡"
	case enums.ItemCategoryAccessory:
		icon = "📿"
	}

	return fmt.Sprintf("%s %s", icon, m.ItemNames[item.Name])
}
