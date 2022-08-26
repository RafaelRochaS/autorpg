package enemy

import (
	"autorpg/item"
	"autorpg/utils"
	"strings"
)

const LEVEL_DIFFERENCE = 2
const LEVEL_DIFFERENCE_WEAPON = 10
const LEVEL_DIFFERENCE_ARMOR_DEFENSE = 15
const LEVEL_DIFFERENCE_ARMOR_WEIGHT = 2

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
var weaponFirstNames_1 = []string{"Basic Sword", "Basic Spear", "Basic Bow", "Basic Staff"}
var weaponLastNames_1 = []string{"of the Little Bitch", "that kinda sucks", "that is dull", "that needs repair"}
var weaponFirstNames_2 = []string{"Good Sword", "Good Spear", "Good Bow", "Good Staff"}
var weaponLastNames_2 = []string{"that could be better", "that does the job", "that kinda works", "good enough"}
var weaponFirstNames_3 = []string{"Epic Sword", "Devastating Spear", "Mythical Bow", "Staff do Caralho"}
var weaponLastNames_3 = []string{"that vaporizes", "that kills'em all", "that blows, in a good way", "that destroys"}

/** ARMOR **/
var armorFirstNames_1 = []string{"Basic Mail", "Basic Leather", "Basic Steel", "Basic Robe"}
var armorLastNames_1 = []string{"kinda broken", "with some holes", "that shows skin", "for weaklings"}
var armorFirstNames_2 = []string{"Good Mail", "Good Leather", "Good Steel", "Good Robe"}
var armorLastNames_2 = []string{"that could be better", "that does the job", "that kinda works", "good enough"}
var armorFirstNames_3 = []string{"Fantastic Mail", "Special Leather", "Fuckin Great Steel", "Nice Robe"}
var armorLastNames_3 = []string{"top na balada", "that is a fucking wall", ", unbreakable", "that nothing gets through"}

func getLevelBasedOnPlayer(level int) int {
	return utils.GetRandomNumberInRange(level-LEVEL_DIFFERENCE, level+LEVEL_DIFFERENCE)
}

func getXpGiven(level, playerLevel int) int {
	return 4*level + 3*playerLevel + 4
}

func getHP(level int) int {
	return 3*level + 15 + utils.GetRandomNumberInRange(level, 3*level)
}

func getRandomNameByLevel(level int) string {
	name := ""

	if level <= 10 {
		name = firstNames_1[utils.GetRandomNumberInRange(0, 3)]
		name += " "
		name += lastNames_1[utils.GetRandomNumberInRange(0, 3)]
	} else if level <= 20 {
		name = firstNames_1[utils.GetRandomNumberInRange(0, 3)]
		name += " "
		name += lastNames_2[utils.GetRandomNumberInRange(0, 3)]
	} else if level <= 30 {
		name = firstNames_1[utils.GetRandomNumberInRange(0, 3)]
		name += " "
		name += lastNames_3[utils.GetRandomNumberInRange(0, 3)]
	} else if level <= 40 {
		name = firstNames_2[utils.GetRandomNumberInRange(0, 3)]
		name += " "
		name += lastNames_1[utils.GetRandomNumberInRange(0, 3)]
	} else if level <= 50 {
		name = firstNames_2[utils.GetRandomNumberInRange(0, 3)]
		name += " "
		name += lastNames_2[utils.GetRandomNumberInRange(0, 3)]
	} else if level <= 60 {
		name = firstNames_2[utils.GetRandomNumberInRange(0, 3)]
		name += " "
		name += lastNames_3[utils.GetRandomNumberInRange(0, 3)]
	} else if level <= 70 {
		name = firstNames_3[utils.GetRandomNumberInRange(0, 3)]
		name += " "
		name += lastNames_1[utils.GetRandomNumberInRange(0, 3)]
	} else if level <= 80 {
		name = firstNames_3[utils.GetRandomNumberInRange(0, 3)]
		name += " "
		name += lastNames_2[utils.GetRandomNumberInRange(0, 3)]
	} else {
		name = firstNames_3[utils.GetRandomNumberInRange(0, 3)]
		name += " "
		name += lastNames_3[utils.GetRandomNumberInRange(0, 3)]
	}

	return name
}

func getWeaponByLevel(level int, luck int) item.Weapon {
	weapon := &item.WeaponImpl{}
	weaponItem := &item.ItemImpl{}
	weaponItem.Name = nameWeapon(level)
	weaponItem.Level = level

	if strings.Contains(weaponItem.Name, "Staff") {
		weapon.DamageType = item.MAGICAL
		weapon.Damage, weapon.AttackSpeed = getWeaponDamageAndAttackSpeedMagical(level, luck)
		weaponItem.StrReq = 1 + utils.GetRandomNumberInRange(0, level)
		weaponItem.DexReq = 5 + utils.GetRandomNumberInRange(0, level)
		weaponItem.IntReq = level + 10 + utils.GetRandomNumberInRange(0, level)
	} else {
		weapon.DamageType = item.NORMAL
		weapon.Damage, weapon.AttackSpeed = getWeaponDamageAndAttackSpeed(level, luck)
		weaponItem.StrReq = level + 10 + utils.GetRandomNumberInRange(0, level)
		weaponItem.DexReq = 4 + utils.GetRandomNumberInRange(0, level)
		weaponItem.IntReq = 1 + utils.GetRandomNumberInRange(0, level)
	}
	weaponItem.Type = item.WEAPON

	weapon.Item = weaponItem

	return weapon
}

func nameWeapon(level int) string {
	name := ""

	if level <= 10 {
		name = weaponFirstNames_1[utils.GetRandomNumberInRange(0, 3)]
		name += " "
		name += weaponLastNames_1[utils.GetRandomNumberInRange(0, 3)]
	} else if level <= 20 {
		name = weaponFirstNames_1[utils.GetRandomNumberInRange(0, 3)]
		name += " "
		name += weaponLastNames_2[utils.GetRandomNumberInRange(0, 3)]
	} else if level <= 30 {
		name = weaponFirstNames_1[utils.GetRandomNumberInRange(0, 3)]
		name += " "
		name += weaponLastNames_3[utils.GetRandomNumberInRange(0, 3)]
	} else if level <= 40 {
		name = weaponFirstNames_2[utils.GetRandomNumberInRange(0, 3)]
		name += " "
		name += weaponLastNames_1[utils.GetRandomNumberInRange(0, 3)]
	} else if level <= 50 {
		name = weaponFirstNames_2[utils.GetRandomNumberInRange(0, 3)]
		name += " "
		name += weaponLastNames_2[utils.GetRandomNumberInRange(0, 3)]
	} else if level <= 60 {
		name = weaponFirstNames_2[utils.GetRandomNumberInRange(0, 3)]
		name += " "
		name += weaponLastNames_3[utils.GetRandomNumberInRange(0, 3)]
	} else if level <= 70 {
		name = weaponFirstNames_3[utils.GetRandomNumberInRange(0, 3)]
		name += " "
		name += weaponLastNames_1[utils.GetRandomNumberInRange(0, 3)]
	} else if level <= 80 {
		name = weaponFirstNames_3[utils.GetRandomNumberInRange(0, 3)]
		name += " "
		name += weaponLastNames_2[utils.GetRandomNumberInRange(0, 3)]
	} else {
		name = weaponFirstNames_3[utils.GetRandomNumberInRange(0, 3)]
		name += " "
		name += weaponLastNames_3[utils.GetRandomNumberInRange(0, 3)]
	}

	return name
}

func getWeaponDamageAndAttackSpeed(level int, luck int) (int, float32) {
	damage := utils.GetRandomNumberInRange(level-LEVEL_DIFFERENCE_WEAPON, level+LEVEL_DIFFERENCE_WEAPON+luck)
	var attackSpeed float32

	if damage > level {
		attackSpeed = utils.GetRandomFloatInRange(0.5, 1.5)
	} else {
		attackSpeed = utils.GetRandomFloatInRange(1.5, 3)
	}

	return damage, attackSpeed
}

func getWeaponDamageAndAttackSpeedMagical(level int, luck int) (int, float32) {
	damage := utils.GetRandomNumberInRange(level-LEVEL_DIFFERENCE_WEAPON, level+LEVEL_DIFFERENCE_WEAPON+luck)
	var attackSpeed float32

	if damage > level {
		attackSpeed = utils.GetRandomFloatInRange(0.4, 0.7)
	} else {
		attackSpeed = utils.GetRandomFloatInRange(0.8, 1.1)
	}

	return damage, attackSpeed
}

func getArmorByLevel(level int, luck int) item.Armor {
	armor := item.ArmorImpl{}

	armor.Item = makeArmorItem(level)
	armor.Defense = utils.GetRandomNumberInRange(level-LEVEL_DIFFERENCE_ARMOR_DEFENSE, level+LEVEL_DIFFERENCE_ARMOR_DEFENSE+luck)
	armor.Weight = utils.GetRandomNumberInRange(level-LEVEL_DIFFERENCE_ARMOR_WEIGHT, level+LEVEL_DIFFERENCE_ARMOR_WEIGHT+luck)

	return armor
}

func makeArmorItem(level int) item.Item {
	armorItem := item.ItemImpl{}

	armorItem.Name = nameArmor(level)
	armorItem.Level = level
	armorItem.Type = item.ARMOR
	armorItem.StrReq = level + utils.GetRandomNumberInRange(0, level)
	armorItem.DexReq = level + utils.GetRandomNumberInRange(0, level)
	armorItem.IntReq = level + utils.GetRandomNumberInRange(0, level)

	return armorItem
}

func nameArmor(level int) string {
	name := ""

	if level <= 10 {
		name = armorFirstNames_1[utils.GetRandomNumberInRange(0, 3)]
		name += " "
		name += armorLastNames_1[utils.GetRandomNumberInRange(0, 3)]
	} else if level <= 20 {
		name = armorFirstNames_1[utils.GetRandomNumberInRange(0, 3)]
		name += " "
		name += armorLastNames_2[utils.GetRandomNumberInRange(0, 3)]
	} else if level <= 30 {
		name = armorFirstNames_1[utils.GetRandomNumberInRange(0, 3)]
		name += " "
		name += armorLastNames_3[utils.GetRandomNumberInRange(0, 3)]
	} else if level <= 40 {
		name = armorFirstNames_2[utils.GetRandomNumberInRange(0, 3)]
		name += " "
		name += armorLastNames_1[utils.GetRandomNumberInRange(0, 3)]
	} else if level <= 50 {
		name = armorFirstNames_2[utils.GetRandomNumberInRange(0, 3)]
		name += " "
		name += armorLastNames_2[utils.GetRandomNumberInRange(0, 3)]
	} else if level <= 60 {
		name = armorFirstNames_2[utils.GetRandomNumberInRange(0, 3)]
		name += " "
		name += armorLastNames_3[utils.GetRandomNumberInRange(0, 3)]
	} else if level <= 70 {
		name = armorFirstNames_3[utils.GetRandomNumberInRange(0, 3)]
		name += " "
		name += armorLastNames_1[utils.GetRandomNumberInRange(0, 3)]
	} else if level <= 80 {
		name = armorFirstNames_3[utils.GetRandomNumberInRange(0, 3)]
		name += " "
		name += armorLastNames_2[utils.GetRandomNumberInRange(0, 3)]
	} else {
		name = armorFirstNames_3[utils.GetRandomNumberInRange(0, 3)]
		name += " "
		name += armorLastNames_3[utils.GetRandomNumberInRange(0, 3)]
	}

	return name
}
