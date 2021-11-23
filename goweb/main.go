package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {
	var port int
	var base_folder string
	var gzip bool

	flag.StringVar(&base_folder, "f", "", "Folder to serve files from")
	flag.IntVar(&port, "p", 0, "Port to use")
	flag.BoolVar(&gzip, "g", false, "Whether to use gzip, if client supports it")
	flag.Parse()

	if base_folder == "" {
		log.Fatal("Empty base folder!")
	}

	if port <= 0 || port > 65536 {
		log.Fatal("Invalid port, or no port specified")
	}

	fmt.Printf("Starting goweb at %d", port)

	if err := http.ListenAndServe(":8080", FileServer{base_folder, gzip}); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Shutting down goweb")
}
