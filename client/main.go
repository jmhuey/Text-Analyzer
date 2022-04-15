package main

import (
	"bufio"
	"bytes"
	"encoding/json"
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

type CharsFound struct {
	Filename        string
	Letters         map[string]*int
	Numbers         map[string]*int
	LettersNotFound map[string]*int
	NumbersNotFound map[string]*int
}

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

		fmt.Println()
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

	//fmt.Print(content)

	parseJson(content)

	//fmt.Print(content)
}

//-----------------
// Print Functions
//-----------------

func parseJson(c []byte) {

	//fileData := make(map[string]map[string]int)

	fileData := new(CharsFound)

	json.Unmarshal(c, &fileData)

	fmt.Println("Here is the data about the file:", fileData.Filename)

	printLettersFound(fileData)
	printLettersNotFound(fileData)
	printNumbersFound(fileData)
	printNumbersNotFound(fileData)

	fmt.Println("------------------------------------------------------")
	fmt.Println()

}

func printLettersFound(fd *CharsFound) {
	fmt.Println()
	fmt.Println("These are the letters found in the text: ")

	for l := 'a'; l <= 'z'; l++ {
		if _, ok := fd.Letters[string(l)]; ok {
			fmt.Printf("%c: %d | ", l, *fd.Letters[string(l)])
		}
	}
	fmt.Println()
}

func printLettersNotFound(fd *CharsFound) {
	if len(fd.LettersNotFound) != 0 {
		fmt.Println()
		fmt.Println("These are the letters not found in the text: ")
		for l := 'a'; l <= 'z'; l++ {
			if _, ok := fd.LettersNotFound[string(l)]; ok {
				fmt.Printf("%c: %d | ", l, *fd.LettersNotFound[string(l)])
			}
		}
		fmt.Println()
	}
}

func printNumbersFound(fd *CharsFound) {
	fmt.Println()
	fmt.Println("These are the numbers found in the text:")

	for n := 0; n < 10; n++ {
		if _, ok := fd.Numbers[strconv.Itoa(n)]; ok {
			fmt.Printf("%d: %d | ", n, *fd.Numbers[strconv.Itoa(n)])
		}
	}
	fmt.Println()
}

func printNumbersNotFound(fd *CharsFound) {
	if len(fd.NumbersNotFound) != 0 {
		fmt.Println()
		fmt.Println("These are the numbers not found in the text: ")
		for n := 0; n < 10; n++ {
			if _, ok := fd.NumbersNotFound[strconv.Itoa(n)]; ok {
				fmt.Printf("%d: %d | ", n, *fd.NumbersNotFound[strconv.Itoa(n)])
			}
		}
		fmt.Println()
	}

}
