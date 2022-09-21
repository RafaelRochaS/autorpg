package character

import (
	"autorpg/item"
	"autorpg/person"
)

type Character interface {
	Create() error
	SetName(string)
	GetClass() Class
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
	HandleWeaponDrop(item.Weapon)
	HandleArmorDrop(item.Armor)
	ResetHP()
	GetLuck() int
	GetCurrentXp() int
}
