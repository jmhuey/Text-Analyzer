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
	Filename string
	CharInfo CharInfo
	WordInfo WordInfo
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
	if err != nil {
		log.Fatal(err)
	}

	io.Copy(file, f)

	//---------------------
	// Text analysis calls
	//---------------------
	charData := scanFileChars(handler.Filename)
	wordData := scanFileWord(handler.Filename)

	//-----------------------------------------
	// Create and send response back to client
	//-----------------------------------------
	textAnalysis := new(TextAnalysis)
	textAnalysis.Filename = handler.Filename
	textAnalysis.CharInfo = *charData
	textAnalysis.WordInfo = *wordData

	resp, err := json.Marshal(textAnalysis)
	if err != nil {
		fmt.Print("Unable to marshal information provided")
	}

	w.Write(resp)
}
