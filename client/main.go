package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter the file containing the text you want analyzed: ")

	var wg sync.WaitGroup

	for scanner.Scan() {

		input := scanner.Text()

		for _, fileName := range strings.Fields(input) {
			_, err := os.Stat(fileName)
			if err != nil {
				fmt.Println(fileName, "is not a valid file")

				continue
			}

			wg.Add(1)

			go func(f string) {
				defer wg.Done()
				sendRequest(f)
			}(fileName)

		}

		wg.Wait()

		fmt.Print("Enter the file or files containing the text you want analyzed: ")
	}

}

func sendRequest(s string) {

	client := &http.Client{}

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	fw, err := writer.CreateFormFile("file", s)
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.Open(s)
	if err != nil {
		log.Fatal(err)
	}

	_, err = io.Copy(fw, file)
	if err != nil {
		log.Fatal(err)
	}

	writer.Close()

	req, err := http.NewRequest("POST", "http://127.0.0.1:8080/", body)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(content))
	fmt.Println()
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
