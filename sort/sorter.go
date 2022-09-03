package sort

import (
	"fmt"
	"io/fs"
)

type Sorter = func([]fs.DirEntry) []fs.DirEntry

func NewSorter(field Field, order Order) (Sorter, error) {
	switch field {
	case Name:
		return newNameDirEntrySorter(order)
	case ModifiedAt:
		return newModifiedAtDirEntrySorter(order)
	}

	return nil, fmt.Errorf("No sorter defined for field: %s, order %s", field, order)
}
