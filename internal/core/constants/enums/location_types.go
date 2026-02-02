package enums

import (
	"slices"

	"github.com/sirupsen/logrus"
)

type FightLocation string

const (
	LocationTypeForest    FightLocation = "forest"
	LocationTypeGraveyard FightLocation = "graveyard"
	LocationTypeCastle    FightLocation = "castle"
)

func AllFightLocations() []FightLocation {
	return []FightLocation{
		LocationTypeForest,
		LocationTypeGraveyard,
		LocationTypeCastle,
	}
}

func FightLocationLevelRange(loc FightLocation) (int, int) {
	values, ok := map[FightLocation][]int{
		LocationTypeForest: {
			1,
			2,
		},
		LocationTypeGraveyard: {
			3,
			4,
		},
		LocationTypeCastle: {
			5,
			6,
		},
	}[loc]

	if !ok {
		logrus.Errorf("undefined location for levels: %s", string(loc))

		return 1, 2
	}

	return values[0], values[1]
}

func GetFightLocation(str string) FightLocation {
	loc := FightLocation(str)

	if !slices.Contains(AllFightLocations(), loc) {
		logrus.Errorf("invalid location passed: %s", str)

		return LocationTypeForest
	}

	return loc
}
