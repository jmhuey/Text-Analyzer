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

// Holds information returned by the server
type CharsFound struct {
	Filename        string
	Letters         map[string]*int
	Numbers         map[string]*int
	LettersNotFound map[string]*int
	NumbersNotFound map[string]*int
}

func main() {

	// Read in user input
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter the file or files containing the text you want analyzed or enter exit to exit the program: ")

	var wg sync.WaitGroup

	// Reads in user input until EXIT is passed
	for scanner.Scan() {

		input := scanner.Text()

		if input == "exit" {
			os.Exit(0)
		}

		// Check all user input to see if the files exist in current directory before calling the sendRequest function in a goroutine
		for _, fileName := range strings.Fields(input) {

			_, err := os.Stat(fileName)
			if err != nil {
				log.Println(fileName, "was not found in current directory")
				continue
			}

			wg.Add(1)

			go func(f string) {
				defer wg.Done()
				sendRequest(f)
			}(fileName)

		}

		// Waits for all files to finish being analyzed before asking the user for input
		wg.Wait()

		fmt.Print("\nEnter the file or files containing the text you want analyzed or enter exit to exit the program: ")
	}

}

// Sends a POST request to the local server containing the file to be analyzed
func sendRequest(s string) {

	// Ensure that the client can continue reading input even if error occurs during the request process
	defer func() {
		if r := recover(); r != nil {
			log.Println("Unable to send POST request for file:", s)
		}
	}()

	//--------------------------------------------------
	// Prepare the file to be sent through POST request
	//--------------------------------------------------
	client := &http.Client{}
	body := &bytes.Buffer{}

	writer := multipart.NewWriter(body)

	fw, err := writer.CreateFormFile("file", s)
	if err != nil {
		log.Println(err)
	}

	file, err := os.Open(s)
	if err != nil {
		log.Println(err)
	}

	_, err = io.Copy(fw, file)
	if err != nil {
		log.Println(err)
	}

	writer.Close()

	//----------------------------------------------------
	// Create, setup and send POST request with file info
	//----------------------------------------------------
	req, err := http.NewRequest("POST", "http://127.0.0.1:8080/", body)
	if err != nil {
		log.Println(err)
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())

	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
	}

	//-----------------------------------------------------
	// Receive and parse the data received from the server
	//-----------------------------------------------------
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}

	parseJson(content)
}

// Takes in content returned by the server and prints it out
func parseJson(c []byte) {

	fileData := new(CharsFound)

	json.Unmarshal(c, &fileData)

	fmt.Println("\n------------------------------------------------------")
	fmt.Println("Here is the data about the file:", fileData.Filename)

	fileData.printLettersFound()
	fileData.printLettersNotFound()
	fileData.printNumbersFound()
	fileData.printNumbersNotFound()

	fmt.Println("------------------------------------------------------")
}

//-----------------
// Print Functions
//-----------------

func (fd *CharsFound) printLettersFound() {
	fmt.Println("\nThese are the letters found in the text: ")

	for l := 'a'; l <= 'z'; l++ {
		if _, ok := fd.Letters[string(l)]; ok {
			fmt.Printf("%c: %d | ", l, *fd.Letters[string(l)])
		}
	}
	fmt.Println()
}

func (fd *CharsFound) printLettersNotFound() {
	if len(fd.LettersNotFound) != 0 {
		fmt.Println("\nThese are the letters not found in the text: ")
		for l := 'a'; l <= 'z'; l++ {
			if _, ok := fd.LettersNotFound[string(l)]; ok {
				fmt.Printf("%c: %d | ", l, *fd.LettersNotFound[string(l)])
			}
		}
		fmt.Println()
	}
}

func (fd *CharsFound) printNumbersFound() {
	fmt.Println("\nThese are the numbers found in the text:")

	for n := 0; n < 10; n++ {
		if _, ok := fd.Numbers[strconv.Itoa(n)]; ok {
			fmt.Printf("%d: %d | ", n, *fd.Numbers[strconv.Itoa(n)])
		}
	}
	fmt.Println()
}

func (fd *CharsFound) printNumbersNotFound() {
	if len(fd.NumbersNotFound) != 0 {
		fmt.Println("\nThese are the numbers not found in the text: ")
		for n := 0; n < 10; n++ {
			if _, ok := fd.NumbersNotFound[strconv.Itoa(n)]; ok {
				fmt.Printf("%d: %d | ", n, *fd.NumbersNotFound[strconv.Itoa(n)])
			}
		}
		fmt.Println()
	}

}
