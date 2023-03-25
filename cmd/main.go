package main

import (
	"flag"
	"os"
	correction "tele_zad1/correctionCodes"
)

func main() {
	var encode bool
	var decode bool
	var hamming bool
	var matrix bool

	var inputFile string
	var outputFile string

	var encodef func(uint8) uint16
	var decodef func(uint16) uint8

	flag.BoolVar(&encode, "e", false, "encode file")
	flag.BoolVar(&decode, "d", false, "decode file")
	flag.BoolVar(&hamming, "h", false, "hamming encoding")
	flag.BoolVar(&matrix, "m", false, "hamming encoding")
	flag.StringVar(&inputFile, "in", "", "input file")
	flag.StringVar(&outputFile, "out", "", "output file")

	flag.Parse()

	if hamming {
		encodef = correction.SetParity
		decodef = correction.DecodeData
	} else if matrix {
		encodef = correction.EncodeTo16Bits
		decodef = correction.DecodeTo8Bits
	} else {
		os.Exit(1)
	}

	if encode {
		data := correction.EncryptFile(inputFile, encodef)
		correction.WriteEncryptedToFile(outputFile, data)
		return
	}
	if decode {
		data := correction.DecryptFile(inputFile, decodef)
		correction.WriteDecryptedToFile(outputFile, data)
		return
	}
	os.Exit(1)
}
