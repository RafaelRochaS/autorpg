package item

import "testing"

func TestWeaponString(t *testing.T) {
	t.Run("Check weapon string is not empty", func(t *testing.T) {
		weapon := &WeaponImpl{
			Item: &ItemImpl{
				Name: "Testing Weapon",
			},
		}

		if weapon.String() == "" {
			t.Error("Failed to print weapon specs")
		}
	})
}
