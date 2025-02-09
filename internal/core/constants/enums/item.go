package enums

type ItemCategory string

const (
	ItemEquipsOnHand     = "hand"      // a sword
	ItemEquipsOnTwoHands = "two_hands" // battle axe
	ItemEquipsOnHands    = "hands"     // gloves
	ItemEquipsOnBody     = "body"
	ItemEquipsOnHead     = "head"
	ItemEquipsOnNeck     = "neck"
	ItemEquipsOnFeet     = "feet"
	ItemEquipsOnFinger   = "finger" // maximum 2 per hero
	ItemEquipsOnWaist    = "waist"

	ItemCategoryWeapon    ItemCategory = "weapon"
	ItemCategoryArmor     ItemCategory = "armor"
	ItemCategoryAccessory ItemCategory = "accessory"
)

func ItemCategoryAll() []ItemCategory {
	return []ItemCategory{
		ItemCategoryWeapon,
		ItemCategoryArmor,
		ItemCategoryAccessory,
	}
}
