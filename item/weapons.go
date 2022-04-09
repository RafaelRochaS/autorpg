package item

type WeaponImpl struct {
	Item
	DamageType DamageType
	Damage     int
}

func (w WeaponImpl) GetDamage() int {
	return w.Damage
}

func (w WeaponImpl) GetDamageType() DamageType {
	return w.DamageType
}
