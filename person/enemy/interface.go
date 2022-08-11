package enemy

import (
	"autorpg/item"
	"autorpg/person"
)

type Enemy interface {
	GetName() string
	GetLevel() int
	GetWeapon() item.Weapon
	GetArmor() item.Armor
	GetXpGiven() int
	GetHP() int
	GetDrop() item.Item
	String() string
	GetPerson() person.Person
	TakeDamage(int)
}
