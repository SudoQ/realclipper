package main

import (
	cb "github.com/atotto/clipboard"
	"log"
	"net/http"
	"io/ioutil"
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
	addr := fmt.Sprintf("http://%s:%s/realclipper/api/clipboard", host, port)

	// Get content
	resp, err := http.Get(addr)
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
	content := string(b)

	// Write to clipboard
	err = cb.WriteAll(content)
	if err != nil {
		log.Println(err)
		return
	}
}
