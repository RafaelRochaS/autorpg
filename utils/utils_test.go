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
			t.Errorf("TestRandoms :: random number : received  number not in range: %d", randNumber)
		}
	})

	t.Run("Assert random float returns a float between the range", func(t *testing.T) {
		randFloat := GetRandomFloatInRange(10, 20)
		if randFloat > 20 || randFloat < 10 {
			t.Errorf("TestRandoms :: random float : received  float not in range: %f", randFloat)
		}
	})
}
