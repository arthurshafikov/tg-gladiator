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

	PotionEffectTypeHeal        = "heal"
	PotionEffectTypeAttackBuff  = "attack_buff"
	PotionEffectTypeDefenseBuff = "defense_buff"
	PotionEffectTypeEvasionBuff = "evasion_buff"
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
