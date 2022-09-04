package sort

import (
	"io/fs"
	internalSort "sort"
)

func newModifiedAtDirEntrySorterFunc(order Order, data []fs.DirEntry) []fs.DirEntry {
	internalSort.Slice(data, func(i, j int) bool {
		iinfo, err := data[i].Info()

		if err != nil {
			panic(err)
		}

		jinfo, err := data[j].Info()

		if err != nil {
			panic(err)
		}

		if order == ASC {
			return iinfo.ModTime().Before(jinfo.ModTime())
		}

		return iinfo.ModTime().After(jinfo.ModTime())
	})

	return data
}

type ModifiedAtDirEntrySorter struct {
	order Order
}

func (m *ModifiedAtDirEntrySorter) Sort(entries []fs.DirEntry) []fs.DirEntry {
	return newModifiedAtDirEntrySorterFunc(m.order, entries)
}

func NewModifiedAtDirEntrySorter(order Order) (*ModifiedAtDirEntrySorter, error) {
	return &ModifiedAtDirEntrySorter{order: order}, nil
}
