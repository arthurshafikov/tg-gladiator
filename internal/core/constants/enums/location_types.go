package enums

import (
	"slices"

	"github.com/sirupsen/logrus"
)

type FightLocation string

const (
	LocationTypeForest          FightLocation = "forest"
	LocationTypeGraveyard       FightLocation = "graveyard"
	LocationTypeCastle          FightLocation = "castle"
	LocationTypeDragonMountains FightLocation = "dragon_mountains"
	LocationTypeDarkForest      FightLocation = "dark_forest"
	LocationTypeIceCaves        FightLocation = "ice_caves"
	LocationTypeVolcanicPlains  FightLocation = "volcanic_plains"
	LocationTypeDesertDunes     FightLocation = "desert_dunes"
	LocationTypeOceanDepths     FightLocation = "ocean_depths"
	LocationTypeStormPeaks      FightLocation = "storm_peaks"
)

func AllFightLocations() []FightLocation {
	return []FightLocation{
		LocationTypeForest,
		LocationTypeGraveyard,
		LocationTypeCastle,
		LocationTypeDragonMountains,
		LocationTypeDarkForest,
		LocationTypeIceCaves,
		LocationTypeVolcanicPlains,
		LocationTypeDesertDunes,
		LocationTypeOceanDepths,
		LocationTypeStormPeaks,
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
		LocationTypeDragonMountains: {
			7,
			8,
		},
		LocationTypeDarkForest: {
			9,
			10,
		},
		LocationTypeIceCaves: {
			11,
			12,
		},
		LocationTypeVolcanicPlains: {
			13,
			14,
		},
		LocationTypeDesertDunes: {
			15,
			16,
		},
		LocationTypeOceanDepths: {
			17,
			18,
		},
		LocationTypeStormPeaks: {
			19,
			20,
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
