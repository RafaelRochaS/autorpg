package item

type ItemType int

const (
	WEAPON ItemType = iota
	ARMOR
)

func (it ItemType) String() string {
	switch it {
	case WEAPON:
		return "Weapon"
	case ARMOR:
		return "Armor"
	}
	return "Unknown"
}
