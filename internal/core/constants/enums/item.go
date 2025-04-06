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
	ItemCategoryPotion    ItemCategory = "potion"

	// @todo what's left for potions
	// using them in a fight with a unique effects
	// allowing to have more than 1 potion in the bag (quantity column is possible + also to the shop table?)
	// updating the potions list
	// a separate shop section for potions?
	PotionEffectTypeHeal = "heal"
	// PotionEffectTypeAttackBuff  = "attack_buff"
	// PotionEffectTypeDefenseBuff = "defense_buff"
	// PotionEffectTypeEvasionBuff = "evasion_buff"
)

func ItemEquipsOnAll() []ItemEquipsOn {
	return []ItemEquipsOn{
		ItemEquipsOnHand,
		ItemEquipsOnTwoHands,
		ItemEquipsOnHands,
		ItemEquipsOnBody,
		ItemEquipsOnHead,
		ItemEquipsOnNeck,
		ItemEquipsOnFeet,
		ItemEquipsOnFinger,
	}
}

func ItemCategoryAll() []ItemCategory {
	return []ItemCategory{
		ItemCategoryWeapon,
		ItemCategoryArmor,
		ItemCategoryAccessory,
		ItemCategoryPotion,
	}
}
