package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	fmt.Println("Listening to port 8080")
	http.HandleFunc("/", parseFile)
	http.ListenAndServe(":8080", nil)

}

func parseFile(w http.ResponseWriter, request *http.Request) {
	// open the file
	file, err := os.Open("test.txt")
	if err != nil {
		fmt.Println(err)
	}

	//---------------------------
	// Character operation calls
	//---------------------------

	char := scanFileChar(file)

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

	//----------------------
	// Word operation calls
	//----------------------

	word := scanFileWord(file)

	for key, value := range word {
		fmt.Println(key, ":", value)
	}
}
