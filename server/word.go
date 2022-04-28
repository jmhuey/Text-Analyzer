package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

type WordInfo struct {
	TotalWords int
	Words      map[string]*int
}

// scans the file and puts unique values into a map while recording the number of times a repeated value appears
func scanFileWord(s string) *WordInfo {
	//-------------------------------------
	// Set up the struct to hold file data
	//------------------------------------
	wordInfo := new(WordInfo)

	wordInfo.Words = make(map[string]*int)

	//------------------------------------------
	// Process text and save values into struct
	//------------------------------------------
	file, err := os.Open(s)
	if err != nil {
		fmt.Println(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		token := scanner.Text()

		rstr, err := regexp.Compile(`[^\w]`)
		if err != nil {
			fmt.Print("Error: Unable to parse regular expression")
		}

		token = rstr.ReplaceAllString(token, "")

		if val, ok := wordInfo.Words[token]; ok {
			*val++
		} else {
			t := 1
			wordInfo.Words[token] = &t
		}
		wordInfo.TotalWords++
	}

	return wordInfo
}
