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
	AbsPath  string
	RelPath  string
	IsDir    bool
	FileInfo FileInfo
	Children []DirectoryStructure
}

func Walk(dirAbsPath string, sorter sort.Sorter) ([]DirectoryStructure, error) {
	root := DirectoryStructure{
		Level:    0,
		AbsPath:  dirAbsPath,
		Children: []DirectoryStructure{},
	}

	err := walk(&root, sorter, dirAbsPath)

	return root.Children, err
}

func walk(parent *DirectoryStructure, sorter sort.Sorter, rootAbsPath string) error {
	entries, err := os.ReadDir(parent.AbsPath)

	if err != nil {
		return err
	}

	sorter(entries)

	for _, v := range entries {
		fileInfo, err := v.Info()

		if err != nil {
			return err
		}

		absPath := filepath.Join(parent.AbsPath, v.Name())
		relpath, err := filepath.Rel(rootAbsPath, absPath)

		if err != nil {
			return err
		}

		self := DirectoryStructure{
			Level:   parent.Level + 1,
			AbsPath: absPath,
			RelPath: relpath,
			Name:    v.Name(),
			IsDir:   v.IsDir(),
			FileInfo: FileInfo{
				ModifiedAt: fileInfo.ModTime(),
				Size:       fileInfo.Size(),
			},
			Children: []DirectoryStructure{},
		}

		if self.IsDir {
			err = walk(&self, sorter, rootAbsPath)

			if err != nil {
				return err
			}
		}

		parent.Children = append(parent.Children, self)
	}

	return nil
}
