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
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter the file containing the text you want analyzed: ")

	for scanner.Scan() {

		fileName := scanner.Text()

		_, err := os.Stat(fileName)
		if err != nil {
			fmt.Println("Please enter a valid file containing the text you want analyzed: ")

			continue
		}

		sendRequest(fileName)
		fmt.Print("Enter the file containing the text you want analyzed: ")
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

	file, err := os.Open("test.txt")
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

	fmt.Println(content)
}
