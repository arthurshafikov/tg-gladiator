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

	FightActionPunch        FightActionType = "FightActionPunch"
	FightActionStrongPunch  FightActionType = "FightActionStrongPunch"
	FightActionPrecisePunch FightActionType = "FightActionPrecisePunch"
)
