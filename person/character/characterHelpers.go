package character

import (
	stringsRPG "autorpg/strings"
	"autorpg/utils"
	"fmt"
	"log"
)

func printCreation() {
	utils.ClearScreen()
	fmt.Println(stringsRPG.CharCreation)
	fmt.Print(stringsRPG.Separator)
	fmt.Println(stringsRPG.StatsScreen)
	fmt.Print(stringsRPG.Separator)
}

func printName() {
	fmt.Print(stringsRPG.SetCharName)
}

func printChooseClass() {
	fmt.Println()
	fmt.Print(stringsRPG.ChooseClass)
	fmt.Println()
}

func printChooseStats() {
	fmt.Print(stringsRPG.SetStats)
}

func printCurrentStats(c CharacterImpl) {
	fmt.Print(stringsRPG.Separator)
	fmt.Println("Current Stats:")
	fmt.Printf(`
HP: %d
Level: %d
Strength: %d (+%d)
Const: %d (+%d)
Dexterity: %d (+%d)
Intelligence: %d (+%d)
Luck: %d (+%d)
`,
		c.Stats.HP,
		c.Person.Level,
		c.Stats.Str,
		c.Stats.Str_UP,
		c.Stats.Const,
		c.Stats.Const_UP,
		c.Stats.Dex,
		c.Stats.Dex_UP,
		c.Stats.Int,
		c.Stats.Int_UP,
		c.Stats.Luck,
		c.Stats.Luck_UP)
}

func printAddPoints() {
	fmt.Print(stringsRPG.Separator)
	fmt.Println("Now you have 10 total points to distribute amongst your stats:")
}

func printChar(c CharacterImpl) {
	fmt.Println("Done! Your character is created:")
	fmt.Printf(`
Name: %v
Class: %v`, c.Person.Name, c.Class)

	printCurrentStats(c)
}

func printGear(c CharacterImpl) {
	fmt.Print(stringsRPG.Separator)
	fmt.Print("\nCurrently equipped gear:\n")
	fmt.Printf(`
Weapon:
%v
Level: %d
Damage: %d
Damage Type: %v
Attack Speed: %.2f`, c.Person.Weapon.GetName(), c.Person.Weapon.GetLevel(), c.Person.Weapon.GetDamage(), c.Person.Weapon.GetDamageType(), c.Person.Weapon.GetAttackSpeed())
	fmt.Println()
	fmt.Printf(`
Armor:
%v
Level: %d
Defense: %d
Weight: %d

`, c.Person.Armor.GetName(), c.Person.Armor.GetLevel(), c.Person.Armor.GetDefense(), c.Person.Armor.GetWeight())
}

func checkForErrors(errArr []error) error {
	log.Printf("Check For Errors :: array length: '%d'", len(errArr))
	for i := 0; i < len(errArr); i++ {
		log.Printf("Check For Errors :: current item: '%v'", errArr[i])
		if errArr[i] != nil {
			log.Printf("Check For Errors :: got error on index '%d'", i)
			return errArr[i]
		}
	}

	return nil
}
