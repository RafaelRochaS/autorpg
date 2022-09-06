package character

import (
	"autorpg/item"
	"autorpg/person"
	"autorpg/utils"
	"fmt"
	"io"
	"log"
	"math"
)

type CharacterImpl struct {
	Person      person.Person
	Stats       Stats
	CurrentXp   int
	Class       Class
	damageTaken int
	io          io.Reader
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

func NewCharacter(io io.Reader) Character {
	return &CharacterImpl{
		io: io,
	}
}

/*** Character Functions **/

func (c *CharacterImpl) Create() {
	printCreation()
	printName()

	if utils.DEBUG != "True" {
		str, err := utils.ReadString(c.io)
		if err != nil {
			log.Fatalf("Error getting character name")
		}
		c.SetName(str)
	} else {
		fmt.Println("\n***[DEBUG] mode enabled, choosing default char name...")
		c.SetName("Joe Do Debug")
	}

	fmt.Printf("\nCool, '%v', that's a good name, I guess. I don't really know.\n", c.Person.Name)

	printChooseClass()
	selected := false

	if utils.DEBUG != "True" {
		for {
			if selected {
				break
			}

			class, err := utils.ReadInt(c.io)
			if err != nil {
				fmt.Println("**************************************************************")
				fmt.Println(c.io)
				fmt.Println("**************************************************************")
				log.Fatal(err.Error())
			}

			switch class - 1 {
			case int(WARRIOR):
				c.setClass(WARRIOR)
				selected = true

			case int(ROGUE):
				c.setClass(ROGUE)
				selected = true

			case int(WIZARD):
				c.setClass(WIZARD)
				selected = true

			case int(BARBARIAN):
				c.setClass(BARBARIAN)
				selected = true

			default:
				fmt.Println("Value not recognized, please try again.")
			}
		}
	} else {
		fmt.Println("\n***[DEBUG] mode enabled, choosing default char class...")
		c.setClass(BARBARIAN)
	}

	fmt.Printf("\nSelected class: %v\n", c.Class)

	printChooseStats()
	c.setStats()

	printAddPoints()
	c.addPoints()

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
		weapon, armor = GetBarbarianDefaults()
	}

	c.AttachWeapon(weapon)
	c.AttachArmor(armor)

	printGear(*c)
}

func (c *CharacterImpl) SetName(name string) {
	c.Person.Name = name
}

func (c *CharacterImpl) setStats() {
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

	c.Person.Level = 1

	printCurrentStats(*c)
}

func (c *CharacterImpl) setClass(class Class) { // Assumes class is a valid int that belongs to an existing class
	c.Class = class
}

func (c *CharacterImpl) addPoints() {
	var total int

	if utils.DEBUG != "True" {
		for {
			fmt.Println("How many points to add to strength?")
			str, _ := utils.ReadInt(c.io)
			total += str
			fmt.Println("How many points to add to Constitution?")
			consti, _ := utils.ReadInt(c.io)
			total += consti
			fmt.Println("How many points to add to Dexterity?")
			dex, _ := utils.ReadInt(c.io)
			total += dex
			fmt.Println("How many points to add to Intelligence?")
			intel, _ := utils.ReadInt(c.io)
			total += intel
			fmt.Println("How many points to add to Luck?")
			luck, _ := utils.ReadInt(c.io)
			total += luck

			if total > 10 {
				fmt.Println("Too many points added, try again.")
				total = 0
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
		c.Stats.Str += 5
		c.Stats.Const += 3
		c.Stats.Dex += 1
		c.Stats.Int += 0
		c.Stats.Luck += 1
	}
}

func (c *CharacterImpl) AttachWeapon(w item.Weapon) {
	c.Person.Weapon = w
}

func (c *CharacterImpl) AttachArmor(a item.Armor) {
	c.Person.Armor = a
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
		fmt.Printf("\nCharacter leveled up!\nLevel: %d\n", c.Person.Level)
		c.increaseStatsByLevelUp()
		printCurrentStats(*c)
	} else {
		fmt.Printf("Current XP: %d\nTo Level Up: %d\n", c.CurrentXp, xpNeeded)
	}
}

func (c *CharacterImpl) increaseStatsByLevelUp() {
	switch c.Class {
	case WARRIOR:
		c.Stats.Str += WAR_STR_UP
		c.Stats.Dex += WAR_DEX_UP
		c.Stats.Int += WAR_INT_UP
		c.Stats.Const += WAR_CONST_UP / 3
		c.Stats.Luck += WAR_LUCK_UP
	case ROGUE:
		c.Stats.Str += ROG_STR_UP
		c.Stats.Dex += ROG_DEX_UP
		c.Stats.Int += ROG_INT_UP
		c.Stats.Const += ROG_CONST_UP / 3
		c.Stats.Luck += ROG_LUCK_UP
	case WIZARD:
		c.Stats.Str += WIZ_STR_UP
		c.Stats.Dex += WIZ_DEX_UP
		c.Stats.Int += WIZ_INT_UP
		c.Stats.Const += WIZ_CONST_UP / 3
		c.Stats.Luck += WIZ_LUCK_UP
	case BARBARIAN:
		c.Stats.Str += BAR_STR_UP
		c.Stats.Dex += BAR_DEX_UP
		c.Stats.Int += BAR_INT_UP
		c.Stats.Const += BAR_CONST_UP / 3
		c.Stats.Luck += BAR_LUCK_UP
	}

	c.Stats.HP += c.Stats.Const
}

func (c CharacterImpl) GetPerson() person.Person {
	return c.Person
}

func (c CharacterImpl) GetArmor() item.Armor {
	return c.Person.Armor
}

func (c CharacterImpl) GetWeapon() item.Weapon {
	return c.Person.Weapon
}

func (c *CharacterImpl) IncreaseXP(xp int) {
	c.CurrentXp += xp
}

func (c *CharacterImpl) TakeDamage(dmg int) {
	c.damageTaken += dmg
	c.Stats.HP -= dmg
}

func (c CharacterImpl) GetHP() int {
	return c.Stats.HP
}

func (c *CharacterImpl) HandleArmorDrop(drop item.Armor) {
	if c.canItemBeUsed(drop) {
		if utils.DEBUG == "True" {
			fmt.Print("[***DEBUG] ")
			fmt.Printf("HandleArmorDrop: item can be used")
		}

		if (drop.GetDefense() > c.GetArmor().GetDefense()) &&
			(drop.GetWeight() <= c.GetArmor().GetDefense()) {
			c.Person.SetArmor(drop)
		}

		if (float32(drop.GetDefense()) > (float32(c.GetArmor().GetDefense()) * 1.33)) &&
			(drop.GetWeight() > c.GetArmor().GetWeight()) {
			c.Person.SetArmor(drop)
		}
	} else if utils.DEBUG == "True" {
		fmt.Print("[***DEBUG] ")
		fmt.Printf("HandleArmorDrop: item cannot be used")
	}

	printCurrentStats(*c)
	printGear(*c)
}

func (c *CharacterImpl) HandleWeaponDrop(drop item.Weapon) {
	if c.canItemBeUsed(drop) {
		if utils.DEBUG == "True" {
			fmt.Print("[***DEBUG] ")
			fmt.Printf("HandleWeaponDrop: item can be used\n")
		}

		if (drop.GetDamage() > c.Person.Weapon.GetDamage()) &&
			(drop.GetAttackSpeed() >= c.Person.Weapon.GetAttackSpeed()) {
			c.Person.SetWeapon(drop)
		}

		if (float32(drop.GetDamage()) > (float32(c.Person.Weapon.GetDamage()) * 1.33)) &&
			(drop.GetAttackSpeed() < c.Person.Weapon.GetAttackSpeed()) {
			c.Person.SetWeapon(drop)
		}
	} else if utils.DEBUG == "True" {
		fmt.Print("[***DEBUG] ")
		fmt.Printf("HandleWeaponDrop: item cannot be used\n")
	}

	printCurrentStats(*c)
	printGear(*c)
}

func (c CharacterImpl) canItemBeUsed(item item.Item) bool {
	return (item.GetStrReq() < c.Stats.Str) ||
		(item.GetIntReq() < c.Stats.Int) ||
		(item.GetDexReq() < c.Stats.Int)
}

func (c *CharacterImpl) ResetHP() {
	c.Stats.HP += c.damageTaken
	c.damageTaken = 0
}

func (c CharacterImpl) GetLuck() int {
	return c.Stats.Luck
}

func (c CharacterImpl) GetClass() Class {
	return c.Class
}
