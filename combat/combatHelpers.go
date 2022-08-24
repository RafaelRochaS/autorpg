package combat

import (
	"autorpg/item"
	"autorpg/person"
	stringsRPG "autorpg/strings"
	"autorpg/utils"
	"fmt"
	"time"
)

func calculateDps(c *CombatImpl) {
	c.playerDPS = getBaseDPS(c.player.GetPerson())
	c.enemyDPS = getBaseDPS(c.enemy.GetPerson())

	if c.player.GetWeapon().GetDamageType() == item.NORMAL {
		c.playerDPS -= float32(c.enemy.GetArmor().GetDefense())
	}

	if c.enemy.GetWeapon().GetDamageType() == item.NORMAL {
		c.enemyDPS -= float32(c.player.GetArmor().GetDefense())
	}

	if c.playerDPS <= 0 {
		c.playerDPS = 1
	}

	if c.enemyDPS <= 0 {
		c.enemyDPS = 1
	}
}

func getBaseDPS(p person.Person) float32 {
	return (float32(p.Weapon.GetDamage()) * p.Weapon.GetAttackSpeed()) - (float32(p.Armor.GetWeight()) / 3)
}

func printStartCombat(c CombatImpl) {
	fmt.Print(stringsRPG.Separator)
	fmt.Printf("Initiating combat...\n")
	fmt.Printf("%s vs. %s!!\n", c.player.GetPerson().Name, c.enemy.GetName())
	fmt.Print("Fight!\n")
}

func afterCombatPause() {
	time.Sleep(time.Duration(utils.COMBAT_PAUSE_TIME) * time.Second)
}
