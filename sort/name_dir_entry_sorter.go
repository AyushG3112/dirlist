package sort

import (
	"io/fs"
	internalSort "sort"
)

func newNameDirEntrySorterFunc(order Order, data []fs.DirEntry) []fs.DirEntry {
	internalSort.Slice(data, func(i, j int) bool {
		if order == ASC {
			return data[i].Name() < data[j].Name()
		}

		return data[i].Name() > data[j].Name()
	})

	return data
}

type NameDirEntrySorter struct {
	order Order
}

func (n *NameDirEntrySorter) Sort(entries []fs.DirEntry) []fs.DirEntry {
	return newNameDirEntrySorterFunc(n.order, entries)
}

func NewNameDirEntrySorter(order Order) (*NameDirEntrySorter, error) {
	return &NameDirEntrySorter{order: order}, nil
}
