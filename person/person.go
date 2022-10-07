package person

import "github.com/RafaelRochaS/autorpg/item"

type Person struct {
	Name   string
	Level  int
	Weapon item.Weapon
	Armor  item.Armor
}

func (p *Person) SetWeapon(w item.Weapon) {
	p.Weapon = w
}

func (p *Person) SetArmor(a item.Armor) {
	p.Armor = a
}
