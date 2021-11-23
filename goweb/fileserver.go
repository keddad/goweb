package main

import (
	"bytes"
	"compress/gzip"
	"errors"
	"log"
	"net/http"
	"os"
	"path"
	"strings"
)

type FileServer struct {
	baseFolder string
	gzip       bool
}

func (f FileServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(405)
		return
	}

	targetPath := path.Join(f.baseFolder, r.URL.Path)

	if !path.IsAbs(targetPath) { // Actually it looks like path.Join somehow kills .. related exploits
		w.WriteHeader(403)
		return
	}

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

	if f.gzip && strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
		var buf bytes.Buffer
		zw := gzip.NewWriter(&buf)
		defer zw.Close()

		_, err := zw.Write(file)
		if err != nil {
			log.Fatal(err)
		}

		if err := zw.Close(); err != nil {
			log.Fatal(err)
		}

		w.Header().Add("Content-Encoding", "gzip") // Should go before w.Write, or add w.WriteHeader
		w.Write(buf.Bytes())
	} else {
		w.Write(file)
	}
	return
}
