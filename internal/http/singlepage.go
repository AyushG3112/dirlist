package http

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"

	"github.com/ayushg3112/dirlist/internal/templates"
	"github.com/ayushg3112/dirlist/walk"
)

type ServerOptions struct {
	Port           string
	RootDirAbsPath string
}

func StartSinglePageServer(structure []walk.DirectoryStructure, options ServerOptions) error {
	html, err := templates.GenerateSinglePageTemplateHTML(structure)

	if err != nil {
		return err
	}

	log.Printf("starting the server at port %s", options.Port)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path

		if path == "" || path == "/" {
			fmt.Fprint(w, html)
		}

		filePath := filepath.Join(options.RootDirAbsPath, path)

		http.ServeFile(w, r, filePath)
	})

	return http.ListenAndServe(":"+options.Port, nil)
}
