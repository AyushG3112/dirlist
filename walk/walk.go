package walk

import (
	"os"
	"path/filepath"
	"time"

	"github.com/ayushg3112/dirlist/sort"
)

type FileInfo struct {
	ModifiedAt time.Time
	Size       int64
}

type DirectoryStructure struct {
	Level    int
	Name     string
	Path     string
	IsDir    bool
	FileInfo FileInfo
	Children []DirectoryStructure
}

func Walk(dir string, sorter sort.Sorter) ([]DirectoryStructure, error) {
	root := DirectoryStructure{
		Level:    0,
		Path:     dir,
		Children: []DirectoryStructure{},
	}

	err := walk(&root, sorter)

	return root.Children, err
}

func walk(parent *DirectoryStructure, sorter sort.Sorter) error {
	entries, err := os.ReadDir(parent.Path)

	if err != nil {
		return err
	}

	sorter(entries)

	for _, v := range entries {
		fileInfo, err := v.Info()

		if err != nil {
			return err
		}

		self := DirectoryStructure{
			Level: parent.Level + 1,
			Path:  filepath.Join(parent.Path, v.Name()),
			Name:  v.Name(),
			IsDir: v.IsDir(),
			FileInfo: FileInfo{
				ModifiedAt: fileInfo.ModTime(),
				Size:       fileInfo.Size(),
			},
			Children: []DirectoryStructure{},
		}

		if self.IsDir {
			err = walk(&self, sorter)

			if err != nil {
				return err
			}
		}

		parent.Children = append(parent.Children, self)
	}

	return nil
}
