package enemy

import (
	"autorpg/item"
	"autorpg/person"
	"fmt"
)

type EnemyImpl struct {
	Person      person.Person
	HP          int
	XpGiven     int
	PlayerLevel int
}

func NewEnemy(playerLevel int) Enemy {
	enemy := &EnemyImpl{}
	enemy.PlayerLevel = playerLevel

	enemy.instantiate()

	return enemy
}

func (e *EnemyImpl) instantiate() {
	e.Person.Level = getLevelBasedOnPlayer(e.PlayerLevel)
	e.XpGiven = getXpGiven(e.Person.Level, e.PlayerLevel)
	e.HP = getHP(e.Person.Level)
	e.Person.Name = getRandomNameByLevel(e.Person.Level)
	e.Person.Weapon = getWeaponByLevel(e.Person.Level)
	e.Person.Armor = getArmorByLevel(e.Person.Level)
}

func (e EnemyImpl) GetName() string {
	return e.Person.Name
}

func (e EnemyImpl) GetLevel() int {
	return e.Person.Level
}

func (e EnemyImpl) GetWeapon() item.Weapon {
	return e.Person.Weapon
}

func (e EnemyImpl) GetArmor() item.Armor {
	return e.Person.Armor
}

func (e EnemyImpl) GetXpGiven() int {
	return e.XpGiven
}

func (e EnemyImpl) GetHP() int {
	return e.HP
}

func (e EnemyImpl) GetDrop() item.Item {

}

func (e EnemyImpl) String() string {
	str := "Current Enemy:\n"
	str += fmt.Sprintf("Name: %d\n", e.GetName())
	str += fmt.Sprintf("Level: %d\n", e.GetLevel())

	return str
}
