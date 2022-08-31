package person

import (
	"autorpg/item"
	"testing"
)

func setupPersonTesting() Person {
	return Person{
		Name:  "Test",
		Level: 1,
	}
}

func TestBasicPersonFunctionality(t *testing.T) {
	testPerson := setupPersonTesting()

	t.Run("Add basic armor", func(t *testing.T) {
		const testingArmorName = "Testing Armor"
		const testingArmorWeight = 1
		const testingArmorDefense = 1

		testPerson.SetArmor(&item.ArmorImpl{
			Item: item.ItemImpl{
				Name: testingArmorName,
			},
			Defense: testingArmorDefense,
			Weight:  testingArmorWeight,
		})

		if testPerson.Armor.GetName() != testingArmorName {
			t.Errorf("Person Armor Error - Name :: got %s, wanted %s", testPerson.Armor.GetName(), testingArmorName)
		}

		if testPerson.Armor.GetDefense() != testingArmorDefense {
			t.Errorf("Person Armor Error - Defense :: got %d, wanted %d", testPerson.Armor.GetDefense(), testingArmorDefense)
		}

		if testPerson.Armor.GetWeight() != testingArmorWeight {
			t.Errorf("Person Armor Error - Weight :: got %d, wanted %d", testPerson.Armor.GetWeight(), testingArmorWeight)
		}
	})

	t.Run("Add basic weapon", func(t *testing.T) {
		const testingWeaponName = "Testing Weapon"
		const testingWeaponDamageType = item.NORMAL
		const testingWeaponDamage = 1
		const testingWeaponAttackSpeed = 1.0

		testPerson.SetWeapon(&item.WeaponImpl{
			Item: item.ItemImpl{
				Name: testingWeaponName,
			},
			DamageType:  testingWeaponDamageType,
			Damage:      testingWeaponDamage,
			AttackSpeed: testingWeaponAttackSpeed,
		})

		if testPerson.Weapon.GetName() != testingWeaponName {
			t.Errorf("Person Weapon Error - Name :: got %s, wanted %s", testPerson.Weapon.GetName(), testingWeaponName)
		}

		if testPerson.Weapon.GetDamageType() != testingWeaponDamageType {
			t.Errorf("Person Weapon Error - DamageType :: got %d, wanted %d", testPerson.Weapon.GetDamageType(), testingWeaponDamageType)
		}

		if testPerson.Weapon.GetDamage() != testingWeaponDamage {
			t.Errorf("Person Weapon Error - Damage :: got %d, wanted %d", testPerson.Weapon.GetDamage(), testingWeaponDamage)
		}

		if testPerson.Weapon.GetAttackSpeed() != testingWeaponAttackSpeed {
			t.Errorf("Person Weapon Error - AttackSpeed :: got %f, wanted %f", testPerson.Weapon.GetAttackSpeed(), testingWeaponAttackSpeed)
		}
	})
}
