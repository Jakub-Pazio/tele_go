package zad1

import (
	"testing"
)

func TestCheckBit(t *testing.T) {
	t.Run("Tested bit is turned on", func(t *testing.T) {
		firstBitOn := uint8(0x1)

		want := true
		got, err := CheckBit(firstBitOn, 0)

		if got != want {
			t.Errorf("wanted %t got %t", want, got)
		}
		if err != nil {
			t.Error("error not expected")
		}
	})

	t.Run("Tested bit is not turned on", func(t *testing.T) {
		firstBitOff := uint8(0xFE)

		want := false
		got, err := CheckBit(firstBitOff, 0)

		if got != want {
			t.Errorf("wanted %t got %t", want, got)
		}
		if err != nil {
			t.Error("error not expected")
		}
	})

}
