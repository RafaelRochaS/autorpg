package item

import "testing"

func TestItemString(t *testing.T) {
	t.Run("Check item type - weapon - string is not empty", func(t *testing.T) {
		itemType := WEAPON

		if itemType.String() == "" {
			t.Error("Failed to print item specs")
		}
	})

	t.Run("Check item type - armor - string is not empty", func(t *testing.T) {
		itemType := ARMOR

		if itemType.String() == "" {
			t.Error("Failed to print item specs")
		}
	})
}
