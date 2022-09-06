package character

import (
	"autorpg/item"
	"autorpg/utils"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

const charTestName = "Testing Character Name"
const weaponTestName = "Testing Weapon Name"
const armorTestName = "Testing Armor Name"

func TestCharacterCreation(t *testing.T) {
	utils.DEBUG = "True"

	char := createTestChar()

	t.Run("Create debug character", func(t *testing.T) {
		char.Create()
	})

	t.Run("Set character name", func(t *testing.T) {
		char.SetName(charTestName)
		if char.GetPerson().Name != charTestName {
			t.Errorf("Character Creation Error - Name :: got %s, wanted %s",
				char.GetPerson().Name, charTestName)
		}
	})

	t.Run("Check class is defined", func(t *testing.T) {
		if char.GetClass().String() == "" {
			t.Errorf("Character Creation Error - Class :: got %s, wanted %s",
				"", "any class set")
		}
	})

	t.Run("Check stats are defined", func(t *testing.T) {
		if char.GetHP() <= 0 {
			t.Errorf("Character Creation Error - Stats :: got %d, wanted more than 0",
				char.GetHP())
		}
	})
}

func TestCharacterItems(t *testing.T) {
	char := createTestChar()
	weapon := createTestWeapon()
	armor := createTestArmor()

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
}

func createTestChar() Character {
	return NewCharacter(os.Stdin)
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
