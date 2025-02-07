package enums

type FighterType string
type FightActionType string

const (
	FightStatusActive = "active"
	FightStatusFleed  = "fleed"
	FightStatusEnded  = "ended"

	OpponentTypeMob = "mob"

	FighterTypeHero     FighterType = "hero"
	FighterTypeOpponent FighterType = "opponent"

	FightActionSimpleStrike  FightActionType = "FightActionSimpleStrike"
	FightActionStrongStrike  FightActionType = "FightActionStrongStrike"
	FightActionPreciseStrike FightActionType = "FightActionPreciseStrike"
)
