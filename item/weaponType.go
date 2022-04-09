package item

type DamageType int

const (
	NORMAL DamageType = iota
	MAGICAL
)

func (dt DamageType) String() string {
	switch dt {
	case NORMAL:
		return "Normal"
	case MAGICAL:
		return "Magical"
	}
	return "Unknown"
}
