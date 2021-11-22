package main

import (
	"errors"
	"log"
	"net/http"
	"os"
	"path"
)

type FileServer struct {
	baseFolder string
}

func (f FileServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(405)
		return
	}

	targetPath := path.Join(f.baseFolder, r.URL.Path)

	// TODO Check for ..

	file, err := os.ReadFile(targetPath)

	if err != nil {
		log.Printf("Met %e while reading file", err)

		if errors.Is(err, os.ErrNotExist) {
			w.WriteHeader(404)
			return
		}

		w.WriteHeader(503)
		return
	}

	w.Write(file)
	return
}
