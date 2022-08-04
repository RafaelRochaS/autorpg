package person

import "autorpg/item"

type Person struct {
	Name   string
	Level  int
	Weapon item.Weapon
	Armor  item.Armor
}
