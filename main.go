package main

import (
	"fmt"
	"os"
)

func main() {

	// open the file
	file, err := os.Open("test.txt")
	if err != nil {
		fmt.Println(err)
	}

	char := scanFile(file)

	letters_found, num_found := sortCharacters(char)
	letters_notFound, num_notFound := make([]string, 0), make([]int, 0)

	// checks if there were any letters not found in the text
	checkMissingLetters(&letters_notFound, char)

	// checks if there were any numbers not found in text
	checkMissingNumbers(&num_notFound, char)

	// prints out letters found
	printLettersFound(letters_found, char)

	// prints out letters not found if any
	printLettersNotFound(letters_notFound)

	// prints the numbers found
	printNumbersFound(num_found, char)

	// prints out numbers not found if any
	printNumbersNotFound(num_notFound)

}
