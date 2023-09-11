package main

import (
	"dev-xero/caesar-cipher/log"
	"dev-xero/caesar-cipher/utils"
	"errors"
	"flag"
	"fmt"
	"os"
	"strconv"
)

func main() {
	log.Title("CAESAR CIPHER")
	var export bool
	flag.BoolVar(&export, "e", false, "Option to export the generated cipher")

	flag.Parse()

	if flag.NArg() < 2 {
		log.Error(errors.New("not enough arguments supplied"))
		os.Exit(1)
	}

	if word := flag.Arg(0); word != "" {
		if shift, err := strconv.Atoi(flag.Arg(1)); err == nil {
			encryptedText := utils.Cipher(word, shift)

			fmt.Println("Text:", word)
			fmt.Println("Shift:", shift)
			log.Result(encryptedText)
		} else {
			log.Error(errors.New("shift is not a valid integer"))
			os.Exit(1)
		}
	} else {
		log.Error(errors.New("a text to encrypt must be supplied"))
		os.Exit(1)
	}
}
