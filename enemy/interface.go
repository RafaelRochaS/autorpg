package enemy

import "autorpg/item"

type Enemy interface {
	GetName() string
	GetLevel() int
	GetWeapon() item.Weapon
	GetArmor() item.Armor
	GetXpGiven() int
	GetHP() int
}
