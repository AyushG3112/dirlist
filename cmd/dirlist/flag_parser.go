package main

import (
	"os"

	"github.com/spf13/pflag"
)

func parseCLIFlags() cliOptions {
	flag := pflag.NewFlagSet("dirlist", pflag.ExitOnError)
	rootDir := flag.StringP("root-dir", "d", "", "Root directory to start directory listing")
	port := flag.IntP("port", "p", 8000, "Port on which to start the listing server")
	sortField := flag.StringP("sort-field", "f", "modifiedAt", "Field to sort by")
	sortOrder := flag.StringP("sort-order", "", "ASC", "Sorting order. ASC/DESC.")
	help := flag.Bool("help", false, "View help")
	flag.Parse(os.Args[1:])

	options := cliOptions{
		sortField: *sortField,
		sortOrder: *sortOrder,
		rootDir:   *rootDir,
		port:      *port,
		help:      *help,
		usages:    flag.FlagUsages(),
	}

	return options
}
