package main

import (
	"fmt"
	"github.com/go-martini/martini"
	"log"
	"net/http"
	"os"
)

func main() {
	m := martini.Classic()
	m.Use(martini.Static("assets"))
	env_port := os.Getenv("REALCLIPPER_WEB_PORT")
	if env_port == "" {
		env_port = "8080"
	}
	port := fmt.Sprintf(":%s", env_port)
	log.Fatal(http.ListenAndServe(port, m))
}
