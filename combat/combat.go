package combat

import (
	"fmt"
	"os"

	stringsRPG "github.com/RafaelRochaS/autorpg/strings"

	"github.com/RafaelRochaS/autorpg/person/enemy"

	"github.com/RafaelRochaS/autorpg/person/character"

	"github.com/RafaelRochaS/autorpg/utils"
)

type CombatImpl struct {
	player    character.Character
	enemy     enemy.Enemy
	playerDPS float32
	enemyDPS  float32
}

func NewCombat() Combat {
	return &CombatImpl{}
}

func (c *CombatImpl) InitiateCombat(player character.Character, enemy enemy.Enemy) {
	c.player = player
	c.enemy = enemy

	calculateDps(c)
	c.executeCombat()
}

func (c *CombatImpl) executeCombat() {
	printStartCombat(*c)

	if utils.DEBUG == "True" {
		fmt.Print(stringsRPG.Debug)
		fmt.Printf("Player DPS: %f\tEnemy DPS: %f\n", c.playerDPS, c.enemyDPS)
	}

	for c.enemy.GetHP() > 0 && c.player.GetHP() > 0 {
		c.enemy.TakeDamage(int(c.playerDPS))
		fmt.Printf("%s took %d points of damage. Remaining HP: %d\n", c.enemy.GetName(), int(c.playerDPS), c.enemy.GetHP())

		if c.enemy.GetHP() <= 0 {
			c.handleCombatVictory()
			return
		}

		c.player.TakeDamage(int(c.enemyDPS))
		fmt.Printf("%s took %d points of damage. Remaining HP: %d\n", c.player.GetPerson().Name, int(c.enemyDPS), c.player.GetHP())

		if c.player.GetHP() <= 0 {
			c.handleCombatDefeat()
			return
		}
	}
}

func (c *CombatImpl) handleCombatVictory() {
	c.handleDrop()
	c.player.IncreaseXP(c.enemy.GetXpGiven())
	c.player.LevelUp()
	c.player.ResetHP()
	afterCombatPause()
}

func (c CombatImpl) handleDrop() {
	num := utils.GetRandomNumberInRange(0, 1)
	if num != 0 {
		drop := c.enemy.GetDropWeapon(c.player.GetLuck())
		fmt.Printf("\n%s defeated! Got drop: %s\n", c.enemy.GetName(), drop)
		c.player.HandleWeaponDrop(drop)
	} else {
		drop := c.enemy.GetDropArmor(c.player.GetLuck())
		fmt.Printf("\n%s defeated! Got drop: %s\n", c.enemy.GetName(), drop)
		c.player.HandleArmorDrop(drop)
	}
}

func (c CombatImpl) handleCombatDefeat() {
	fmt.Print(stringsRPG.Separator)
	fmt.Printf("Tough luck! %s defeated you!\n\n", c.enemy.GetName())
	fmt.Print("That's that, better luck next time!\n")
	os.Exit(0)
}
