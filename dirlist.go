package dirlist

import (
	"github.com/ayushg3112/dirlist/internal/http"
	"github.com/ayushg3112/dirlist/sort"
	"github.com/ayushg3112/dirlist/walk"
)

func StartServer(options ProcessingOptions) error {
	sorter, err := sort.NewSorter(options.SortField, options.SortOrder)

	if err != nil {
		return err
	}

	structure, err := walk.Walk(options.RootDirAbsPath, sorter)

	if err != nil {
		return err
	}

	err = http.StartSinglePageServer(structure, http.ServerOptions{
		Port:           options.HTTPPort,
		RootDirAbsPath: options.RootDirAbsPath,
	})

	return err
}
