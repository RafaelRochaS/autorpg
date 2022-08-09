package character

import (
	"autorpg/item"
	"autorpg/person"
	stringsRPG "autorpg/strings"
	"autorpg/utils"
	"fmt"
	"math"
)

type CharacterImpl struct {
	Person    person.Person
	Stats     Stats
	CurrentXp int
	Class     Class
}

type Stats struct {
	HP       int
	Str      int
	Const    int
	Dex      int
	Int      int
	Luck     int
	Str_UP   int
	Const_UP int
	Dex_UP   int
	Int_UP   int
	Luck_UP  int
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
	fmt.Printf(`
HP: %d
Strength: %d (+%d)
Const: %d (+%d)
Dexterity: %d (+%d)
Intelligence: %d (+%d)
Luck: %d (+%d)`,
		c.Stats.HP,
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
Defense: %d`, c.Person.Armor.GetName(), c.Person.Armor.GetLevel(), c.Person.Armor.GetDefense())
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

	fmt.Printf("\nCool, '%v', that's a good name, I guess. I don't really know.\n", c.Person.Name)

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

	printAddPoints()
	c.AddPoints()

	printChar(*c)

	var weapon item.Weapon
	var armor item.Armor

	if utils.DEBUG != "True" {
		switch c.Class {
		case WARRIOR:
			weapon, armor = GetWarriorDefaults()
		case ROGUE:
			weapon, armor = GetRogueDefaults()
		case WIZARD:
			weapon, armor = GetWizardDefaults()
		case BARBARIAN:
			weapon, armor = GetBarbarianDefaults()
		}
	} else {
		weapon, armor = GetWizardDefaults()
	}

	c.AttachWeapon(weapon)
	c.AttachArmor(armor)

	printGear(*c)
}

func (c *CharacterImpl) SetName(name string) {
	c.Person.Name = name
}

func (c *CharacterImpl) SetStats() {
	switch c.Class {
	case WARRIOR:
		c.Stats.HP = WAR_HP
		c.Stats.Str = WAR_STR
		c.Stats.Str_UP = WAR_STR_UP
		c.Stats.Const = WAR_CONST
		c.Stats.Const_UP = WAR_CONST_UP
		c.Stats.Dex = WAR_DEX
		c.Stats.Dex_UP = WAR_DEX_UP
		c.Stats.Int = WAR_INT
		c.Stats.Int_UP = WAR_INT_UP
		c.Stats.Luck = WAR_LUCK
		c.Stats.Luck_UP = WAR_LUCK_UP

	case ROGUE:
		c.Stats.HP = ROG_HP
		c.Stats.Str = ROG_STR
		c.Stats.Str_UP = ROG_STR_UP
		c.Stats.Const = ROG_CONST
		c.Stats.Const_UP = ROG_CONST_UP
		c.Stats.Dex = ROG_DEX
		c.Stats.Dex_UP = ROG_DEX_UP
		c.Stats.Int = ROG_INT
		c.Stats.Int_UP = ROG_INT_UP
		c.Stats.Luck = ROG_LUCK
		c.Stats.Luck_UP = ROG_LUCK_UP

	case WIZARD:
		c.Stats.HP = WIZ_HP
		c.Stats.Str = WIZ_STR
		c.Stats.Str_UP = WIZ_STR_UP
		c.Stats.Const = WIZ_CONST
		c.Stats.Const_UP = WIZ_CONST_UP
		c.Stats.Dex = WIZ_DEX
		c.Stats.Dex_UP = WIZ_DEX_UP
		c.Stats.Int = WIZ_INT
		c.Stats.Int_UP = WIZ_INT_UP
		c.Stats.Luck = WIZ_LUCK
		c.Stats.Luck_UP = WIZ_LUCK_UP

	case BARBARIAN:
		c.Stats.HP = BAR_HP
		c.Stats.Str = BAR_STR
		c.Stats.Str_UP = BAR_STR_UP
		c.Stats.Const = BAR_CONST
		c.Stats.Const_UP = BAR_CONST_UP
		c.Stats.Dex = BAR_DEX
		c.Stats.Dex_UP = BAR_DEX_UP
		c.Stats.Int = BAR_INT
		c.Stats.Int_UP = BAR_INT_UP
		c.Stats.Luck = BAR_LUCK
		c.Stats.Luck_UP = BAR_LUCK_UP
	}

	printCurrentStats(*c)
}

func (c *CharacterImpl) SetClass(class Class) { // Assumes class is a valid int that belongs to an existing class
	c.Class = class
}

func (c *CharacterImpl) AddPoints() {
	var total int

	if utils.DEBUG != "True" {
		for {
			fmt.Println("How many points to add to strength?")
			str := utils.ReadInt()
			total += str
			fmt.Println("How many points to add to Constitution?")
			consti := utils.ReadInt()
			total += consti
			fmt.Println("How many points to add to Dexterity?")
			dex := utils.ReadInt()
			total += dex
			fmt.Println("How many points to add to Intelligence?")
			intel := utils.ReadInt()
			total += intel
			fmt.Println("How many points to add to Luck?")
			luck := utils.ReadInt()
			total += luck

			if total > 10 {
				fmt.Println("Too many points added, try again.")
				continue
			}

			c.Stats.Str += str
			c.Stats.Const += consti
			c.Stats.Dex += dex
			c.Stats.Int += intel
			c.Stats.Luck += luck
			break
		}
	} else {
		fmt.Println("\n***[DEBUG] mode enabled, choosing default stat points...")
		c.Stats.Str += 0
		c.Stats.Const += 3
		c.Stats.Dex += 1
		c.Stats.Int += 5
		c.Stats.Luck += 1
	}
}

func (c *CharacterImpl) AttachWeapon(w item.Weapon) {
	c.Person.Weapon = w
}

func (c *CharacterImpl) AttachArmor(w item.Armor) {
	c.Person.Armor = w
}

func (c *CharacterImpl) RemoveItem(t item.ItemType) {
	switch t {
	case item.ARMOR:
		c.Person.Armor = &item.ArmorImpl{}
	case item.WEAPON:
		c.Person.Weapon = &item.WeaponImpl{}
	}
}

func (c CharacterImpl) CheckLevelItem(i item.Item) bool {
	return c.Person.Level >= i.GetLevel()
}

func (c CharacterImpl) CheckStatRequirementsItem(i item.Item) bool {
	return c.Stats.Str >= i.GetStrReq() &&
		c.Stats.Dex >= i.GetDexReq() &&
		c.Stats.Int >= i.GetIntReq()
}

func (c *CharacterImpl) LevelUp() {
	xpNeeded := int(math.Pow(float64(c.Person.Level*5), 2) + 65)

	if c.CurrentXp >= xpNeeded {
		c.Person.Level++
		c.CurrentXp = 0
	}
}

func (c CharacterImpl) GetPerson() person.Person {
	return c.Person
}