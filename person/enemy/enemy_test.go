package enemy

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnemyCreation(t *testing.T) {
	t.Run("New enemy creates correctly", func(t *testing.T) {
		testPlayerLevel := 3
		enemy := NewEnemy(testPlayerLevel)

		t.Run("Name is created", func(t *testing.T) {
			if enemy.GetName() == "" {
				t.Error("TestEnemyCreation - Name :: name came empty")
			}
		})

		t.Run("Level is appropriate", func(t *testing.T) {
			if enemy.GetLevel() < testPlayerLevel-LEVEL_DIFFERENCE ||
				enemy.GetLevel() > testPlayerLevel+LEVEL_DIFFERENCE {
				t.Errorf(
					"TestEnemyCreation - Level :: level was set outside of range, got: %d - player level was: %d",
					enemy.GetLevel(),
					testPlayerLevel)
			}
		})

		t.Run("Has Weapon", func(t *testing.T) {
			assert.NotPanics(t, func() { enemy.GetWeapon() })
		})

		t.Run("Has Armor", func(t *testing.T) {
			assert.NotPanics(t, func() { enemy.GetArmor() })
		})

		t.Run("Has XpGiven", func(t *testing.T) {
			assert.NotPanics(t, func() { enemy.GetXpGiven() })
		})

		t.Run("Has Hp", func(t *testing.T) {
			assert.NotPanics(t, func() { enemy.GetHP() })
		})

		t.Run("Has person", func(t *testing.T) {
			assert.NotPanics(t, func() { enemy.GetPerson() })
		})

		t.Run("String renders correctly", func(t *testing.T) {
			assert.NotPanics(t, func() { fmt.Print(enemy.String()) })
		})
	})
}

func TestEnemyFunctionality(t *testing.T) {
	t.Run("Different Levels", func(t *testing.T) {
		playerLevels := []int{3, 10, 20, 30, 40, 50, 60, 70, 80, 90, 100}

		for _, level := range playerLevels {

			enemy := NewEnemy(level)
			weapon := enemy.GetDropWeapon(10)
			armor := enemy.GetDropArmor(10)

			if weapon.GetLevel() != level {
				t.Errorf("TestEnemyFunctionality - Drops :: Weapon got wrong level: got '%d', wanted '%d'",
					weapon.GetLevel(), level)
			}

			if armor.GetLevel() != level {
				t.Errorf("TestEnemyFunctionality - Drops :: Armor got wrong level: got '%d', wanted '%d'",
					armor.GetLevel(), level)
			}
		}
	})

	t.Run("Take Damage", func(t *testing.T) {
		damageToTake := 3
		enemy := NewEnemy(5)
		currentHp := enemy.GetHP()
		enemy.TakeDamage(damageToTake)

		if enemy.GetHP() != currentHp-damageToTake {
			t.Errorf("TestEnemyFunctionality - Take Damage :: got '%d', wanted '%d'",
				enemy.GetHP(), currentHp-damageToTake)
		}
	})
}
