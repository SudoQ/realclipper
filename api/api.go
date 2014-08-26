package main

import (
	"net/http"
	"fmt"
	"log"
	"bytes"
	"github.com/go-martini/martini"
	"os"
)


var clipboard string

func ClipboardGroup(r martini.Router) {
	r.Get("", CopyClipboard)
	r.Post("", PasteClipboard)
}

func CopyClipboard()(int, string) {
	log.Printf("Copied %s\n", clipboard)
	return 200, clipboard
}

func PasteClipboard(w http.ResponseWriter, r *http.Request) {
	buf:= new(bytes.Buffer)
	buf.ReadFrom(r.Body)
	s := buf.String()
	clipboard = s
	log.Printf("Pasted %s\n", clipboard)
	fmt.Fprintf(w, "ok")
}

func main(){
	m := martini.Classic()
	//http.HandleFunc("/realclipper/api/clipboard", handler)
	m.Group("/realclipper/api/clipboard", ClipboardGroup)
	env_port := os.Getenv("REALCLIPPER_API_PORT")
	if env_port == "" {
		env_port = "8080"
	}
	port := fmt.Sprintf(":%s", env_port)
	log.Fatal(http.ListenAndServe(port, m))
}
