package sort

import (
	"io/fs"
	internalSort "sort"
)

func newNameDirEntrySorter(order Order) (Sorter, error) {
	return func(data []fs.DirEntry) []fs.DirEntry {
		internalSort.Slice(data, func(i, j int) bool {
			if order == ASC {
				return data[i].Name() < data[j].Name()
			}

			return data[i].Name() > data[j].Name()
		})

		return data
	}, nil
}
