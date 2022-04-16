package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type TextAnalysis struct {
	CharInfo CharInfo
}

func main() {
	fmt.Println("Listening to port 8080")
	http.HandleFunc("/", parseFile)
	http.ListenAndServe(":8080", nil)
}

// Creates a new file in the directory with the contents from the POST request before analyzing it
func parseFile(w http.ResponseWriter, r *http.Request) {

	//----------------------------------------------------
	// Preparing file from POST request for text analysis
	//----------------------------------------------------
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

	//---------------------
	// Text analysis calls
	//---------------------
	data := scanFileChars(file, handler.Filename)

	word := scanFileWord(file)

	for key, value := range word {
		fmt.Println(key, ":", value)
	}

	//-----------------------------------------
	// Create and send response back to client
	//-----------------------------------------
	textAnalysis := new(TextAnalysis)
	textAnalysis.CharInfo = *data

	resp, err := json.Marshal(textAnalysis)
	if err != nil {
		fmt.Print("Unable to marshal information provided")
	}

	w.Write(resp)
}
