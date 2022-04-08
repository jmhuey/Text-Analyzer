package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

		//is_alphanumeric := regexp.MustCompile(`^[a-zA-Z]*$`).MatchString(token)

		switch []rune(token)[0] {
		case ' ', '\t', '\n', '\f', '\r', '\v':
			continue

		default:
			//if !is_alphanumeric {
			//	continue
			//}
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
