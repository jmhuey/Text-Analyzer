package main

import (
	"bufio"
	"fmt"
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

		/*
			// Temporarily added to focus on letters only -----------------------
			if !isLetter(token) && !isNumber(token) {
				continue
			}
			// ------------------------------------------------------------------
		*/
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

/*
// Sorts the characters into a slice of letters and numbers before further sorting alphanumerically
func sortCharacters(m map[string]*int) ([]string, []int) {

	letters_found := make([]string, 0)
	num_found := make([]int, 0)

	// sorting the characters found
	for k := range m {
		if isLetter(k) {
			letters_found = append(letters_found, k)
		} else if isNumber(k) {

			knum, err := strconv.Atoi(k)
			if err != nil {
				fmt.Print("Error: Unable to convert string num to int")
			}

			num_found = append(num_found, knum)
		}
	}

	sort.Strings(letters_found)
	sort.Ints(num_found)

	return letters_found, num_found
}
*/

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

//-----------------
// Print Functions
//-----------------
func printLettersFound(lf []string, m map[string]*int) {
	fmt.Println()
	fmt.Println("These are the letters found in the text:")

	for _, letter := range lf {
		fmt.Println(letter, ": ", *m[letter])
	}
}

func printLettersNotFound(lnf []string) {
	if len(lnf) != 0 {
		fmt.Println()
		fmt.Println("These are the letters not found in the text:")

		for _, letter := range lnf {
			fmt.Println(letter)
		}
	}
}

func printNumbersFound(nf []int, m map[string]*int) {
	fmt.Println()
	fmt.Println("These are the numbers found in the text:")

	for _, num := range nf {
		fmt.Println(num, ": ", *m[strconv.Itoa(num)])
	}
}

func printNumbersNotFound(nnf []int) {
	if len(nnf) != 0 {
		fmt.Println()
		fmt.Println("These are the numbers not found in the text:")

		for _, num := range nnf {
			fmt.Println(num)
		}
	}

}
