package main

import (
	"autorpg/character"
	"autorpg/utils"
)

func main() {
	utils.LoadEnvs()

	utils.PrintIntro()
	if utils.DEBUG != "True" {
		utils.AwaitInput()
	}

	character := character.NewCharacter()
	character.Create()
}
