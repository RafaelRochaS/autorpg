package combat

import (
	"autorpg/item"
	"autorpg/person/character"
	"autorpg/person/enemy"
	"autorpg/utils"
	"os"
	"testing"

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
		weakEnemy := enemy.NewEnemy(1)

		assert.NotPanics(t, func() { NewCombat().InitiateCombat(strongPlayer, weakEnemy) })
	})
}
