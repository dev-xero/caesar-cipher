package main

import (
	"dev-xero/caesar-cipher/log"
	"fmt"
	"os"
)

func main() {
	log.PrintTitle("Caesar Cipher")

	var word string
	fmt.Print("Word: ")
	_, err := fmt.Scanln(&word)
	if err != nil {
		log.PrintError(err)
		os.Exit(1)
	}

	fmt.Println(word)
}
