package combat

import (
	"autorpg/person/character"
	"autorpg/person/enemy"
)

type Combat interface {
	InitiateCombat(character.Character, enemy.Enemy)
}
