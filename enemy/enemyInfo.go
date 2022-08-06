package enemy

import (
	"autorpg/item"
	"autorpg/utils"
)

const LEVEL_DIFFERENCE = 3

/* Naming Scheme:
Level 1-10: first name level 1 last name level 1
Level 10-20: first name level 2 last name level 1
Level 20-30: first name level 3 last name level 1
Level 30-40: first name level 1 last name level 2 ...
*/

/** NAMES **/
var firstNames_1 = []string{"Foot Soldier", "Foot Archer", "Foot Horsemen", "Foot Wizard"}
var lastNames_1 = []string{"Weak boy", "Trash tier", "No skill", "N00b"}
var firstNames_2 = []string{"Sargeant", "Ranger", "Knight", "Mage"}
var lastNames_2 = []string{"Upcoming", "Kinda good", "Shows promise", "Getting there"}
var firstNames_3 = []string{"Champion", "Legolas", "Charger", "BattleMage"}
var lastNames_3 = []string{"Next Level", "MFer", "Hardcore", "Killer"}

/** WEAPONS **/
var weaponFirstNames_1 = []string{"Foot Soldier", "Foot Archer", "Foot Horsemen", "Foot Wizard"}
var weaponLastNames_1 = []string{"Weak boy", "Trash tier", "No skill", "N00b"}
var weaponFirstNames_2 = []string{"Sargeant", "Ranger", "Knight", "Mage"}
var weaponLastNames_2 = []string{"Upcoming", "Kinda good", "Shows promise", "Getting there"}
var weaponFirstNames_3 = []string{"Champion", "Legolas", "Charger", "BattleMage"}
var weaponLastNames_3 = []string{"Next Level", "MFer", "Hardcore", "Killer"}

/** ARMOR **/
var armorFirstNames_1 = []string{"Foot Soldier", "Foot Archer", "Foot Horsemen", "Foot Wizard"}
var armorLastNames_1 = []string{"Weak boy", "Trash tier", "No skill", "N00b"}
var armorFirstNames_2 = []string{"Sargeant", "Ranger", "Knight", "Mage"}
var armorLastNames_2 = []string{"Upcoming", "Kinda good", "Shows promise", "Getting there"}
var armorFirstNames_3 = []string{"Champion", "Legolas", "Charger", "BattleMage"}
var armorLastNames_3 = []string{"Next Level", "MFer", "Hardcore", "Killer"}

func getLevelBasedOnPlayer(level int) int {
	return utils.GetRandomNumberInRange(level-LEVEL_DIFFERENCE, level+LEVEL_DIFFERENCE)
}

func getXpGiven(level, playerLevel int) int {
	return 4*level + 3*playerLevel + 4
}

func getHP(level int) int {
	return 2*level + 15 + utils.GetRandomNumberInRange(level, 3*level)
}

func getRandomNameByLevel(level int) string {
	name := ""

	if level <= 10 {
		name = firstNames_1[utils.GetRandomNumberInRange(0, 3)]
		name += lastNames_1[utils.GetRandomNumberInRange(0, 3)]
	} else if level <= 20 {
		name = firstNames_1[utils.GetRandomNumberInRange(0, 3)]
		name += lastNames_2[utils.GetRandomNumberInRange(0, 3)]
	} else if level <= 30 {
		name = firstNames_1[utils.GetRandomNumberInRange(0, 3)]
		name += lastNames_3[utils.GetRandomNumberInRange(0, 3)]
	} else if level <= 40 {
		name = firstNames_2[utils.GetRandomNumberInRange(0, 3)]
		name += lastNames_1[utils.GetRandomNumberInRange(0, 3)]
	} else if level <= 50 {
		name = firstNames_2[utils.GetRandomNumberInRange(0, 3)]
		name += lastNames_2[utils.GetRandomNumberInRange(0, 3)]
	} else if level <= 60 {
		name = firstNames_2[utils.GetRandomNumberInRange(0, 3)]
		name += lastNames_3[utils.GetRandomNumberInRange(0, 3)]
	} else if level <= 70 {
		name = firstNames_3[utils.GetRandomNumberInRange(0, 3)]
		name += lastNames_1[utils.GetRandomNumberInRange(0, 3)]
	} else if level <= 80 {
		name = firstNames_3[utils.GetRandomNumberInRange(0, 3)]
		name += lastNames_2[utils.GetRandomNumberInRange(0, 3)]
	} else {
		name = firstNames_3[utils.GetRandomNumberInRange(0, 3)]
		name += lastNames_3[utils.GetRandomNumberInRange(0, 3)]
	}

	return name
}

func getWeaponByLevel(level int) item.Weapon {
	weapon := &item.WeaponImpl{}
	weaponItem := &item.ItemImpl{}

	weaponItem.Name = nameWeapon(level)
	weaponItem.Level = level
	// weaponItem.Type =

	return weapon
}

func nameWeapon(level int) string {
	name := ""

	if level <= 10 {
		name = weaponFirstNames_1[utils.GetRandomNumberInRange(0, 3)]
		name += weaponLastNames_1[utils.GetRandomNumberInRange(0, 3)]
	} else if level <= 20 {
		name = weaponFirstNames_1[utils.GetRandomNumberInRange(0, 3)]
		name += weaponLastNames_2[utils.GetRandomNumberInRange(0, 3)]
	} else if level <= 30 {
		name = weaponFirstNames_1[utils.GetRandomNumberInRange(0, 3)]
		name += weaponLastNames_3[utils.GetRandomNumberInRange(0, 3)]
	} else if level <= 40 {
		name = weaponFirstNames_2[utils.GetRandomNumberInRange(0, 3)]
		name += weaponLastNames_1[utils.GetRandomNumberInRange(0, 3)]
	} else if level <= 50 {
		name = weaponFirstNames_2[utils.GetRandomNumberInRange(0, 3)]
		name += weaponLastNames_2[utils.GetRandomNumberInRange(0, 3)]
	} else if level <= 60 {
		name = weaponFirstNames_2[utils.GetRandomNumberInRange(0, 3)]
		name += weaponLastNames_3[utils.GetRandomNumberInRange(0, 3)]
	} else if level <= 70 {
		name = weaponFirstNames_3[utils.GetRandomNumberInRange(0, 3)]
		name += weaponLastNames_1[utils.GetRandomNumberInRange(0, 3)]
	} else if level <= 80 {
		name = weaponFirstNames_3[utils.GetRandomNumberInRange(0, 3)]
		name += weaponLastNames_2[utils.GetRandomNumberInRange(0, 3)]
	} else {
		name = weaponFirstNames_3[utils.GetRandomNumberInRange(0, 3)]
		name += weaponLastNames_3[utils.GetRandomNumberInRange(0, 3)]
	}

	return name
}

func nameArmor(level int) string {
	name := ""

	if level <= 10 {
		name = armorFirstNames_1[utils.GetRandomNumberInRange(0, 3)]
		name += armorLastNames_1[utils.GetRandomNumberInRange(0, 3)]
	} else if level <= 20 {
		name = armorFirstNames_1[utils.GetRandomNumberInRange(0, 3)]
		name += armorLastNames_2[utils.GetRandomNumberInRange(0, 3)]
	} else if level <= 30 {
		name = armorFirstNames_1[utils.GetRandomNumberInRange(0, 3)]
		name += armorLastNames_3[utils.GetRandomNumberInRange(0, 3)]
	} else if level <= 40 {
		name = armorFirstNames_2[utils.GetRandomNumberInRange(0, 3)]
		name += armorLastNames_1[utils.GetRandomNumberInRange(0, 3)]
	} else if level <= 50 {
		name = armorFirstNames_2[utils.GetRandomNumberInRange(0, 3)]
		name += armorLastNames_2[utils.GetRandomNumberInRange(0, 3)]
	} else if level <= 60 {
		name = armorFirstNames_2[utils.GetRandomNumberInRange(0, 3)]
		name += armorLastNames_3[utils.GetRandomNumberInRange(0, 3)]
	} else if level <= 70 {
		name = armorFirstNames_3[utils.GetRandomNumberInRange(0, 3)]
		name += armorLastNames_1[utils.GetRandomNumberInRange(0, 3)]
	} else if level <= 80 {
		name = armorFirstNames_3[utils.GetRandomNumberInRange(0, 3)]
		name += armorLastNames_2[utils.GetRandomNumberInRange(0, 3)]
	} else {
		name = armorFirstNames_3[utils.GetRandomNumberInRange(0, 3)]
		name += armorLastNames_3[utils.GetRandomNumberInRange(0, 3)]
	}

	return name
}
