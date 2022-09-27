package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrints(t *testing.T) {
	t.Run("Test functions that just print, for some reason...", func(t *testing.T) {
		assert.NotPanics(t, func() { PrintIntro() })
	})

	t.Run("Test functions that just print, for some reason...", func(t *testing.T) {
		assert.NotPanics(t, func() { AwaitInput() })
	})
}

func TestRandoms(t *testing.T) {
	t.Run("Assert random number returns a number between the range", func(t *testing.T) {
		randNumber := GetRandomNumberInRange(10, 20)
		if randNumber > 20 || randNumber < 10 {
			t.Errorf("TestRandoms - random number :: received  number not in range: %d", randNumber)
		}
	})

	t.Run("Assert random float returns a float between the range", func(t *testing.T) {
		randFloat := GetRandomFloatInRange(10, 20)
		if randFloat > 20 || randFloat < 10 {
			t.Errorf("TestRandoms - random float :: received  float not in range: %f", randFloat)
		}
	})

	t.Run("Negative mins and maxs return 0 - int", func(t *testing.T) {
		shouldBeZero := GetRandomNumberInRange(-1, -1)

		if shouldBeZero != 0 {
			t.Errorf("TestRandoms - negative mins and maxes :: did not return 0, got %d", shouldBeZero)
		}
	})

	t.Run("Negative mins and maxs return 0 - float", func(t *testing.T) {
		shouldBeLessThanOne := GetRandomFloatInRange(-1, -1)

		if shouldBeLessThanOne >= 1 {
			t.Errorf("TestRandoms - negative mins and maxes :: did not return correct, got %f", shouldBeLessThanOne)
		}
	})
}
