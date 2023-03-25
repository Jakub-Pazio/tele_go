package correction

import (
	"encoding/binary"
	"errors"
	"fmt"
	"math"
	"os"
)

var ParityIndexes = [4][8]int{
	{1, 3, 5, 7, 9, 11, 13, 15},
	{2, 3, 6, 7, 10, 11, 14, 15},
	{4, 5, 6, 7, 12, 13, 14, 15},
	{8, 9, 10, 11, 12, 13, 14, 15},
}

func CheckBit8(value uint8, possition int) (bool, error) {
	if possition > 7 {
		return false, errors.New("wrong possition")
	}
	val := value & (uint8(0x1) << possition)

	if val != 0 {
		return true, nil
	}
	return false, nil
}
func CheckBit16(value uint16, possition int) (bool, error) {
	if possition > 15 {
		return false, errors.New("wrong possition")
	}
	val := value & (uint16(0x1) << possition)

	if val != 0 {
		return true, nil
	}
	return false, nil
}

func SetBit16(value *uint16, possition int) error {
	if possition > 15 {
		return errors.New("wrong possition")
	}
	*value = *value | uint16(0x1)<<possition

	return nil
}

func SetBit8(value *uint8, possition int) error {
	if possition > 7 {
		return errors.New("wrong possition")
	}
	*value = *value | uint8(0x1)<<possition

	return nil
}
func FlipBit16(value *uint16, possition int) error {
	if possition > 15 {
		return errors.New("wrong possition")
	}
	*value = *value ^ uint16(0x1)<<possition

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
func preHammingEncode(input uint8) uint16 {
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
		isSet, _ := CheckBit8(input, i)
		if isSet {
			SetBit16(&result, v)
		}
	}
	return result
}

func HammingEncode(input uint8) uint16 {
	input16 := preHammingEncode(input)
	for i, indexArray := range ParityIndexes {
		sum := 0
		for _, v := range indexArray {
			isSet, _ := CheckBit16(input16, v)
			if isSet {
				sum++
			}
		}
		if sum%2 == 1 {
			SetBit16(&input16, int(math.Pow(2, float64(i))))
		}
	}

	return input16
}

func HammingDecode(intput uint16) uint8 {
	HammingDataCorrect(&intput)
	result := uint8(0)
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
		isSet, _ := CheckBit16(intput, v)
		if isSet {
			SetBit8(&result, i)
		}
	}
	return result
}

func HammingDataCorrect(input *uint16) {
	indexToCorrect := 0
	for i := 0; i < 16; i++ {
		on, _ := CheckBit16(*input, i)
		if on {
			indexToCorrect = indexToCorrect ^ i
		}
	}
	if indexToCorrect != 0 {
		FlipBit16(input, indexToCorrect)
	}
}

func readFileToVec(name string) []uint8 {
	array, _ := os.ReadFile(name)
	return array
}

func EncryptFile(name string, encoding func(uint8) uint16) []uint16 {
	fromFileArray := readFileToVec(name)
	retsultArray := make([]uint16, 0)

	for _, v := range fromFileArray {
		retsultArray = append(retsultArray, encoding(v))
	}

	return retsultArray
}

func DecryptFile(name string, decoding func(uint16) uint8) []uint8 {
	fromFileArray := ReadEncryptedFromFile(name)
	decryptedArray := make([]uint8, 0)

	for i := 0; i < len(fromFileArray); i++ {
		decryptedArray = append(decryptedArray, decoding(fromFileArray[i]))
	}

	return decryptedArray
}

func WriteEncryptedToFile(name string, data []uint16) {
	f, _ := os.OpenFile(name, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	err := binary.Write(f, binary.BigEndian, data)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
	}
}

func ReadEncryptedFromFile(name string) []uint16 {
	array, _ := os.ReadFile(name)
	resultNumber := make([]uint16, 0)
	for i := 0; i < len(array)/2; i++ {
		number := binary.BigEndian.Uint16(array[i*2 : i*2+2])
		//CorrectData(&number)
		resultNumber = append(resultNumber, number)
	}
	return resultNumber
}

func WriteDecryptedToFile(name string, data []uint8) {
	f, _ := os.OpenFile(name, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	err := binary.Write(f, binary.BigEndian, data)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
	}
}
