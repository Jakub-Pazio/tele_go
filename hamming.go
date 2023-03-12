package zad1

import (
	"errors"
)

func CheckBit(value uint8, possition int) (bool, error) {
	if possition > 7 {
		return false, errors.New("Wrong possition")
	}
	val := value & (uint8(0x1) << possition)

	if val != 0 {
		return true, nil
	}
	return false, nil
}

func SetBit(value *uint16, possition int) error {
	if possition > 15 {
		return errors.New("Wrong possition")
	}
	*value = *value | uint16(0x1)<<possition

	return nil
}

// I count bits from right to left
// We want to encode data on bits with numbers:
//
//	data: 7  6  5  4  3  2  1  0
//	 out: 12 11 10 9  7  6  5  3
//
// Bits 1 2 4 8 are left for parity check; Bit 0 for whole block check
// Other bits (13, 14, 15) are not needed and will have 0 assigned
func EncodeData(input uint8) uint16 {
	result := uint16(0)
	bitMapping := [8]int{
		0: 3,
		1: 5,
		2: 6,
		3: 7,
		4: 9,
		5: 10,
		6: 11,
		7: 12,
	}
	for i, v := range bitMapping {
		isSet, _ := CheckBit(input, i)
		if isSet == true {
			SetBit(&result, v)
		}
	}

	return result
}
