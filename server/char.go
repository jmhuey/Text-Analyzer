package main

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// scans the file and puts unique values into a map while recording the number of times a repeated value appears
func scanFileChar(f *os.File) map[string]map[string]*int {
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanRunes)

	//char := make(map[string]*int)
	charFound := make(map[string]map[string]*int)
	charFound["Letters"] = map[string]*int{}
	charFound["Numbers"] = map[string]*int{}
	charFound["LettersNotFound"] = map[string]*int{}
	charFound["NumbersNotFound"] = map[string]*int{}

	for scanner.Scan() {
		token := scanner.Text()

		if isLetter(token) {
			token = strings.ToLower(token)
			if val, ok := charFound["Letters"][token]; ok {
				*val++
			} else {
				t := 1
				charFound["Letters"][token] = &t
			}
		} else if isNumber(token) {
			if val, ok := charFound["Numbers"][token]; ok {
				*val++

			} else {
				t := 1
				charFound["Numbers"][token] = &t
			}
		} else {
			continue
		}

	}

	checkMissingLetters(charFound)
	checkMissingNumbers(charFound)

	return charFound
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

func checkMissingLetters(cf map[string]map[string]*int) {
	for l := 'a'; l < 'z'; l++ {
		if _, ok := cf["Letters"][string(l)]; !ok {
			t := 0
			cf["LettersNotFound"][string(l)] = &t
		}
	}
}

func checkMissingNumbers(cf map[string]map[string]*int) {
	for n := 0; n < 10; n++ {
		if _, ok := cf["Numbers"][strconv.Itoa(n)]; !ok {
			t := 0
			cf["NumbersNotFound"][strconv.Itoa(n)] = &t
		}
	}
}
