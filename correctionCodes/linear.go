package correction

import (
	"math/bits"
)

var EncodingMatixFull = [8]uint16{
	uint16(0xF080), // {1, 1, 1, 1, 0, 0, 0, 0,   1, 0, 0, 0, 0, 0, 0, 0}
	uint16(0xCC40), // {1, 1, 0, 0, 1, 1, 0, 0,   0, 1, 0, 0, 0, 0, 0, 0}
	uint16(0xAA20), // {1, 0, 1, 0, 1, 0, 1, 0,   0, 0, 1, 0, 0, 0, 0, 0}
	uint16(0x5610), // {0, 1, 0, 1, 0, 1, 1, 0,   0, 0, 0, 1, 0, 0, 0, 0}
	uint16(0xE908), // {1, 1, 1, 0, 1, 0, 0, 1,   0, 0, 0, 0, 1, 0, 0, 0}
	uint16(0x9504), // {1, 0, 0, 1, 0, 1, 0, 1,   0, 0, 0, 0, 0, 1, 0, 0}
	uint16(0x7B02), // {0, 1, 1, 1, 1, 0, 1, 1,   0, 0, 0, 0, 0, 0, 1, 0}
	uint16(0xE701), // {1, 1, 1, 0, 0, 1, 1, 1,   0, 0, 0, 0, 0, 0, 0, 1}
}

var EncodingMatrixPart = [8]uint8{
	uint8(0xF0), // {1, 1, 1, 1, 0, 0, 0, 0, |  1, 0, 0, 0, 0, 0, 0, 0}
	uint8(0xCC), // {1, 1, 0, 0, 1, 1, 0, 0, |  0, 1, 0, 0, 0, 0, 0, 0}
	uint8(0xAA), // {1, 0, 1, 0, 1, 0, 1, 0, |  0, 0, 1, 0, 0, 0, 0, 0}
	uint8(0x56), // {0, 1, 0, 1, 0, 1, 1, 0, |  0, 0, 0, 1, 0, 0, 0, 0}
	uint8(0xE9), // {1, 1, 1, 0, 1, 0, 0, 1, |  0, 0, 0, 0, 1, 0, 0, 0}
	uint8(0x95), // {1, 0, 0, 1, 0, 1, 0, 1, |  0, 0, 0, 0, 0, 1, 0, 0}
	uint8(0x7B), // {0, 1, 1, 1, 1, 0, 1, 1, |  0, 0, 0, 0, 0, 0, 1, 0}
	uint8(0xE7), // {1, 1, 1, 0, 0, 1, 1, 1, |  0, 0, 0, 0, 0, 0, 0, 1}
}

var DecodingMatrix = [16]uint8{
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

func MatrixEncoding(input uint8) uint16 {
	result := uint16(input)
	result <<= 8

	for i := 0; i < 8; i++ {
		xored := EncodingMatrixPart[i] & input
		numberOfBitsUp := bits.OnesCount8(xored)
		isOdd := numberOfBitsUp & 0x1
		if isOdd != 0 {
			SetBit16(&result, 7-i)
		}
	}

	return result
}

func MatrixDecoding(input uint16) uint8 {
	result := uint8(input >> 8 & 255)
	mistakeMatix := ErrorVector(input)
	MatrixErrorCorrection(&result, mistakeMatix)

	return result
}

func ErrorVector(input uint16) uint8 {
	result := uint8(0x0)
	for i := 0; i < 8; i++ {
		multiMatrix := bits.OnesCount16(input & EncodingMatixFull[i])
		parity := multiMatrix % 2
		if parity != 0 {
			result |= 0x1 << (7 - uint8(i))
		}
		// trzeba zobaczyć w macierzy transponowanej
		// która z tych wartości odpowiada blędowi i ją flipnąć
		// to działa dla będów 1 bitowych
		// dla błędów 2 bitowych inaczej jakoś
	}
	return result
}

func MatrixErrorCorrection(input *uint8, mistake uint8) {
	for i := 0; i < 8; i++ {
		if mistake == DecodingMatrix[i] {
			*input ^= 0x1 << uint8(7-i)
		}
	}
	for i := 0; i < 16; i++ {
		for j := i; j < 16; j++ {
			mistakeMask := DecodingMatrix[i] ^ DecodingMatrix[j]
			if mistakeMask == mistake {
				*input ^= 0x1 << uint8(7-i)
				*input ^= 0x1 << uint8(7-j)
			}
		}
	}
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
		//fmt.Printf("%d\n", bits.OnesCount8(temp))
		if bits.OnesCount8(temp) > 4 {
			result |= uint8(0x1) << i
		}
		input >>= 8
	}
	return result
}
