package main

import (
	"fmt"
	"log"
	"os"

	"github.com/RafaelRochaS/autorpg/combat"
	"github.com/RafaelRochaS/autorpg/person/character"
	"github.com/RafaelRochaS/autorpg/person/enemy"
	"github.com/RafaelRochaS/autorpg/utils"
)

func main() {
	utils.LoadEnvs()

	utils.PrintIntro()
	if utils.DEBUG != "True" {
		utils.AwaitInput()
	}

	combatEngine := combat.NewCombat()

	character := character.NewCharacter(os.Stdin)
	err := character.Create()
	if err != nil {
		log.Printf("Error creating character: %s", err.Error())
		log.Fatal("Exiting...")
	}

	for {
		fmt.Print("Generating enemy...\n")
		enemy := enemy.NewEnemy(character.GetPerson().Level)
		fmt.Print(enemy)

		combatEngine.InitiateCombat(character, enemy)
	}
}
