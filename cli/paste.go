package main

import (
	cb "github.com/atotto/clipboard"
	"log"
	"net/http"
	"io/ioutil"
	"bytes"
	"os"
	"fmt"
)

func main() {
	// Read enviromental variables
	host := os.Getenv("REALCLIPPER_API_HOST")
	if host == "" {
		host = "localhost"
	}
	port := os.Getenv("REALCLIPPER_API_PORT")
	if port == "" {
		port = "8080"
	}

	// Construct address
	addr := fmt.Sprintf("http://%s:%s/realclipper/api/v0.1/clipboard", host, port)

	// Read the content of the clipboard
	content, err := cb.ReadAll()

	// Convert to io.Reader
	bufptr := bytes.NewBufferString(content)

	// Send a POST request
	resp, err := http.Post(addr, "text/plain", bufptr)
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
