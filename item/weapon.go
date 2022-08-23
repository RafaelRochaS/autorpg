package item

import "fmt"

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

func (w WeaponImpl) String() string {
	return fmt.Sprintf(`%s
Damage Type: %v
Damage: %d,
Attack Speed: %f
Str Req: %d
Dex Req: %d
Int Req: %d
`, w.GetName(), w.DamageType, w.Damage, w.AttackSpeed, w.GetStrReq(), w.GetDexReq(), w.GetIntReq())
}
