package character

import "autorpg/item"

type Class int

const (
	WARRIOR Class = iota
	ROGUE
	WIZARD
	BARBARIAN
)

const MAX_STARTING_STATS = 45

/*** WARRIOR Defaults **/
const (
	WAR_HP       = 40
	WAR_STR      = 9
	WAR_STR_UP   = 3
	WAR_CONST    = 10
	WAR_CONST_UP = 2
	WAR_DEX      = 7
	WAR_DEX_UP   = 2
	WAR_INT      = 4
	WAR_INT_UP   = 1
	WAR_LUCK     = 5
	WAR_LUCK_UP  = 2
)

func GetWarriorDefaults() (item.Weapon, item.Armor) {
	warWeapon := item.WeaponImpl{
		Item: item.ItemImpl{
			Name:   "Basic Sword",
			Level:  1,
			Type:   item.WEAPON,
			StrReq: 1,
			DexReq: 1,
			IntReq: 1,
		},
		DamageType:  item.NORMAL,
		Damage:      5,
		AttackSpeed: 1.6,
	}

	warArmor := item.ArmorImpl{
		Item: item.ItemImpl{
			Name:   "Basic Plate",
			Level:  1,
			Type:   item.ARMOR,
			StrReq: 1,
			DexReq: 1,
			IntReq: 1,
		},
		Defense: 7,
		Weight:  3,
	}

	return warWeapon, warArmor
}

/*** ROGUE Defaults **/
const (
	ROG_HP       = 36
	ROG_STR      = 4
	ROG_STR_UP   = 1
	ROG_CONST    = 8
	ROG_CONST_UP = 2
	ROG_DEX      = 12
	ROG_DEX_UP   = 3
	ROG_INT      = 5
	ROG_INT_UP   = 1
	ROG_LUCK     = 6
	ROG_LUCK_UP  = 2
)

func GetRogueDefaults() (item.Weapon, item.Armor) {
	warWeapon := item.WeaponImpl{
		Item: item.ItemImpl{
			Name:   "Basic Dagger",
			Level:  1,
			Type:   item.WEAPON,
			StrReq: 1,
			DexReq: 1,
			IntReq: 1,
		},
		DamageType:  item.NORMAL,
		Damage:      7,
		AttackSpeed: 2.8,
	}

	warArmor := item.ArmorImpl{
		Item: item.ItemImpl{
			Name:   "Basic Leather",
			Level:  1,
			Type:   item.ARMOR,
			StrReq: 1,
			DexReq: 1,
			IntReq: 1,
		},
		Defense: 9,
		Weight:  2,
	}

	return warWeapon, warArmor
}

/*** WIZARD Defaults **/
const (
	WIZ_HP       = 33
	WIZ_STR      = 3
	WIZ_STR_UP   = 1
	WIZ_CONST    = 5
	WIZ_CONST_UP = 2
	WIZ_DEX      = 3
	WIZ_DEX_UP   = 1
	WIZ_INT      = 15
	WIZ_INT_UP   = 4
	WIZ_LUCK     = 9
	WIZ_LUCK_UP  = 2
)

func GetWizardDefaults() (item.Weapon, item.Armor) {
	warWeapon := item.WeaponImpl{
		Item: item.ItemImpl{
			Name:   "Basic Staff",
			Level:  1,
			Type:   item.WEAPON,
			StrReq: 1,
			DexReq: 1,
			IntReq: 1,
		},
		DamageType:  item.MAGICAL,
		Damage:      3,
		AttackSpeed: 1.6,
	}

	warArmor := item.ArmorImpl{
		Item: item.ItemImpl{
			Name:   "Basic Robe",
			Level:  1,
			Type:   item.ARMOR,
			StrReq: 1,
			DexReq: 1,
			IntReq: 1,
		},
		Defense: 5,
		Weight:  1,
	}

	return warWeapon, warArmor
}

/*** BARBARIAN Defaults **/
const (
	BAR_HP       = 44
	BAR_STR      = 13
	BAR_STR_UP   = 4
	BAR_CONST    = 13
	BAR_CONST_UP = 3
	BAR_DEX      = 2
	BAR_DEX_UP   = 1
	BAR_INT      = 2
	BAR_INT_UP   = 1
	BAR_LUCK     = 5
	BAR_LUCK_UP  = 1
)

func GetBarbarianDefaults() (item.Weapon, item.Armor) {
	warWeapon := item.WeaponImpl{
		Item: item.ItemImpl{
			Name:   "Basic Club",
			Level:  1,
			Type:   item.WEAPON,
			StrReq: 1,
			DexReq: 1,
			IntReq: 1,
		},
		DamageType:  item.NORMAL,
		Damage:      8,
		AttackSpeed: 1,
	}

	warArmor := item.ArmorImpl{
		Item: item.ItemImpl{
			Name:   "Basic Chainmail",
			Level:  1,
			Type:   item.ARMOR,
			StrReq: 1,
			DexReq: 1,
			IntReq: 1,
		},
		Defense: 8,
		Weight:  3,
	}

	return warWeapon, warArmor
}

func (c Class) String() string {
	switch c {
	case WARRIOR:
		return "Warrior"
	case ROGUE:
		return "Rogue"
	case WIZARD:
		return "Wizard"
	case BARBARIAN:
		return "Barbarian"
	}
	return "Unknown"
}
