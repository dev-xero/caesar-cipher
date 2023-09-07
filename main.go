package main

import (
	"dev-xero/caesar-cipher/log"
	"dev-xero/caesar-cipher/utils"
	"fmt"
	"os"
)

func main() {
	log.Title("Caesar Cipher")

	var word string
	fmt.Print("Word: ")
	_, err := fmt.Scanln(&word)
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}

	var shift int
	fmt.Print("Shift: ")
	_, err = fmt.Scanln(&shift)
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}

	cipher := utils.GenerateCipher(word, shift)
	log.Result(cipher)
}
