package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", Output)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
