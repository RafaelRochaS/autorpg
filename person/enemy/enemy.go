package enemy

import (
	"fmt"

	"github.com/RafaelRochaS/autorpg/item"
	"github.com/RafaelRochaS/autorpg/person"
	stringsRPG "github.com/RafaelRochaS/autorpg/strings"
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
	e.Person.Weapon = getWeaponByLevel(e.Person.Level, 0)
	e.Person.Armor = getArmorByLevel(e.Person.Level, 0)
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

func (e EnemyImpl) GetDropWeapon(luck int) item.Weapon {
	return getWeaponByLevel(e.PlayerLevel, luck)
}

func (e EnemyImpl) GetDropArmor(luck int) item.Armor {
	return getArmorByLevel(e.PlayerLevel, luck)
}

func (e EnemyImpl) String() string {
	fmt.Print(stringsRPG.Separator)
	str := "Current Enemy:\n"
	str += "-------------------------------------------------\n"
	str += fmt.Sprintf("Name: %s\n", e.GetName())
	str += fmt.Sprintf("Level: %d\n", e.GetLevel())
	str += fmt.Sprintf("HP: %d\n", e.GetHP())
	str += fmt.Sprintf("XP Given: %d\n", e.GetXpGiven())
	str += "-------------------------------------------------\n"
	str += fmt.Sprintf("Weapon: %s\n", e.Person.Weapon.GetName())
	str += fmt.Sprintf("Level: %d\n", e.Person.Weapon.GetLevel())
	str += fmt.Sprintf("Damage Type: %v\n", e.Person.Weapon.GetDamageType())
	str += fmt.Sprintf("Damage: %v\n", e.Person.Weapon.GetDamage())
	str += fmt.Sprintf("Attack Speed: %v\n", e.Person.Weapon.GetAttackSpeed())
	str += "-------------------------------------------------\n"
	str += fmt.Sprintf("Armor: %s\n", e.Person.Armor.GetName())
	str += fmt.Sprintf("Level: %d\n", e.Person.Armor.GetLevel())
	str += fmt.Sprintf("Defense: %d\n", e.Person.Armor.GetDefense())
	str += fmt.Sprintf("Weight: %d\n", e.Person.Armor.GetWeight())
	str += "-------------------------------------------------\n"

	return str
}

func (e EnemyImpl) GetPerson() person.Person {
	return e.Person
}

func (e *EnemyImpl) TakeDamage(dmg int) {
	e.HP -= dmg
}
