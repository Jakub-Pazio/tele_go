package correction

import (
	"fmt"
	"math/bits"
)

var DecodingMatrix = [8]uint16{
	uint16(0xF080), // {1, 1, 1, 1, 0, 0, 0, 0,   1, 0, 0, 0, 0, 0, 0, 0}
	uint16(0xCC40), // {1, 1, 0, 0, 1, 1, 0, 0,   0, 1, 0, 0, 0, 0, 0, 0}
	uint16(0xAA20), // {1, 0, 1, 0, 1, 0, 1, 0,   0, 0, 1, 0, 0, 0, 0, 0}
	uint16(0x5610), // {0, 1, 0, 1, 0, 1, 1, 0,   0, 0, 0, 1, 0, 0, 0, 0}
	uint16(0xE908), // {1, 1, 1, 0, 1, 0, 0, 1,   0, 0, 0, 0, 1, 0, 0, 0}
	uint16(0x9504), // {1, 0, 0, 1, 0, 1, 0, 1,   0, 0, 0, 0, 0, 1, 0, 0}
	uint16(0x7B02), // {0, 1, 1, 1, 1, 0, 1, 1,   0, 0, 0, 0, 0, 0, 1, 0}
	uint16(0xE701), // {1, 1, 1, 0, 0, 1, 1, 1,   0, 0, 0, 0, 0, 0, 0, 1}
}

var DecodingMatrix2 = [8]uint8{
	uint8(0xF0), // {1, 1, 1, 1, 0, 0, 0, 0,   1, 0, 0, 0, 0, 0, 0, 0}
	uint8(0xCC), // {1, 1, 0, 0, 1, 1, 0, 0,   0, 1, 0, 0, 0, 0, 0, 0}
	uint8(0xAA), // {1, 0, 1, 0, 1, 0, 1, 0,   0, 0, 1, 0, 0, 0, 0, 0}
	uint8(0x56), // {0, 1, 0, 1, 0, 1, 1, 0,   0, 0, 0, 1, 0, 0, 0, 0}
	uint8(0xE9), // {1, 1, 1, 0, 1, 0, 0, 1,   0, 0, 0, 0, 1, 0, 0, 0}
	uint8(0x95), // {1, 0, 0, 1, 0, 1, 0, 1,   0, 0, 0, 0, 0, 1, 0, 0}
	uint8(0x7B), // {0, 1, 1, 1, 1, 0, 1, 1,   0, 0, 0, 0, 0, 0, 1, 0}
	uint8(0xE7), // {1, 1, 1, 0, 0, 1, 1, 1,   0, 0, 0, 0, 0, 0, 0, 1}
}

var CodingMatrix = [16]uint8{
	uint8(0xED),
	uint8(0xDB),
	uint8(0xAB),
	uint8(0x96),
	uint8(0x6A),
	uint8(0x55),
	uint8(0x33),
	uint8(0xF),
	uint8(0x80),
	uint8(0x40),
	uint8(0x20),
	uint8(0x10),
	uint8(0x8),
	uint8(0x4),
	uint8(0x2),
	uint8(0x1),
}

// Does not work
func EncodeTo16Bits(input uint8) uint16 {
	result := uint16(0x0)

	for i := 0; i < 16; i++ {
		xored := CodingMatrix[i] & input
		numberOfBitsUp := bits.OnesCount8(xored)
		isOdd := numberOfBitsUp & 0x1
		fmt.Printf("%d\n", isOdd)

		// TODO: Flip corresponding bit in the result
		if isOdd != 0 {
			SetBit(&result, 15-i)
		}
	}

	return result
}

// Does not work
func DecodeTo8Bits(input uint16) uint8 {
	result := uint8(input & 255)
	//input >>= 8
	checkError := uint8(input >> 8 & 255)
	//checkError = bits.Reverse8(checkError)
	//isError := uint8(0x0)
	fmt.Printf("res: %d check: %d\n", result, checkError)

	for i := 0; i < 8; i++ {
		xored := DecodingMatrix[7-i] & input
		//numberOfBitsUp := bits.OnesCount8(xored)
		//isOdd := numberOfBitsUp & 0x1
		fmt.Printf("%d\n", bits.OnesCount16(xored))
	}

	return result
}

func RepeatEncoder(input uint8) uint64 {
	result := uint64(0x0)

	for i := 0; i < 8; i++ {
		val, _ := CheckBit8(input, 7-i)
		result <<= 8
		if val {
			result = result | uint64(0xFF)
		}
	}
	return result
}

func RepeatDecoder(input uint64) uint8 {
	result := uint8(0x0)
	for i := 0; i < 8; i++ {
		temp := uint8(input)
		fmt.Printf("%d\n", bits.OnesCount8(temp))
		if bits.OnesCount8(temp) > 4 {
			result |= uint8(0x1) << i
		}
		input >>= 8
	}
	return result
}
