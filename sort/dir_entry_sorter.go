package sort

import (
	"fmt"
	"io/fs"
)

type DirEntrySorter interface {
	Sort([]fs.DirEntry) []fs.DirEntry
}

func NewDirEntrySorter(field Field, order Order) (DirEntrySorter, error) {
	switch field {
	case Name:
		return NewModifiedAtDirEntrySorter(order)
	case ModifiedAt:
		return NewModifiedAtDirEntrySorter(order)
	}

	return nil, fmt.Errorf("No sorter defined for field: %s, order %s", field, order)
}
