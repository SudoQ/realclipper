package main

import (
	"net/http"
	"fmt"
	"log"
	"bytes"
)


var clipboard string

func handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
		case "GET":
			log.Printf("Copied %s\n", clipboard)
			fmt.Fprintf(w, clipboard)
		case "POST":
			buf:= new(bytes.Buffer)
			buf.ReadFrom(r.Body)
			s := buf.String()
			clipboard = s
			log.Printf("Pasted %s\n", clipboard)
			fmt.Fprintf(w, "ok")
		default:
			log.Println("Error")
	}
}

func main(){
	http.HandleFunc("/realclipper/api/v0.1/clipboard", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
