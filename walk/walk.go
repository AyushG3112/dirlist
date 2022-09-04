package walk

import (
	"errors"
	"io/fs"
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

type Walker interface {
	Walk(sorter sort.DirEntrySorter) ([]DirectoryStructure, error)
}

type walker struct {
	dirAbsPath string
}

func (w *walker) Walk(sorter sort.DirEntrySorter) ([]DirectoryStructure, error) {
	root := DirectoryStructure{
		Level:    0,
		AbsPath:  w.dirAbsPath,
		Children: []DirectoryStructure{},
	}

	err := w.walk(&root, sorter, w.dirAbsPath)

	return root.Children, err
}

func (w *walker) walk(parent *DirectoryStructure, sorter sort.DirEntrySorter, rootAbsPath string) error {
	entries, err := os.ReadDir(parent.AbsPath)

	if err != nil {
		return err
	}

	sorter.Sort(entries)

	for _, v := range entries {
		fileInfo, err := v.Info()
		if err != nil {
			if errors.Is(err, fs.ErrPermission) {
				continue
			}

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
			err = w.walk(&self, sorter, rootAbsPath)

			if err != nil {
				if errors.Is(err, fs.ErrPermission) {
					continue
				}

				return err
			}
		}

		parent.Children = append(parent.Children, self)
	}

	return nil
}

func NewWalker(dirAbsPath string) (Walker, error) {
	return &walker{dirAbsPath: dirAbsPath}, nil
}
