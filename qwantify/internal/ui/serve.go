package ui

import (
	"embed"
	"fmt"
	"io"
	"io/fs"
	"log"
	"mime"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

//go:embed static/*
var content embed.FS

func NewClient() (s *http.ServeMux) {

	fsys := fs.FS(content)
	html, err := fs.Sub(fsys, "static")
	if err != nil {
		log.Fatal("failed to get ui fs", err)
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		//must be a get method (browser)
		if r.Method != "GET" {
			http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
			return
		}

		path := filepath.Clean(r.URL.Path)

		path = strings.TrimPrefix(path, `\`)

		if path == `/` || path == "" { // Add other paths that you route on the UI side here
			path = "index.html"
		}

		file, err := html.Open(path)
		if err != nil {
			if os.IsNotExist(err) {
				log.Println("file", path, "not found:", err)
				http.NotFound(w, r)
				return
			}
			log.Println("file", path, "cannot be read:", err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		contentType := mime.TypeByExtension(filepath.Ext(path))
		w.Header().Set("Content-Type", contentType)

		// if strings.HasPrefix(path, "static/") {
		w.Header().Set("Cache-Control", "public, max-age=31536000")
		// }
		stat, err := file.Stat()
		if err == nil && stat.Size() > 0 {
			w.Header().Set("Content-Length", fmt.Sprintf("%d", stat.Size()))
		}

		n, _ := io.Copy(w, file)

		log.Println("file", path, "copied", n, "bytes")
	})

	return mux
}
