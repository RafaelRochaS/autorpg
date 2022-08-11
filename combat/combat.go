package combat

import (
	"autorpg/person/character"
	"autorpg/person/enemy"
	stringsRPG "autorpg/strings"
	"autorpg/utils"
	"fmt"
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
	c.playerDPS = getBaseDPS(c.player.GetPerson()) - float32(c.enemy.GetArmor().GetDefense())
	c.enemyDPS = getBaseDPS(c.enemy.GetPerson()) - float32(c.player.GetArmor().GetDefense())

	c.executeCombat()
}

func (c *CombatImpl) executeCombat() {
	c.printStartCombat()

	if utils.DEBUG == "True" {
		fmt.Print("[***DEBUG] ")
		fmt.Print("Starting with player attacking only, until enemy dies\n")
		fmt.Printf("Player DPS: %f\tEnemy DPS: %f\n", c.playerDPS, c.enemyDPS)
	}

	for c.enemy.GetHP() > 0 {
		c.enemy.TakeDamage(int(c.playerDPS))
		if utils.DEBUG == "True" {
			fmt.Print("[***DEBUG] ")
			fmt.Printf("%s took %d points of damage. Remaining HP: %d\n", c.enemy.GetName(), int(c.playerDPS), c.enemy.GetHP())
		}
	}

	c.handleCombatVictory()
}

func (c CombatImpl) printStartCombat() {
	fmt.Print(stringsRPG.Separator)
	fmt.Printf("Initiating combat...\n")
	fmt.Printf("%s vs. %s!!\n", c.player.GetPerson().Name, c.enemy.GetName())
	fmt.Print("Fight!\n")
}

func (c *CombatImpl) handleCombatVictory() {
	drop := c.enemy.GetDrop()
	fmt.Printf("\n%s defeated! Got drop: %s\n", c.enemy.GetName(), drop.GetName())

	c.player.IncreaseXP(c.enemy.GetXpGiven())
	c.player.LevelUp()
}
