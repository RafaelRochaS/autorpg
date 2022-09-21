package character

import (
	"autorpg/item"
	"autorpg/utils"
	"errors"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

const charTestName = "Testing Character Name"
const weaponTestName = "Testing Weapon Name"
const armorTestName = "Testing Armor Name"
const itemTestName = "Testing Item Name"
const testingDamage = 3

func TestCharacterBasics(t *testing.T) {
	utils.DEBUG = "True"

	char := createTestChar()

	t.Run("Create debug character", func(t *testing.T) {
		char.Create()
	})

	t.Run("Set character name", func(t *testing.T) {
		char.SetName(charTestName)
		if char.GetPerson().Name != charTestName {
			t.Errorf("Character Basics Error - Name :: got %s, wanted %s",
				char.GetPerson().Name, charTestName)
		}
	})

	t.Run("Check class is defined", func(t *testing.T) {
		if char.GetClass().String() == "" {
			t.Errorf("Character Basics Error - Class :: got %s, wanted %s",
				"", "any class set")
		}
	})

	t.Run("Check stats are defined", func(t *testing.T) {
		if char.GetHP() <= 0 {
			t.Errorf("Character Basics Error - Stats :: got %d, wanted more than 0",
				char.GetHP())
		}
	})

	t.Run("Reader failure", func(t *testing.T) {
		mockChar := createTestCharBroken()
		utils.DEBUG = "False"
		assert.Panics(t,
			func() { mockChar.Create() },
			"Character Basics Error - Mocked Reader :: did not panic")
	})

	t.Run("Check damage is done", func(t *testing.T) {
		hpBeforeDmg := char.GetHP()
		char.TakeDamage(testingDamage)
		if char.GetHP() != hpBeforeDmg-testingDamage {
			t.Errorf("Character Basics Error - Take Damage :: got %d, wanted %d",
				char.GetHP(), hpBeforeDmg-testingDamage)
		}
	})

	t.Run("Check hp is reset", func(t *testing.T) {
		hpAfterDmg := char.GetHP()
		char.ResetHP()
		if char.GetHP() != hpAfterDmg+testingDamage {
			t.Errorf("Character Basics Error - Reset HP :: got %d, wanted %d",
				char.GetHP(), hpAfterDmg-testingDamage)
		}
	})

	t.Run("Check Get Luck", func(t *testing.T) {
		if char.GetLuck() <= 0 {
			t.Errorf("Character Basics Error - Get Luck:: got %d, wanted more than 0",
				char.GetLuck())
		}
	})
}

func TestCharacterItems(t *testing.T) {
	utils.DEBUG = "True"

	char := createTestChar()
	char.Create()
	weapon := createTestWeapon()
	armor := createTestArmor()
	defaultWeapon := char.GetWeapon()
	defaultArmor := char.GetArmor()

	setDefaultEquipment := func() {
		char.AttachWeapon(defaultWeapon)
		char.AttachArmor(defaultArmor)
	}

	t.Run("Set test weapon", func(t *testing.T) {
		char.AttachWeapon(weapon)
		if char.GetWeapon().GetName() != weaponTestName {
			t.Errorf("Character Items Error - Weapon :: got %s, wanted %s",
				char.GetWeapon().GetName(), weaponTestName)
		}
	})

	t.Run("Set test armor", func(t *testing.T) {
		char.AttachArmor(armor)
		if char.GetArmor().GetName() != armorTestName {
			t.Errorf("Character Items Error - Armor :: got %s, wanted %s",
				char.GetWeapon().GetName(), armorTestName)
		}
	})

	t.Run("Remove weapon", func(t *testing.T) {
		char.RemoveItem(item.WEAPON)
		assert.Panics(t,
			func() { char.GetWeapon().GetName() },
			"Character Items Error - Remove Weapon :: item still exists")
	})

	t.Run("Remove armor", func(t *testing.T) {
		char.RemoveItem(item.ARMOR)
		assert.Panics(t,
			func() { char.GetArmor().GetName() },
			"Character Items Error - Remove Armor :: item still exists")
	})

	t.Run("Check item level - higher", func(t *testing.T) {
		testItem := createTestItem(999, 1, 1, 1)
		if char.CheckLevelItem(testItem) {
			t.Error("Character Items Error - Level Higher :: got 'true', wanted 'false'")
		}
	})

	t.Run("Check item level - same/lower", func(t *testing.T) {
		testItem := createTestItem(1, 1, 1, 1)
		if !char.CheckLevelItem(testItem) {
			t.Error("Character Items Error - Level same/lower :: got 'false', wanted 'true'")
		}
	})

	t.Run("Check stat requirement - higher", func(t *testing.T) {
		testItem := createTestItem(1, 999, 999, 999)
		if char.CheckStatRequirementsItem(testItem) {
			t.Error("Character Items Error - Level Higher :: got 'true', wanted 'false'")
		}
	})

	t.Run("Check stat requirement - lower", func(t *testing.T) {
		testItem := createTestItem(1, 1, 1, 1)
		if !char.CheckStatRequirementsItem(testItem) {
			t.Error("Character Items Error - Level Higher :: got 'false', wanted 'true'")
		}
	})

	t.Run("Check Weapon Drop - useable/bad", func(t *testing.T) {
		setDefaultEquipment()
		testWeaponItem := createTestItem(1, 1, 1, 1)
		testWeapon := &item.WeaponImpl{
			Item:        testWeaponItem,
			DamageType:  item.NORMAL,
			Damage:      1,
			AttackSpeed: 1,
		}

		char.HandleWeaponDrop(testWeapon)
		if char.GetWeapon().GetName() == itemTestName {
			t.Error("Character Items Error - Handle weapon drop - useable/bad :: was not supposed to equip")
		}
	})

	t.Run("Check Weapon Drop - useable/good", func(t *testing.T) {
		setDefaultEquipment()
		testWeaponItem := createTestItem(1, 1, 1, 1)
		testWeapon := &item.WeaponImpl{
			Item:        testWeaponItem,
			DamageType:  item.NORMAL,
			Damage:      100,
			AttackSpeed: 100,
		}

		char.HandleWeaponDrop(testWeapon)
		if char.GetWeapon().GetName() != itemTestName {
			t.Error("Character Items Error - Handle weapon drop - useable/good :: was supposed to equip")
		}
	})

	t.Run("Check Weapon Drop - useable/good - low attackspeed", func(t *testing.T) {
		setDefaultEquipment()
		testWeaponItem := createTestItem(1, 1, 1, 1)
		testWeapon := &item.WeaponImpl{
			Item:        testWeaponItem,
			DamageType:  item.NORMAL,
			Damage:      999,
			AttackSpeed: 0.01,
		}

		char.HandleWeaponDrop(testWeapon)
		if char.GetWeapon().GetName() != itemTestName {
			t.Error("Character Items Error - Handle weapon drop - useable/good :: was supposed to equip")
		}
	})

	t.Run("Check Weapon Drop - unuseable", func(t *testing.T) {
		setDefaultEquipment()
		testWeaponItem := createTestItem(1, 999, 999, 999)
		testWeapon := &item.WeaponImpl{
			Item:        testWeaponItem,
			DamageType:  item.NORMAL,
			Damage:      100,
			AttackSpeed: 100,
		}

		char.HandleWeaponDrop(testWeapon)
		if char.GetWeapon().GetName() == itemTestName {
			t.Error("Character Items Error - Handle weapon drop - unuseable :: was not supposed to equip")
		}
	})

	t.Run("Check Armor Drop - unuseable", func(t *testing.T) {
		setDefaultEquipment()
		testArmorItem := createTestItem(1, 999, 999, 999)
		testArmor := &item.ArmorImpl{
			Item:    testArmorItem,
			Defense: 1,
			Weight:  1,
		}

		char.HandleArmorDrop(testArmor)
		if char.GetArmor().GetName() == itemTestName {
			t.Error("Character Items Error - Handle armor drop - unuseable :: was not supposed to equip")
		}
	})

	t.Run("Check Armor Drop - useable/bad", func(t *testing.T) {
		setDefaultEquipment()
		testArmorItem := createTestItem(1, 1, 1, 1)
		testArmor := &item.ArmorImpl{
			Item:    testArmorItem,
			Defense: 1,
			Weight:  1,
		}

		char.HandleArmorDrop(testArmor)
		if char.GetArmor().GetName() == itemTestName {
			t.Error("Character Items Error - Handle armor drop - useable/bad :: was not supposed to equip")
		}
	})

	t.Run("Check Armor Drop - useable/good", func(t *testing.T) {
		setDefaultEquipment()
		testArmorItem := createTestItem(1, 1, 1, 1)
		testArmor := &item.ArmorImpl{
			Item:    testArmorItem,
			Defense: 100,
			Weight:  1,
		}

		char.HandleArmorDrop(testArmor)
		if char.GetArmor().GetName() != itemTestName {
			t.Error("Character Items Error - Handle armor drop - useable/good :: was supposed to equip")
		}
	})

	t.Run("Check Armor Drop - useable/good - heavy but good defense", func(t *testing.T) {
		setDefaultEquipment()
		testArmorItem := createTestItem(1, 1, 1, 1)
		testArmor := &item.ArmorImpl{
			Item:    testArmorItem,
			Defense: 1000,
			Weight:  100,
		}

		char.HandleArmorDrop(testArmor)
		if char.GetArmor().GetName() != itemTestName {
			t.Error("Character Items Error - Handle armor drop - useable/good - heavy but good defense :: was supposed to equip")
		}
	})
}

func TestCharacterLevel(t *testing.T) {
	utils.DEBUG = "True"
	char := createTestChar()
	char.Create()
	currentLevel := char.GetPerson().Level

	t.Run("Attempt level - no xp", func(t *testing.T) {
		char.LevelUp()
		if currentLevel < char.GetPerson().Level {
			t.Error("Character Level Error :: Leveled up with no xp")
		}
	})

	t.Run("Give xp", func(t *testing.T) {
		currentXp := char.GetCurrentXp()
		char.IncreaseXP(1000)
		if currentXp >= char.GetCurrentXp() {
			t.Error("Character Level Error :: xp increase did not change xp")
		}
	})

	t.Run("Attempt level - given xp", func(t *testing.T) {
		char.LevelUp()
		if currentLevel >= char.GetPerson().Level {
			t.Error("Character Level Error :: Given xp and not leveled up")
		}
	})
}

func TestStdin(t *testing.T) {
	utils.DEBUG = "False"
	const namingTest = "Test Name\n"
	const namingTestError = "Test Name"

	t.Run("Test Handle Naming", func(t *testing.T) {
		char := &CharacterImpl{
			io: strings.NewReader(namingTest),
		}
		err := char.HandleNaming()
		if err != nil {
			t.Errorf("Character Stdin Error - Naming :: error where it shouldnt")
		}

		if char.GetPerson().Name != strings.Replace(namingTest, "\n", "", -1) {
			t.Errorf("Character Stdin Error - Naming :: got %s, wanted %s",
				char.GetPerson().Name, namingTest)
		}
	})

	t.Run("Test Handle Naming - Error", func(t *testing.T) {
		char := &CharacterImpl{
			io: strings.NewReader(namingTestError),
		}
		err := char.HandleNaming()

		if err == nil {
			t.Errorf("Character Stdin Error - Naming with Error :: failed input did not error")
		}
	})

	t.Run("Set class - Warrior", func(t *testing.T) {
		char := &CharacterImpl{
			io: strings.NewReader("1\n"),
		}
		err := char.HandleClass()

		if err != nil {
			t.Errorf("Character Stdin Error - Class (Warrior) :: got error: %s", err.Error())
		}

		if char.GetClass() != WARRIOR {
			t.Errorf("Character Stdin Error - Class (Warrior) :: got wrong class: %s", char.GetClass())
		}

		if char.GetHP() != WAR_HP {
			t.Errorf("Character Stdin Error - Class (Warrior) :: got wrong HP: %d", char.GetHP())
		}
	})

	t.Run("Set class - Rogue", func(t *testing.T) {
		char := &CharacterImpl{
			io: strings.NewReader("2\n"),
		}
		err := char.HandleClass()

		if err != nil {
			t.Errorf("Character Stdin Error - Class (Rogue) :: got error: %s", err.Error())
		}

		if char.GetClass() != ROGUE {
			t.Errorf("Character Stdin Error - Class (Rogue) :: got wrong class: %s", char.GetClass())
		}

		if char.GetHP() != ROG_HP {
			t.Errorf("Character Stdin Error - Class (Rogue) :: got wrong HP: %d", char.GetHP())
		}
	})

	t.Run("Set class - Wizard", func(t *testing.T) {
		char := &CharacterImpl{
			io: strings.NewReader("3\n"),
		}
		err := char.HandleClass()

		if err != nil {
			t.Errorf("Character Stdin Error - Class (Wizard) :: got error: %s", err.Error())
		}

		if char.GetClass() != WIZARD {
			t.Errorf("Character Stdin Error - Class (Wizard) :: got wrong class: %s", char.GetClass())
		}

		if char.GetHP() != WIZ_HP {
			t.Errorf("Character Stdin Error - Class (Wizard) :: got wrong HP: %d", char.GetHP())
		}
	})

	t.Run("Set class - Barbarian", func(t *testing.T) {
		char := &CharacterImpl{
			io: strings.NewReader("4\n"),
		}
		err := char.HandleClass()

		if err != nil {
			t.Errorf("Character Stdin Error - Class (Barbarian) :: got error: %s", err.Error())
		}

		if char.GetClass() != BARBARIAN {
			t.Errorf("Character Stdin Error - Class (Barbarian) :: got wrong class: %s", char.GetClass())
		}

		if char.GetHP() != BAR_HP {
			t.Errorf("Character Stdin Error - Class (Barbarian) :: got wrong HP: %d", char.GetHP())
		}
	})

	t.Run("Set class - Failed input", func(t *testing.T) {
		char := &CharacterImpl{
			io: strings.NewReader("banana\n"),
		}
		err := char.HandleClass()

		if err == nil {
			t.Errorf("Character Stdin Error - Class - Failed input :: did not get error")
		}
	})
}

func createTestChar() Character {
	return NewCharacter(os.Stdin)
}

func createTestCharBroken() Character {
	return NewCharacter(&ReaderMock{})
}

func createTestWeapon() item.Weapon {
	return &item.WeaponImpl{
		Item: &item.ItemImpl{
			Name: weaponTestName,
		},
	}
}

func createTestArmor() item.Armor {
	return &item.ArmorImpl{
		Item: &item.ItemImpl{
			Name: armorTestName,
		},
	}
}

func createTestItem(level, strReq, dexReq, intReq int) item.Item {
	return &item.ItemImpl{
		Name:   itemTestName,
		Level:  level,
		StrReq: strReq,
		DexReq: dexReq,
		IntReq: intReq,
	}
}

type ReaderMock struct{}

func (r ReaderMock) Read(p []byte) (int, error) {
	return -1, errors.New("mocking error")
}
