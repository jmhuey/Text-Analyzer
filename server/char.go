package main

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// Holds the all the info found on the text related to characters
type CharInfo struct {
	Filename        string
	Letters         map[string]*int
	Numbers         map[string]*int
	LettersNotFound map[string]*int
	NumbersNotFound map[string]*int
}

// scans the file and puts unique values into a map while recording the number of times a repeated value appears
func scanFileChars(f *os.File, s string) *CharInfo {

	//-------------------------------------
	// Set up the struct to hold file data
	//------------------------------------
	charInfo := new(CharInfo)
	charInfo.Filename = s

	charInfo.Letters = make(map[string]*int)
	charInfo.Numbers = make(map[string]*int)
	charInfo.LettersNotFound = make(map[string]*int)
	charInfo.NumbersNotFound = make(map[string]*int)

	//------------------------------------------
	// Process text and save values into struct
	//------------------------------------------
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanRunes)

	for scanner.Scan() {
		token := scanner.Text()

		if isLetter(token) {
			token = strings.ToLower(token)
			if val, ok := charInfo.Letters[token]; ok {
				*val++
			} else {
				t := 1
				charInfo.Letters[token] = &t
			}
		} else if isNumber(token) {
			if val, ok := charInfo.Numbers[token]; ok {
				*val++

			} else {
				t := 1
				charInfo.Numbers[token] = &t
			}
		} else {
			continue
		}

	}

	charInfo.checkMissingAlphanumeric()

	// Return the struct containing info found
	return charInfo
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

// checks if there are any missing alphanumeric characters
func (cf *CharInfo) checkMissingAlphanumeric() {
	for l := 'a'; l < 'z'; l++ {
		if _, ok := cf.Letters[string(l)]; !ok {
			t := 0
			cf.LettersNotFound[string(l)] = &t
		}
	}

	for n := 0; n < 10; n++ {
		if _, ok := cf.Numbers[strconv.Itoa(n)]; !ok {
			t := 0
			cf.NumbersNotFound[strconv.Itoa(n)] = &t
		}
	}
}
