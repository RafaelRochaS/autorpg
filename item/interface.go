package item

type Item interface {
	GetLevel() int
	GetName() string
	GetType() ItemType
	GetStrReq() int
	GetIntReq() int
	GetDexReq() int
}

type Armor interface {
	Item
	GetDefense() int
	GetWeight() int
}

type Weapon interface {
	Item
	GetDamage() int
	GetDamageType() DamageType
	GetAttackSpeed() float32
}
