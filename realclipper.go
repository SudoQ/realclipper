package main

import (
	"net/http"
	"fmt"
	"log"
	"bytes"
)


var clipboard string

func handler(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "Hi there I love %s!", r.URL.Path[1:])
	switch r.Method {
		case "GET":
			log.Println("COPY")
			fmt.Fprintf(w, "copy")
		case "POST":
			log.Println("PASTE")
			buf:= new(bytes.Buffer)
			buf.ReadFrom(r.Body)
			s := buf.String()
			fmt.Fprintf(w, s)
		default:
			log.Println("Error")
	}
}

func main(){
	http.HandleFunc("/realclipper/api/v1.0/clipboard", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
