package combat

import (
	"autorpg/person"
)

func getBaseDPS(p person.Person) float32 {
	return (float32(p.Weapon.GetDamage()) * p.Weapon.GetAttackSpeed()) - (float32(p.Armor.GetWeight()) / 3)
}
