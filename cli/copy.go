package main

import (
	cb "github.com/atotto/clipboard"
	"log"
	"net/http"
	"io/ioutil"
)

func main() {
	// Get content
	resp, err := http.Get("http://localhost:8080/realclipper/api/v0.1/clipboard")
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
