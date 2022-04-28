package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// Holds the all the info found on the text related to characters
type CharInfo struct {
	TotalChar       int
	Letters         map[string]*int
	Numbers         map[string]*int
	NonAlphaNumeric map[string]*int
	LettersNotFound map[string]*int
	NumbersNotFound map[string]*int
}

// scans the file and puts unique values into a map while recording the number of times a repeated value appears
func scanFileChars(s string) *CharInfo {

	//-------------------------------------
	// Set up the struct to hold file data
	//------------------------------------
	charInfo := new(CharInfo)

	charInfo.Letters = make(map[string]*int)
	charInfo.Numbers = make(map[string]*int)
	charInfo.NonAlphaNumeric = make(map[string]*int)
	charInfo.LettersNotFound = make(map[string]*int)
	charInfo.NumbersNotFound = make(map[string]*int)

	//------------------------------------------
	// Process text and save values into struct
	//------------------------------------------

	file, err := os.Open(s)
	if err != nil {
		fmt.Println(err)
	}

	scanner := bufio.NewScanner(file)
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
			charInfo.TotalChar++
		} else if isNumber(token) {
			if val, ok := charInfo.Numbers[token]; ok {
				*val++

			} else {
				t := 1
				charInfo.Numbers[token] = &t
			}
			charInfo.TotalChar++
		} else {
			switch token {
			case "!", "?", ".":
				if val, ok := charInfo.NonAlphaNumeric[token]; ok {
					*val++
				} else {
					t := 1
					charInfo.NonAlphaNumeric[token] = &t
				}
				charInfo.TotalChar++
			default:
				continue

			}
		}

	}

	charInfo.checkMissingAlphanumeric()

	// Return the struct containing info found
	return charInfo
}

// checks if string passed is a letter
func isLetter(s string) bool {
	isLetter := regexp.MustCompile(`^[a-zA-Z]*$`).MatchString(s)
	return isLetter
}

// checks if string passed is a number
func isNumber(s string) bool {
	isNumber := regexp.MustCompile(`^[0-9]*$`).MatchString(s)
	return isNumber
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
