package character

import (
	"autorpg/utils"
	"strings"
	"testing"
)

const BARBARIAN_TESTING_STRING = "TestingBarbarian\n" +
	"4\n" +
	"10\n" +
	"0\n" +
	"0\n" +
	"0\n" +
	"0\n"

func TestCharacterCreation(t *testing.T) {
	utils.DEBUG = "False"

	t.Run("Create Barbarian", func(t *testing.T) {
		barbarian := NewCharacter(strings.NewReader(BARBARIAN_TESTING_STRING))
		barbarian.Create()

		if barbarian.GetPerson().Name != "TestingBarbarian" {
			t.Errorf("Create Barbarian Error - Name :: got %s, wanted %s", barbarian.GetPerson().Name, "TestingBarbarian")
		}

		if barbarian.GetClass() != BARBARIAN {
			t.Errorf("Create Barbarian Error - Class :: got %s, wanted %s", barbarian.GetClass(), BARBARIAN)
		}
	})
}
