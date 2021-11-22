package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	var port = 8080
	var base_folder = "/home/keddad/"

	fmt.Printf("Starting goweb at %d", port)

	if err := http.ListenAndServe(":8080", FileServer{base_folder}); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Shutting down goweb")
}
