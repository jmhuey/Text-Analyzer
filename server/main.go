package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	fmt.Println("Listening to port 8080")
	http.HandleFunc("/", parseFile)
	http.ListenAndServe(":8080", nil)

}

func parseFile(w http.ResponseWriter, r *http.Request) {

	r.ParseMultipartForm(32 << 20)

	f, handler, err := r.FormFile("file")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	file, err := os.Create(handler.Filename)

	io.Copy(file, f)

	// open the file
	file, err = os.Open(handler.Filename)
	if err != nil {
		fmt.Println(err)
	}

	//---------------------------
	// Character operation calls
	//---------------------------

	charFound := scanFileChar(file)

	resp, err := json.Marshal(charFound)
	if err != nil {
		fmt.Print("bleh")
	}

	w.Write(resp)

	word := scanFileWord(file)

	for key, value := range word {
		fmt.Println(key, ":", value)
	}
}
