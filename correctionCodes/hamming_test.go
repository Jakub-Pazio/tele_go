package correction

import (
	"testing"
)

func TestCheckBit(t *testing.T) {
	t.Run("Tested bit is turned on", func(t *testing.T) {
		firstBitOn := uint8(0x1)

		want := true
		got, err := CheckBit8(firstBitOn, 0)

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
		got, err := CheckBit8(firstBitOff, 0)

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
		SetBit16(&firstBitoff, 0)

		if uint(firstBitoff) != want {
			t.Errorf("wanted %d got %d", want, firstBitoff)
		}
	})
	t.Run("Set bit when bit is 1", func(t *testing.T) {
		firstBiton := uint16(0xDDB)

		want := uint16(0xDDB)
		SetBit16(&firstBiton, 0)

		if firstBiton != want {
			t.Errorf("wanted %d got %d", want, firstBiton)
		}
	})
}

func TestEncode(t *testing.T) {
	t.Run("Test basic Hamming encoding", func(t *testing.T) {
		valueToEncode := uint8(0xF7)

		want := uint16(0x1E68)
		got := preHammingEncode(valueToEncode)

		if got != want {
			t.Errorf("wanted %d got %d", want, got)
		}
	})
}

func TestSetParity(t *testing.T) {
	t.Run("Test ParityBits setting", func(t *testing.T) {
		valueToEncode := uint8(0xF7)

		want := uint16(0x1E78)
		got := HammingEncode(valueToEncode)

		if got != want {
			t.Errorf("wanted %d got %d", want, got)
		}
	})
}

func TestDecoding(t *testing.T) {
	t.Run("Decode good data", func(t *testing.T) {
		val := uint8(0x1)
		valToDecode := HammingEncode(val)

		if valToDecode != uint16(0xE) {
			t.Errorf("Error in encoding, wanted 14 got %d", valToDecode)
		}
		got := HammingDecode(valToDecode)
		want := uint8(0x1)
		if got != want {
			t.Errorf("wanted %d got %d", want, got)
		}
	})
}

func TestCorrectData(t *testing.T) {
	t.Run("Data is good so dont change", func(t *testing.T) {
		valueToChange := uint16(0x1E78)

		HammingDataCorrect(&valueToChange)
		if valueToChange != uint16(0x1E78) {
			t.Error("Data corrected for no reason")
		}
	})
	t.Run("Data is bad, change", func(t *testing.T) {
		valueToChange := uint16(0x1E70)

		HammingDataCorrect(&valueToChange)
		if valueToChange != uint16(0x1E78) {
			t.Errorf("got %d want 7792", valueToChange)
		}
	})
	t.Run("Test flip for data value 1", func(t *testing.T) {
		value := uint8(0x1)
		encodedValue := preHammingEncode(value)
		if encodedValue != uint16(0x8) {
			t.Errorf("got %d want 8", encodedValue)
		}
		afterHamming := HammingEncode(value)
		if afterHamming != uint16(0xE) {
			t.Errorf("got %d wanted 14", afterHamming)
		}
		bitFlipped := uint16(0x20E)
		HammingDataCorrect(&bitFlipped)
		if bitFlipped != uint16(0xE) {
			t.Errorf("got %d wanted 14", bitFlipped)
		}
	})
}
