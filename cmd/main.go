package main

import (
	"flag"
	"fmt"
	"os"
	correction "tele_zad1/correctionCodes"
)

func main() {
	var encode bool
	var decode bool

	var inputFile string
	var outputFile string

	flag.BoolVar(&encode, "e", false, "encode file")
	flag.BoolVar(&decode, "d", false, "decode file")
	flag.StringVar(&inputFile, "in", "", "input file")
	flag.StringVar(&outputFile, "out", "", "output file")

	flag.Parse()

	bits2Correction := uint8(0x3)

	linear := correction.EncodeTo16Bits(bits2Correction)
	fmt.Printf("%d\n", linear)

	decodedLinear := correction.DecodeTo8Bits(linear)
	fmt.Printf("%d\n", decodedLinear)

	checkMatrix := correction.LinerCheck(linear)
	fmt.Printf("%d\n", checkMatrix)

	checkMatrixFake := correction.LinerCheck(572)
	fmt.Printf("%d\n", checkMatrixFake)

	checkMatrixFake2 := correction.LinerCheck(1596)
	fmt.Printf("%d\n", checkMatrixFake2)
	// checkMatrixFake := correction.LinerCheck(24690)
	// fmt.Printf("%d\n", checkMatrixFake)

	flippedValue := uint8(0x6)
	correction.CorrectMistake(&flippedValue, checkMatrixFake2)
	fmt.Printf("%d\n", flippedValue)

	if encode {
		data := correction.EncryptFile(inputFile)
		correction.WriteEncryptedToFile(outputFile, data)
		return
	}
	if decode {
		data := correction.DecryptFile(inputFile)
		correction.WriteDecryptedToFile(outputFile, data)
		return
	}
	os.Exit(2)
}
