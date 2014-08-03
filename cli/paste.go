package main

import (
	cb "github.com/atotto/clipboard"
	"log"
	"net/http"
	"io/ioutil"
	"bytes"
)

func main() {
	// Read the content of the clipboard
	content, err := cb.ReadAll()

	// Convert to io.Reader
	bufptr := bytes.NewBufferString(content)

	// Send a POST request
	resp, err := http.Post("http://localhost:8080/realclipper/api/v0.1/clipboard", "text/plain", bufptr)
	if err != nil {
		log.Println(err)
		return
	}
	defer resp.Body.Close()

	// Read response
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return
	}

	// Log status
	status := string(b)
	log.Println(status)
}
