package main

import (
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

	character := character.NewCharacter()
	character.Create()

	enemy := enemy.NewEnemy(character.GetPerson().Level)
	fmt.Print(enemy)
}
