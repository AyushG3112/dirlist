package http

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ayushg3112/dirlist/internal/templates"
	"github.com/ayushg3112/dirlist/walk"
)

type ServerOptions struct {
	Port string
}

func StartSinglePageServer(structure []walk.DirectoryStructure, options ServerOptions) error {
	html, err := templates.GenerateSinglePageTemplateHTML(structure)

	if err != nil {
		return err
	}

	log.Printf("starting the server at port %s", options.Port)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, html)
	})

	return http.ListenAndServe(":"+options.Port, nil)
}
