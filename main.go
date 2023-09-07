package main

import (
	"dev-xero/caesar-cipher/log"
	"dev-xero/caesar-cipher/utils"
	"fmt"
)

func main() {
	log.PrintTitle("Caesar Cipher")

	cipher := utils.GenerateCipher("password", 8)
	fmt.Println(cipher)
}
