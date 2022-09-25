package item

import "testing"

func TestArmorString(t *testing.T) {
	t.Run("Check armor string is not empty", func(t *testing.T) {
		armor := &ArmorImpl{
			Item: &ItemImpl{
				Name: "Testing Armor",
			},
		}

		if armor.String() == "" {
			t.Error("Failed to print armor specs")
		}
	})
}
