package sort

import (
	"io/fs"
	internalSort "sort"
)

func newModifiedAtDirEntrySorter(order Order) (Sorter, error) {
	return func(data []fs.DirEntry) []fs.DirEntry {
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
	}, nil
}
