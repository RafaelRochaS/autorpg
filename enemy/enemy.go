package enemy

import "autorpg/person"

type EnemyImpl struct {
	Person  person.Person
	XpGiven int
}

func NewEnemy() EnemyImpl {
	return &EnemyImpl{}
}
