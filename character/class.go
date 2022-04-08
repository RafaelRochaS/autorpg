package character

type CLASS int

const (
	WARRIOR CLASS = iota
	ROGUE
	WIZARD
	BARBARIAN
)

const MAX_STARTING_STATS = 45

/*** WARRIOR Defaults **/
const (
	WAR_HP       = 20
	WAR_STR      = 9
	WAR_STR_UP   = 3
	WAR_CONST    = 10
	WAR_CONST_UP = 2
	WAR_DEX      = 7
	WAR_DEX_UP   = 2
	WAR_INT      = 4
	WAR_INT_UP   = 1
	WAR_LUCK     = 5
	WAR_LUCK_UP  = 2
)

/*** ROGUE Defaults **/
const (
	ROG_HP       = 16
	ROG_STR      = 4
	ROG_STR_UP   = 1
	ROG_CONST    = 8
	ROG_CONST_UP = 2
	ROG_DEX      = 12
	ROG_DEX_UP   = 3
	ROG_INT      = 5
	ROG_INT_UP   = 1
	ROG_LUCK     = 6
	ROG_LUCK_UP  = 2
)

/*** WIZARD Defaults **/
const (
	WIZ_HP       = 13
	WIZ_STR      = 3
	WIZ_STR_UP   = 0
	WIZ_CONST    = 5
	WIZ_CONST_UP = 2
	WIZ_DEX      = 3
	WIZ_DEX_UP   = 1
	WIZ_INT      = 15
	WIZ_INT_UP   = 4
	WIZ_LUCK     = 9
	WIZ_LUCK_UP  = 3
)

/*** BARBARIAN Defaults **/
const (
	BAR_HP       = 24
	BAR_STR      = 13
	BAR_STR_UP   = 4
	BAR_CONST    = 13
	BAR_CONST_UP = 4
	BAR_DEX      = 2
	BAR_DEX_UP   = 1
	BAR_INT      = 2
	BAR_INT_UP   = 0
	BAR_LUCK     = 5
	BAR_LUCK_UP  = 1
)

func (c CLASS) String() string {
	switch c {
	case WARRIOR:
		return "Warrior"
	case ROGUE:
		return "Rogue"
	case WIZARD:
		return "Wizard"
	case BARBARIAN:
		return "Barbarian"
	}
	return "Unknown"
}
