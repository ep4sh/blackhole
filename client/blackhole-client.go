package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gabriel-vasile/mimetype"
	"gopkg.in/yaml.v2"
)

type URL struct {
	Host string `yaml:"host"`
}

func main() {
	var url URL
	configPath, _ := os.LookupEnv("HOME")
	configPath = configPath + "/" + ".blackhole.yml"
	yamlConfig, err := ioutil.ReadFile(configPath)
	err = yaml.Unmarshal(yamlConfig, &url)
	if err != nil {
		log.Printf("Cannot parse config file %s", configPath)
		return
	}
	// check if user input file to upload
	if len(os.Args) < 2 {
		log.Fatalf("Input file to upload")
		return
	}

	log.Printf("Sending to %s", url.Host)
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
	resp, err := http.Post(url.Host, mime.String(), b)
	if err != nil {
		log.Fatalf("%+v\n", err)
	}

	defer resp.Body.Close()
	rb, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("%+v\n", err)
	} else {
		log.Printf("Data was successfuly send: %v", resp.Status)
		log.Printf("your link ==>  %s", url.Host+"?f="+string(rb))
	}

}
