package main

import (
	"autorpg/combat"
	"autorpg/person/character"
	"autorpg/person/enemy"
	"autorpg/utils"
	"fmt"
)

func main() {
	utils.LoadEnvs()

	utils.PrintIntro()
	if utils.DEBUG != "True" {
		utils.AwaitInput()
	}

	combatEngine := combat.NewCombat()

	character := character.NewCharacter()
	character.Create()

	fmt.Print("Generating enemy...\n")
	enemy := enemy.NewEnemy(character.GetPerson().Level)
	fmt.Print(enemy)

	combatEngine.InitiateCombat(character, enemy)
}
