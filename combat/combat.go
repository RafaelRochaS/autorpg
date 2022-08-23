package combat

import (
	"autorpg/person/character"
	"autorpg/person/enemy"
	stringsRPG "autorpg/strings"
	"autorpg/utils"
	"fmt"
	"os"
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

	if c.playerDPS <= 0 {
		c.playerDPS = 1
	}

	if c.enemyDPS <= 0 {
		c.enemyDPS = 1
	}

	c.executeCombat()
}

// TODO: Implement magic weapon mechanic
func (c *CombatImpl) executeCombat() {
	c.printStartCombat()

	if utils.DEBUG == "True" {
		fmt.Print("[***DEBUG] ")
		fmt.Printf("Player DPS: %f\tEnemy DPS: %f\n", c.playerDPS, c.enemyDPS)
	}

	for c.enemy.GetHP() > 0 && c.player.GetHP() > 0 {
		c.enemy.TakeDamage(int(c.playerDPS))
		if utils.DEBUG == "True" {
			fmt.Print("[***DEBUG] ")
			fmt.Printf("%s took %d points of damage. Remaining HP: %d\n", c.enemy.GetName(), int(c.playerDPS), c.enemy.GetHP())
		}

		utils.AwaitInput()

		c.player.TakeDamage(int(c.enemyDPS))
		if utils.DEBUG == "True" {
			fmt.Print("[***DEBUG] ")
			fmt.Printf("%s took %d points of damage. Remaining HP: %d\n", c.player.GetPerson().Name, int(c.enemyDPS), c.player.GetHP())
		}

		utils.AwaitInput()
	}

	if c.enemy.GetHP() <= 0 {
		c.handleCombatVictory()
	} else {
		c.handleCombatDefeat()
	}
}

func (c CombatImpl) printStartCombat() {
	fmt.Print(stringsRPG.Separator)
	fmt.Printf("Initiating combat...\n")
	fmt.Printf("%s vs. %s!!\n", c.player.GetPerson().Name, c.enemy.GetName())
	fmt.Print("Fight!\n")
}

func (c *CombatImpl) handleCombatVictory() {
	c.handleDrop()
	c.player.IncreaseXP(c.enemy.GetXpGiven())
	c.player.LevelUp()
}

func (c CombatImpl) handleDrop() {
	num := utils.GetRandomNumberInRange(0, 1)
	if num != 0 {
		drop := c.enemy.GetDropWeapon()
		c.player.HandleWeaponDrop(drop)
		fmt.Printf("\n%s defeated! Got drop: %s\n", c.enemy.GetName(), drop)
	} else {
		drop := c.enemy.GetDropArmor()
		c.player.HandleArmorDrop(drop)
		fmt.Printf("\n%s defeated! Got drop: %s\n", c.enemy.GetName(), drop)
	}
}

func (c CombatImpl) handleCombatDefeat() {
	fmt.Print(stringsRPG.Separator)
	fmt.Printf("Tough luck! %s defeated you!\n\n", c.enemy.GetName())
	fmt.Print("That's that, better luck next time!\n")
	os.Exit(0)
}
