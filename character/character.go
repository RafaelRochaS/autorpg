package character

import (
	stringsRPG "autorpg/strings"
	"autorpg/utils"
	"fmt"
)

type Character interface {
	Create()
	SetName(string)
	SetStats()
	SetClass(CLASS)
}

type CharacterImpl struct {
	Name      string
	Stats     Stats
	Level     int
	CurrentXp int
	Class     CLASS
}

type Stats struct {
	HP    int
	Str   int
	Const int
	Dex   int
	Int   int
	Luck  int
}

/*** Helper Functions **/

func NewCharacter() Character {
	return &CharacterImpl{}
}

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
	fmt.Printf("HP: %d\nStrength: %d\nConst: %d\nDexterity: %d\nIntelligence: %d\nLuck: %d\n",
		c.Stats.HP, c.Stats.Str, c.Stats.Const, c.Stats.Dex, c.Stats.Int, c.Stats.Luck)
}

/*** Character Functions **/

func (c *CharacterImpl) Create() {
	printCreation()
	printName()

	if utils.DEBUG != "True" {
		c.SetName(utils.ReadString())
	} else {
		fmt.Println("\n***[DEBUG] mode enabled, choosing default char name...")
		c.SetName("Joe Do Debug")
	}

	fmt.Printf("\nCool, '%v', that's a good name, I guess. I don't really know.\n", c.Name)

	printChooseClass()
	var class int
	selected := false

	if utils.DEBUG != "True" {
		for {
			if selected {
				break
			}

			fmt.Scanf("%d", &class)

			switch class - 1 {
			case int(WARRIOR):
				c.SetClass(WARRIOR)
				selected = true

			case int(ROGUE):
				c.SetClass(ROGUE)
				selected = true

			case int(WIZARD):
				c.SetClass(WIZARD)
				selected = true

			case int(BARBARIAN):
				c.SetClass(BARBARIAN)
				selected = true

			default:
				fmt.Println("Value not recognized, please try again.")
			}
		}
	} else {
		fmt.Println("\n***[DEBUG] mode enabled, choosing default char class...")
		c.SetClass(WIZARD)
	}

	fmt.Printf("\nSelected class: %v\n", c.Class)

	printChooseStats()
	c.SetStats()
}

func (c *CharacterImpl) SetName(name string) {
	c.Name = name
}

func (c *CharacterImpl) SetStats() {
	switch c.Class {
	case WARRIOR:
		c.Stats.HP = WAR_HP
		c.Stats.Str = WAR_STR
		c.Stats.Const = WAR_CONST
		c.Stats.Dex = WAR_DEX
		c.Stats.Int = WAR_INT
		c.Stats.Luck = WAR_LUCK

	case ROGUE:
		c.Stats.HP = ROG_HP
		c.Stats.Str = ROG_STR
		c.Stats.Const = ROG_CONST
		c.Stats.Dex = ROG_DEX
		c.Stats.Int = ROG_INT
		c.Stats.Luck = ROG_LUCK

	case WIZARD:
		c.Stats.HP = WIZ_HP
		c.Stats.Str = WIZ_STR
		c.Stats.Const = WIZ_CONST
		c.Stats.Dex = WIZ_DEX
		c.Stats.Int = WIZ_INT
		c.Stats.Luck = WIZ_LUCK

	case BARBARIAN:
		c.Stats.HP = BAR_HP
		c.Stats.Str = BAR_STR
		c.Stats.Const = BAR_CONST
		c.Stats.Dex = BAR_DEX
		c.Stats.Int = BAR_INT
		c.Stats.Luck = BAR_LUCK
	}

	printCurrentStats(*c)
}

func (c *CharacterImpl) SetClass(class CLASS) { // Assumes class is a valid int that belongs to an existing class
	c.Class = class
}
