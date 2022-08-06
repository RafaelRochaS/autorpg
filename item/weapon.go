package item

type WeaponImpl struct {
	Item
	DamageType  DamageType
	Damage      int
	AttackSpeed float32
}

func (w WeaponImpl) GetDamage() int {
	return w.Damage
}

func (w WeaponImpl) GetDamageType() DamageType {
	return w.DamageType
}

func (w WeaponImpl) GetAttackSpeed() float32 {
	return w.AttackSpeed
}
