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

func encodeData(input byte) []byte {
	return nil
}
