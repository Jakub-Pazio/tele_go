package zad1

import (
	"fmt"
	"math/bits"
)

var verMatrix = [8]uint16{
	uint16(0xF080), // {1, 1, 1, 1, 0, 0, 0, 0,   1, 0, 0, 0, 0, 0, 0, 0}
	uint16(0xCC40), // {1, 1, 0, 0, 1, 1, 0, 0,   0, 1, 0, 0, 0, 0, 0, 0}
	uint16(0xAA20), // {1, 0, 1, 0, 1, 0, 1, 0,   0, 0, 1, 0, 0, 0, 0, 0}
	uint16(0x5610), // {0, 1, 0, 1, 0, 1, 1, 0,   0, 0, 0, 1, 0, 0, 0, 0}
	uint16(0xE908), // {1, 1, 1, 0, 1, 0, 0, 1,   0, 0, 0, 0, 1, 0, 0, 0}
	uint16(0x9504), // {1, 0, 0, 1, 0, 1, 0, 1,   0, 0, 0, 0, 0, 1, 0, 0}
	uint16(0x7B02), // {0, 1, 1, 1, 1, 0, 1, 1,   0, 0, 0, 0, 0, 0, 1, 0}
	uint16(0xE701), // {1, 1, 1, 0, 0, 1, 1, 1,   0, 0, 0, 0, 0, 0, 0, 1}
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

func EncodeTo16Bits(input uint8) uint16 {
	result := uint16(0x0)

	for i := 0; i < 16; i++ {
		xored := CodingMatrix[i] & input
		fmt.Printf("%d\n", xored)
		numberOfBitsUp := bits.OnesCount8(xored)
		isOdd := numberOfBitsUp & 0x1

		// TODO: Flip corresponding bit in the result
		if isOdd != 0 {
			SetBit(&result, 15-i)
		}
	}

	return result
}
