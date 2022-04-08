package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
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
		}
		// ------------------------------------------------------------------

		if isLetter(token) {
			token = strings.ToLower(token)
		}

		// ignores the whitespace | if the token is in the map add 1 to the value otherwise add the token to map
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

	letters_found := make([]string, 0)
	letters_notFound := make([]string, 0)

	num_found := make([]string, 0)
	num_notFound := make([]int, 0)

	// sorting the characters found
	for k := range char {
		if isLetter(k) {
			letters_found = append(letters_found, k)
		} else if isNumber(k) {
			num_found = append(num_found, k)
		}
	}

	// sort letters and numbers alphabetically and numerically respectively
	sort.Strings(letters_found)
	sort.Strings(num_found)

	// checks if there were any letters not found in the text
	for l := 'a'; l < 'z'; l++ {
		if _, ok := char[string(l)]; !ok {
			letters_notFound = append(letters_notFound, string(l))
		}
	}

	// checks if there were any numbers not found in text
	for n := 0; n < 10; n++ {
		if _, ok := char[strconv.Itoa(n)]; !ok {
			num_notFound = append(num_notFound, n)
		}
	}

	// prints the numbers found
	fmt.Println("These are the numbers found in the text:")

	for _, num := range num_found {
		fmt.Println(num, ": ", *char[num])
	}

	// prints out numbers not found if any
	if len(num_notFound) != 0 {
		fmt.Println()
		fmt.Println("These are the numbers not found in the text:")

		for _, num := range num_notFound {
			fmt.Println(num)
		}
	}

	// prints out letters found
	fmt.Println()
	fmt.Println("These are the letters found in the text:")

	for _, letter := range letters_found {
		fmt.Println(letter, ": ", *char[letter])
	}

	// prints out letters not found if any
	if len(letters_notFound) != 0 {
		fmt.Println()
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
