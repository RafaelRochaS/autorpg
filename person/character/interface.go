package character

import (
	"autorpg/item"
	"autorpg/person"
)

type Character interface {
	Create()
	SetName(string)
	SetStats()
	SetClass(Class)
	AttachWeapon(item.Weapon)
	AttachArmor(item.Armor)
	RemoveItem(item.ItemType)
	CheckLevelItem(item.Item) bool
	CheckStatRequirementsItem(item.Item) bool
	LevelUp()
	GetPerson() person.Person
	GetArmor() item.Armor
	GetWeapon() item.Weapon
	IncreaseXP(int)
	TakeDamage(int)
	GetHP() int
}
