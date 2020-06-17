package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
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

func createBasepath() string {
	basePath, ok := os.LookupEnv("BH_PATH")
	if !ok {
		log.Println("Env BH_PATH was not found...")
		log.Println("Creating default /tmp/bh/ ...")
		os.Mkdir("/tmp/bh/", 0755)
		basePath = "/tmp/bh/"
	}
	return basePath
}

func uploadFile(basePath string, w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		keys := r.URL.Query()
		log.Printf("keys: %v", keys)
		// filename := basePath + r.
		//http.ServeFile(w, r, "???")
	case "POST":
		b, err := ioutil.ReadAll(r.Body)

		if err != nil {
			log.Fatalf("%+v\n", err)
		}

		tempfile := basePath + GenFileName("")
		file, err := SaveFile(tempfile, b)
		if err != nil {
			log.Fatalf("%+v\n", err)
		}
		resp := []byte(strings.Split(file.Name(), basePath)[1])
		w.Write(resp)
		fmt.Printf("%s\n", resp)
		fmt.Printf("%s\n", file.Name())
	default:
		fmt.Fprintf(w, "Only GET and POST methods are supported.")
	}

}

func main() {
	basePath := createBasepath()
	http.HandleFunc("/b", func(w http.ResponseWriter, r *http.Request) {
		uploadFile(basePath, w, r)
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
