package enemy

import (
	"autorpg/utils"
)

/* Naming Scheme:
Level 1-10: first name level 1 last name level 1
Level 10-20: first name level 2 last name level 1
Level 20-30: first name level 3 last name level 1
Level 30-40: first name level 1 last name level 2 ...
*/

const LEVEL_DIFFERENCE = 3

var firstNames_1 = []string{"Foot Soldier", "Foot Archer", "Foot Horsemen", "Foot Wizard"}
var lastNames_1 = []string{"Weak boy", "Trash tier", "No skill", "N00b"}
var firstNames_2 = []string{"Sargeant", "Ranger", "Knight", "Mage"}
var lastNames_2 = []string{"Upcoming", "Kinda good", "Shows promise", "Getting there"}
var firstNames_3 = []string{"Champion", "Legolas", "Charger", "BattleMage"}
var lastNames_3 = []string{"Next Level", "MFer", "Hardcore", "Killer"}

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
