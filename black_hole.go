package main

import (
	"encoding/base64"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func Output(w http.ResponseWriter, r *http.Request) {
	base_path, ok := os.LookupEnv("BH_PATH")
	if !ok {
		log.Println("Env BH_PATH was not found...")
		log.Println("Creating default /tmp/bh/ ...")
		os.Mkdir("/tmp/bh/", 0755)
		base_path = "/tmp/bh/"
	}
	tempfile := base_path + GenFileName()
	file, err := SaveFile(tempfile)
	if err != nil {
		log.Fatalf("%+v\n", err)
	}
	fmt.Printf("%s", file.Name())
	io.WriteString(w, "Saved file: "+file.Name()+"\n")

}

func GenFileName() string {
	now := time.Now().String()
	filename := base64.StdEncoding.EncodeToString([]byte(now))
	return filename
}

func SaveFile(tempfile string) (*os.File, error) {
	// accept my bytes

	f, err := os.OpenFile(tempfile, os.O_RDWR|os.O_CREATE, 0755)
	defer f.Close()
	if err != nil {
		log.Fatalf("%+v\n", err)
	}
	// take you bytes here
	s := []byte("Kitzen the CAT!")
	f.Write(s)
	return f, nil
}

func main() {
	http.HandleFunc("/", Output)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
