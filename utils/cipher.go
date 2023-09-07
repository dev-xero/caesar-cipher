package utils

import (
	"fmt"
	"log"
	"strings"
)

func GenerateCipher(word string, shift int) string {
	shiftedLetters := []string{}

	for i := 0; i < len(word); i++ {
		char := word[i]

		if char < 'a' || char > 'z' {
			log.Fatal("[ERR]: Invalid letter encountered")
		}

		index := (int(char - 'a') + shift) % 26
		if index < 0 {
			index += 26
		}

		shiftedLetters = append(shiftedLetters, fmt.Sprintf("%d", index))
	}

	return strings.Join(shiftedLetters, "")
}