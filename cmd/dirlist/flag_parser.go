package main

import (
	"os"

	"github.com/spf13/pflag"
)

func parseCLIFlags() cliOptions {
	flag := pflag.NewFlagSet("dirlist", pflag.ExitOnError)
	rootDir := flag.StringP("root-dir", "d", ".", "Root directory to start directory listing. Defaults to $PWD")
	port := flag.IntP("port", "p", 8000, "Port on which to start the listing server")
	sortField := flag.StringP("sort-field", "f", "modifiedAt", "Field to sort by")
	sortOrder := flag.StringP("sort-order", "", "ASC", "Sorting order. ASC/DESC.")
	help := flag.BoolP("help", "h", false, "View help")
	cached := flag.BoolP("cached", "c", false, "Run in cached mode. Cached mode generates the structure once and always shows that even if the underlying structure has changed")
	flag.Parse(os.Args[1:])

	options := cliOptions{
		sortField: *sortField,
		sortOrder: *sortOrder,
		rootDir:   *rootDir,
		port:      *port,
		help:      *help,
		cached:    *cached,
		usages:    flag.FlagUsages(),
	}

	return options
}
