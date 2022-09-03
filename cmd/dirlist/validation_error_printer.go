package main

import (
	"fmt"
	"io"
)

func printValidationErrors(valErrors []string, stdout io.Writer) error {
	if len(valErrors) == 0 {
		return nil
	}
	fmt.Fprintln(stdout, "Could not process due to following errors: ")
	for _, v := range valErrors {
		_, err := fmt.Fprintf(stdout, " - %s\n", v)
		if err != nil {
			return err
		}
	}
	return nil
}
