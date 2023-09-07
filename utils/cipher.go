package utils

import (
	"dev-xero/caesar-cipher/log"
)

func GenerateCipher(word string, shift int) string {
	shiftedLetters := []rune{}

	for i := 0; i < len(word); i++ {
		char := word[i]

		if char < 'a' || char > 'z' {
			log.PrintError("[ERR]: Invalid letter encountered")
		}

		index := (int(char-'a') + shift) % 26
		if index < 0 {
			index += 26
		}

		shiftedLetters = append(shiftedLetters, rune('a'+index))
	}

	return string(shiftedLetters)
}
