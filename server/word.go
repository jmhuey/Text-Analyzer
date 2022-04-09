package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

// scans the file and puts unique values into a map while recording the number of times a repeated value appears
func scanFileWord(f *os.File) map[string]*int {
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)

	word := make(map[string]*int)

	for scanner.Scan() {
		token := scanner.Text()

		rstr, err := regexp.Compile(`[^\w]`)
		if err != nil {
			fmt.Print("Error: Unable to parse regular expression")
		}

		token = rstr.ReplaceAllString(token, "")

		if val, ok := word[token]; ok {
			*val++
		} else {
			t := 1
			word[token] = &t
		}
	}

	return word
}
