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

func TestSetBit(t *testing.T) {
	t.Run("Set bit when bit is 0", func(t *testing.T) {
		firstBitoff := uint16(0xFFE)

		want := uint(0xFFF)
		SetBit(&firstBitoff, 0)

		if uint(firstBitoff) != want {
			t.Errorf("wanted %d got %d", want, firstBitoff)
		}
	})
	t.Run("Set bit when bit is 1", func(t *testing.T) {
		firstBiton := uint16(0xDDB)

		want := uint16(0xDDB)
		SetBit(&firstBiton, 0)

		if firstBiton != want {
			t.Errorf("wanted %d got %d", want, firstBiton)
		}
	})
}

func TestEncode(t *testing.T) {
	t.Run("Test basic Hamming encoding", func(t *testing.T) {
		valueToEncode := uint8(0xF7)

		want := uint16(0x1E68)
		got := EncodeData(valueToEncode)

		if got != want {
			t.Errorf("wanted %d got %d", want, got)
		}
	})
}
