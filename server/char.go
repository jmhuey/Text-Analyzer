package main

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type CharsFound struct {
	Filename        string
	Letters         map[string]*int
	Numbers         map[string]*int
	LettersNotFound map[string]*int
	NumbersNotFound map[string]*int
}

// scans the file and puts unique values into a map while recording the number of times a repeated value appears
func scanFileChar(f *os.File, s string) *CharsFound {
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanRunes)

	//char := make(map[string]*int)

	charsFound := new(CharsFound)
	charsFound.Filename = s

	charsFound.Letters = make(map[string]*int)
	charsFound.Numbers = make(map[string]*int)
	charsFound.LettersNotFound = make(map[string]*int)
	charsFound.NumbersNotFound = make(map[string]*int)

	for scanner.Scan() {
		token := scanner.Text()

		if isLetter(token) {
			token = strings.ToLower(token)
			if val, ok := charsFound.Letters[token]; ok {
				*val++
			} else {
				t := 1
				charsFound.Letters[token] = &t
			}
		} else if isNumber(token) {
			if val, ok := charsFound.Numbers[token]; ok {
				*val++

			} else {
				t := 1
				charsFound.Numbers[token] = &t
			}
		} else {
			continue
		}

	}

	checkMissingLetters(charsFound)
	checkMissingNumbers(charsFound)

	return charsFound
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

func checkMissingLetters(cf *CharsFound) {
	for l := 'a'; l < 'z'; l++ {
		if _, ok := cf.Letters[string(l)]; !ok {
			t := 0
			cf.LettersNotFound[string(l)] = &t
		}
	}
}

func checkMissingNumbers(cf *CharsFound) {
	for n := 0; n < 10; n++ {
		if _, ok := cf.Numbers[strconv.Itoa(n)]; !ok {
			t := 0
			cf.NumbersNotFound[strconv.Itoa(n)] = &t
		}
	}
}
