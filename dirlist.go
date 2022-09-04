package dirlist

import (
	"github.com/ayushg3112/dirlist/internal/http"
	"github.com/ayushg3112/dirlist/sort"
	"github.com/ayushg3112/dirlist/walk"
)

func StartServer(options ProcessingOptions) error {
	sorter, err := sort.NewDirEntrySorter(options.SortField, options.SortOrder)

	if err != nil {
		return err
	}

	walker, err := walk.NewWalker(options.RootDirAbsPath)

	if err != nil {
		return err
	}

	err = http.StartSinglePageServer(walker, sorter, http.ServerOptions{
		Port:           options.HTTPPort,
		RootDirAbsPath: options.RootDirAbsPath,
		CachedMode:     options.CachedMode,
	})

	return err
}
