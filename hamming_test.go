package zad1

import (
	"fmt"
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

func TestSetParity(t *testing.T) {
	t.Run("Test ParityBits setting", func(t *testing.T) {
		valueToEncode := uint8(0xF7)

		want := uint16(0x1E78)
		got := SetParity(valueToEncode)

		if got != want {
			t.Errorf("wanted %d got %d", want, got)
		}
	})
}

func TestDecoding(t *testing.T) {
	t.Run("Decode good data", func(t *testing.T) {
		val := uint8(0x1)
		valToDecode := SetParity(val)

		if valToDecode != uint16(0xE) {
			t.Errorf("Error in encoding, wanted 14 got %d", valToDecode)
		}
		got := DecodeData(valToDecode)
		want := uint8(0x1)
		if got != want {
			t.Errorf("wanted %d got %d", want, got)
		}
	})
}

func TestCorrectData(t *testing.T) {
	t.Run("Data is good so dont change", func(t *testing.T) {
		valueToChange := uint16(0x1E78)

		CorrectData(&valueToChange)
		if valueToChange != uint16(0x1E78) {
			t.Error("Data corrected for no reason")
		}
	})
	t.Run("Data is bad, change", func(t *testing.T) {
		valueToChange := uint16(0x1E70)

		CorrectData(&valueToChange)
		if valueToChange != uint16(0x1E78) {
			t.Errorf("got %d want 7792", valueToChange)
		}
	})
	t.Run("Test flip for data value 1", func(t *testing.T) {
		value := uint8(0x1)
		encodedValue := EncodeData(value)
		if encodedValue != uint16(0x8) {
			t.Errorf("got %d want 8", encodedValue)
		}
		afterHamming := SetParity(value)
		if afterHamming != uint16(0xE) {
			t.Errorf("got %d wanted 14", afterHamming)
		}
		bitFlipped := uint16(0x20E)
		CorrectData(&bitFlipped)
		if bitFlipped != uint16(0xE) {
			t.Errorf("got %d wanted 14", bitFlipped)
		}

		//writeToFile("we", encryptFile("./test.txt"))

		fromFile := readEncrypted("")
		decryptedArray := make([]uint8, 0)

		for i := 0; i < len(fromFile); i++ {
			decryptedArray = append(decryptedArray, DecodeData(fromFile[i]))
		}

		// decrypted := DecodeData(fromFile)
		// decArray := []uint8{decrypted}
		// writeDecryptedToFile("", decryptedArray)
		// //fmt.Printf("decrypted: %x\n", decryptedArray[])

		// linearResult := EncodeTo16Bits(0x4)
		// fmt.Printf("%d\n", linearResult)

		// decodedResult := DecodeTo8Bits(linearResult)
		// fmt.Printf("%d\n", decodedResult)

		stupidTest := uint8(0x5)
		stupidEncoded := StupidEncoder(stupidTest)
		stupidDecrypted := StupidDecoder(stupidEncoded)
		fmt.Printf("%d\n%d\n", stupidEncoded, stupidDecrypted)
	})
}

func TestFileRead(t *testing.T) {

}
