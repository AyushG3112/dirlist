package main

import (
	"fmt"
	"io"
	"os"

	"github.com/ayushg3112/dirlist"
)

func main() {
	if exitCode := start(os.Stdout); exitCode != 0 {
		os.Exit(exitCode)
	}
}

func start(stdout io.Writer) int {
	options := parseCLIFlags()

	if options.help {
		fmt.Fprintf(stdout, "Usage: \n%s", options.usages)
		return 0
	}

	validationErrors := options.validate()
	if len(validationErrors) != 0 {
		err := printValidationErrors(validationErrors, stdout)
		if err != nil {
			fmt.Fprintf(stdout, "failed to print validation errors: %s\n", err.Error())
			return 3
		}
		return 4
	}

	err := dirlist.StartServer(options.toProcessingOptions())

	if err != nil {
		fmt.Fprintf(stdout, "failed to start server : %s\n", err.Error())
		return 5
	}

	return 0
}
