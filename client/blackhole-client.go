package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gabriel-vasile/mimetype"
)

func main() {
	URL := "http://localhost:8080"
	log.Printf("Sending to %s", URL)
	// open file
	file := os.Args[1]
	mime, err := mimetype.DetectFile(file)
	// read data to buffer
	buf, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatalf("%+v\n", err)
	}
	b := bytes.NewReader(buf)
	// send data
	resp, err := http.Post(URL, mime.String(), b)
	defer resp.Body.Close()
	if err != nil {
		log.Fatalf("%+v\n", err)
	} else {
		log.Printf("Data was successfuly send: %v", resp.Status)
	}

}
