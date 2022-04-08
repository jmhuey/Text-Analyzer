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

	keys := make([]string, 0, len(char))

	for k := range char {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	for _, k := range keys {
		fmt.Println(k, *char[k])

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
