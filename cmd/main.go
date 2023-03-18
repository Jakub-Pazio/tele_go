package main

import (
	"flag"
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
