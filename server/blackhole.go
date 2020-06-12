package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

func GenFileName(fileName string) string {
	now := fileName + time.Now().Format("20060102150405")
	filename := base64.StdEncoding.EncodeToString([]byte(now))
	return filename
}

func SaveFile(tempfile string, b []byte) (*os.File, error) {
	f, err := os.OpenFile(tempfile, os.O_RDWR|os.O_CREATE, 0755)
	defer f.Close()
	if err != nil {
		log.Fatalf("%+v\n", err)
	}

	f.Write(b)
	return f, nil
}

func uploadFormFile(w http.ResponseWriter, r *http.Request) {
	basePath, ok := os.LookupEnv("BH_PATH")
	if !ok {
		log.Println("Env BH_PATH was not found...")
		log.Println("Creating default /tmp/bh/ ...")
		os.Mkdir("/tmp/bh/", 0755)
		basePath = "/tmp/bh/"
	}
	r.ParseMultipartForm(1024 * 1024)
	uploadingFile, handler, err := r.FormFile("file")
	uploadingBytes, err := ioutil.ReadAll(uploadingFile)
	if err != nil {
		log.Fatalf("%+v\n", err)
		return
	}
	log.Printf("File Name:%s Size:%d Header:%s\n", handler.Filename, handler.Size, handler.Header)
	tempfile := basePath + GenFileName(handler.Filename)
	file, err := SaveFile(tempfile, uploadingBytes)
	if err != nil {
		log.Fatalf("%+v\n", err)
	}

	fmt.Printf("%s", file.Name())
}

func uploadFile(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	basePath, ok := os.LookupEnv("BH_PATH")
	if !ok {
		log.Println("Env BH_PATH was not found...")
		log.Println("Creating default /tmp/bh/ ...")
		os.Mkdir("/tmp/bh/", 0755)
		basePath = "/tmp/bh/"
	}

	tempfile := basePath + GenFileName("")
	file, err := SaveFile(tempfile, b)
	if err != nil {
		log.Fatalf("%+v\n", err)
	}

	fmt.Printf("%s\n", file.Name())
}

func main() {
	http.HandleFunc("/", uploadFile)
	http.HandleFunc("/uf", uploadFormFile)
	log.Fatal(http.ListenAndServe(":8080", nil))
}