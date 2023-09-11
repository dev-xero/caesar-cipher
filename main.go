package main

import (
	"bufio"
	"dev-xero/caesar-cipher/log"
	"dev-xero/caesar-cipher/utils"
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
)

func exportResult(encryptedText string, shift int) {
	subdirectory, filename := "exported", fmt.Sprintf("%s_%d.txt", encryptedText, shift)
	fullPath := filepath.Join(subdirectory, filename)

	if err := os.MkdirAll("exported", os.ModePerm); err != nil {
		log.Error(errors.New("could not create directory"))
		os.Exit(1)
	}

	file, err := os.Create(fullPath)
	if err != nil {
		log.Error(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	text := fmt.Sprintf("encrypted text: %s\ndecryption key: %d", encryptedText, -shift)
	_, err = writer.WriteString(text)
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}

	err = writer.Flush()
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}

	fmt.Printf("File '%s' has been written successfully in subdirectory '%s'.\n", filename, subdirectory)
}

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

			if export {
				exportResult(encryptedText, shift)
			}
		} else {
			log.Error(errors.New("shift is not a valid integer"))
			os.Exit(1)
		}
	} else {
		log.Error(errors.New("a text to encrypt must be supplied"))
		os.Exit(1)
	}
}
