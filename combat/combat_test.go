package combat

import (
	"os"
	"testing"

	"github.com/RafaelRochaS/autorpg/item"
	"github.com/RafaelRochaS/autorpg/person/character"
	"github.com/RafaelRochaS/autorpg/person/enemy"
	"github.com/RafaelRochaS/autorpg/utils"

	"github.com/stretchr/testify/assert"
)

func TestCombat(t *testing.T) {
	t.Run("Test execution - player wins", func(t *testing.T) {
		utils.DEBUG = "True"

		strongPlayer := character.NewCharacter(os.Stdin)
		strongPlayer.AttachWeapon(&item.WeaponImpl{
			DamageType:  item.NORMAL,
			Damage:      100_000,
			AttackSpeed: 0.1,
		})
		strongPlayer.AttachArmor(&item.ArmorImpl{
			Weight:  1,
			Defense: 10,
		})
		currentXp := strongPlayer.GetCurrentXp()

		weakEnemy := enemy.NewEnemy(1)
		givenXp := weakEnemy.GetXpGiven()

		assert.NotPanics(t, func() { NewCombat().InitiateCombat(strongPlayer, weakEnemy) })
		if strongPlayer.GetCurrentXp() != (currentXp + givenXp) {
			t.Errorf("Player Wins :: xp error, wanted %d, got %d", (currentXp + givenXp), strongPlayer.GetCurrentXp())
		}
	})

	t.Run("Test execution - enemy wins", func(t *testing.T) {
		utils.DEBUG = "True"

		weakPlayer := character.NewCharacter(os.Stdin)
		weakPlayer.Create()

		strongEnemy := enemy.NewEnemy(300)

		assert.NotPanics(t, func() { NewCombat().InitiateCombat(weakPlayer, strongEnemy) })
	})
}
