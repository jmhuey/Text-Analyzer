package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"
)

func main() {

	// open the file
	file, err := os.Open("test.txt")
	if err != nil {
		fmt.Println(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanRunes)

	char := make(map[string]*int)

	for scanner.Scan() {
		token := scanner.Text()

		// Temporarily added to focus on letters only -----------------------
		if !isLetter(token) && !isNumber(token) {
			continue
		} else if isNumber(token) {
			continue
		}
		// ------------------------------------------------------------------

		if isLetter(token) {
			token = strings.ToLower(token)
		}

		switch []rune(token)[0] {
		case ' ', '\t', '\n', '\f', '\r', '\v':
			continue

		default:
			if val, ok := char[token]; ok {
				*val++
			} else {
				t := 1
				char[token] = &t
			}
		}

	}

	letters_found := make([]string, 0, len(char))
	letters_notFound := make([]string, 0)

	for k := range char {
		letters_found = append(letters_found, k)
	}

	sort.Strings(letters_found)

	// checks if there were any letters not used in the text
	for l := 'a'; l < 'z'; l++ {
		if _, ok := char[string(l)]; !ok {
			letters_notFound = append(letters_notFound, string(l))
		}
	}

	fmt.Println("These are the letters found in the text:")
	for _, letter := range letters_found {
		fmt.Println(letter, *char[letter])
	}

	if len(letters_notFound) != 0 {
		fmt.Println("These are the letters not found in the text:")

		for _, letter := range letters_notFound {
			fmt.Println(letter)
		}
	}

}

// checks if string passed is a letter
func isLetter(s string) bool {
	is_letter := regexp.MustCompile(`^[a-zA-Z]*$`).MatchString(s)
	if is_letter {
		return true
	}
	return false
}

// checks if string passed is a number
func isNumber(s string) bool {
	is_number := regexp.MustCompile(`^[0-9]*$`).MatchString(s)
	if is_number {
		return true
	}
	return false
}
