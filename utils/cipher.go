package utils

import (
	"dev-xero/caesar-cipher/log"
	"errors"
	"os"
	"unicode"
)

func GenerateCipher(word string, shift int) string {
	encryptedText := ""

	for _, char := range word {
		if unicode.IsLetter(char) {
			shiftedChar := shiftChar(char, shift)
			encryptedText += string(shiftedChar)
		} else if unicode.IsDigit(char) {
			shiftedDigit := shiftDigit(char, shift)
			encryptedText += string(shiftedDigit)
		} else {
			log.Error(errors.New("invalid character encountered"))
			os.Exit(1)
		}
	}

	return encryptedText
}

func shiftChar(char rune, shift int) rune {
	var base rune
	if unicode.IsUpper(char) {
		base = 'A'
	} else {
		base = 'a'
	}

	return ((char-base)+rune(shift)+26)%26 + base
}

func shiftDigit(char rune, shift int) rune {
	return ((char-'0')+rune(shift)+10)%10 + '0'
}
