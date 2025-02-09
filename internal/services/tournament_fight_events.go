package services

import (
	"math"

	"github.com/arthurshafikov/tg-gladiator/internal/core/constants/enums"
	"github.com/arthurshafikov/tg-gladiator/internal/core/errors"
	"github.com/arthurshafikov/tg-gladiator/internal/core/helpers"
	"github.com/arthurshafikov/tg-gladiator/internal/core/models"
)

type TournamentFightEventsService struct {
	logger Logger
}

func newTournamentFightEventsService(logger Logger) *TournamentFightEventsService {
	return &TournamentFightEventsService{
		logger: logger,
	}
}

func (s *TournamentFightEventsService) getFightEventFor(
	actionType enums.FightActionType,
	fighter models.Fighter,
	opponent models.Fighter,
) (models.FightEvent, error) {
	fighterEvent := models.FightEvent{
		ActionType: actionType,
	}

	var err error
	fighterEvent.DamageDealt, fighterEvent.IsCritical, err = s.calculateDamageDealt(actionType, fighter)
	if err != nil {
		return fighterEvent, err
	}

	fighterEvent.WasEvaded, fighterEvent.EvadeChancePercent, err = s.calculateIfTheStrikeWasEvaded(actionType, opponent)
	if err != nil {
		return fighterEvent, err
	}

	if !fighterEvent.WasEvaded {
		fighterEvent.DamageReceived, fighterEvent.DamageBlocked = s.calculateReceivedDamage(
			opponent.GetDefense(),
			fighterEvent.DamageDealt,
		)
	}

	return fighterEvent, nil
}

func (s *TournamentFightEventsService) getRandomFightEventFor(
	fighter models.Fighter,
	opponent models.Fighter,
) (models.FightEvent, error) {
	allFightActionTypes := enums.FightActionTypeAll()

	randomIndex, err := s.generateRandomNumberInRange(0, (len(allFightActionTypes) - 1))
	if err != nil {
		return models.FightEvent{}, err
	}

	return s.getFightEventFor(
		allFightActionTypes[randomIndex],
		fighter,
		opponent,
	)
}

func (s *TournamentFightEventsService) calculateDamageDealt(
	actionType enums.FightActionType,
	fighter models.Fighter,
) (int, bool, error) {
	var baseDamage int
	var err error
	if fighter.HasAttackRange() {
		baseDamage, err = s.generateRandomNumberInRange(fighter.GetMinAttack(), fighter.GetMaxAttack())
		if err != nil {
			s.logger.Error(err)

			return 0, false, errors.ErrServerError
		}
	} else {
		baseDamage = fighter.GetMinAttack()
	}

	damageDealt := s.calculateStrikeTypeModificatorDamage(actionType, baseDamage)

	isCritical, err := s.calculateRandomChance(fighter.GetCriticalChancePercent())
	if err != nil {
		return 0, false, err
	}
	if isCritical {
		damageDealt *= 2
	}

	return damageDealt, isCritical, nil
}

func (s *TournamentFightEventsService) calculateIfTheStrikeWasEvaded(
	actionType enums.FightActionType,
	opponent models.Fighter,
) (bool, int, error) {
	var additionalEvasionPercentBasedOnActionType int
	var err error
	switch actionType {
	case enums.FightActionSimpleStrike:
		additionalEvasionPercentBasedOnActionType = 0
	case enums.FightActionStrongStrike:
		additionalEvasionPercentBasedOnActionType, err = s.generateRandomNumberInRange(15, 25)
		if err != nil {
			return false, 0, err
		}
	case enums.FightActionPreciseStrike:
		additionalEvasionPercentBasedOnActionType = -25 // @todo should be propotional instead of raw deduction
	}

	evasionChance := opponent.GetEvasionChancePercent() + additionalEvasionPercentBasedOnActionType
	wasEvaded, err := s.calculateRandomChance(evasionChance)
	if err != nil {
		return false, 0, err
	}

	return wasEvaded, evasionChance, nil
}

func (s *TournamentFightEventsService) calculateStrikeTypeModificatorDamage(
	actionType enums.FightActionType,
	baseDamage int,
) int {
	if baseDamage < 1 {
		return 0
	}

	var damageDealt int
	switch actionType {
	case enums.FightActionSimpleStrike:
		damageDealt = baseDamage
	case enums.FightActionStrongStrike:
		damageDealt = int(math.Round(float64(baseDamage) * 1.5)) // @todo range for variety
	case enums.FightActionPreciseStrike:
		damageDealt = int(math.Round(float64(baseDamage) * 0.7))
	}

	return damageDealt
}

func (s *TournamentFightEventsService) calculateReceivedDamage(defense, damageDealt int) (int, int) {
	damageReductionMultiplier := models.CalculateArmorReductionMultiplier(defense)

	blockedDamage := int(math.Round(float64(damageDealt) * damageReductionMultiplier))

	return damageDealt - blockedDamage, blockedDamage
}

func (s *TournamentFightEventsService) calculateRandomChance(desiredChance int) (bool, error) {
	if desiredChance < 1 {
		return false, nil
	}

	random, err := s.generateRandomNumberInRange(1, 100)
	if err != nil {
		s.logger.Error(err)

		return false, errors.ErrServerError
	}

	return random <= desiredChance, nil
}

func (s *TournamentFightEventsService) generateRandomNumberInRange(min, max int) (int, error) {
	random, err := helpers.GenerateRandomNumberInRange(min, max)
	if err != nil {
		s.logger.Error(err)

		return 0, errors.ErrServerError
	}

	return random, nil
}
