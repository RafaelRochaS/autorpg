package combat

import (
	"github.com/RafaelRochaS/autorpg/person/character"
	"github.com/RafaelRochaS/autorpg/person/enemy"
)

type Combat interface {
	InitiateCombat(character.Character, enemy.Enemy)
}
