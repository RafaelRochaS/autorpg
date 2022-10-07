package enemy

import (
	"github.com/RafaelRochaS/autorpg/item"
	"github.com/RafaelRochaS/autorpg/person"
)

type Enemy interface {
	GetName() string
	GetLevel() int
	GetWeapon() item.Weapon
	GetArmor() item.Armor
	GetXpGiven() int
	GetHP() int
	GetDropWeapon(int) item.Weapon
	GetDropArmor(int) item.Armor
	String() string
	GetPerson() person.Person
	TakeDamage(int)
}
