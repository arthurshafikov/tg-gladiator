package enums

import (
	"slices"

	"github.com/sirupsen/logrus"
)

type FightLocation string

const (
	LocationTypeForest          FightLocation = "forest"
	LocationTypeGraveyard       FightLocation = "graveyard"
	LocationTypeAbandonedCastle FightLocation = "abandoned_castle"
	LocationTypeFireCaves       FightLocation = "fire_caves"
	LocationTypeNetherWorld     FightLocation = "nether_world"
	LocationTypeIceCaves        FightLocation = "ice_caves"
	LocationTypeDesert          FightLocation = "desert"
	LocationTypeEastFortress    FightLocation = "east_fortress"
	LocationTypeWebCatacombs    FightLocation = "web_catacombs"
	LocationTypeSteamCity       FightLocation = "steam_city"
	LocationTypeMagicAcademy    FightLocation = "magic_academy"
	LocationTypeEndLands        FightLocation = "end_lands"
)

func AllFightLocations() []FightLocation {
	return []FightLocation{
		LocationTypeForest,
		LocationTypeGraveyard,
		LocationTypeAbandonedCastle,
		LocationTypeFireCaves,
		LocationTypeNetherWorld,
		LocationTypeIceCaves,
		LocationTypeDesert,
		LocationTypeEastFortress,
		LocationTypeWebCatacombs,
		LocationTypeSteamCity,
		LocationTypeMagicAcademy,
		LocationTypeEndLands,
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
		LocationTypeAbandonedCastle: {
			5,
			6,
		},
		LocationTypeFireCaves: {
			7,
			8,
		},
		LocationTypeNetherWorld: {
			9,
			10,
		},
		LocationTypeIceCaves: {
			11,
			12,
		},
		LocationTypeDesert: {
			13,
			14,
		},
		LocationTypeEastFortress: {
			15,
			16,
		},
		LocationTypeWebCatacombs: {
			17,
			18,
		},
		LocationTypeSteamCity: {
			19,
			20,
		},
		LocationTypeMagicAcademy: {
			21,
			22,
		},
		LocationTypeEndLands: {
			23,
			25,
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
