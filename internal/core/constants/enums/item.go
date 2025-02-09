package enums

type ItemEquipsOn string
type ItemCategory string

const (
	ItemEquipsOnHand     ItemEquipsOn = "hand"      // a sword
	ItemEquipsOnTwoHands ItemEquipsOn = "two_hands" // battle axe
	ItemEquipsOnHands    ItemEquipsOn = "hands"     // gloves
	ItemEquipsOnBody     ItemEquipsOn = "body"
	ItemEquipsOnHead     ItemEquipsOn = "head"
	ItemEquipsOnNeck     ItemEquipsOn = "neck"
	ItemEquipsOnFeet     ItemEquipsOn = "feet"
	ItemEquipsOnFinger   ItemEquipsOn = "finger" // maximum 2 per hero

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
