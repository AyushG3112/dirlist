package sort

import (
	"fmt"
)

type Field string

var (
	Name       Field = "name"
	ModifiedAt Field = "modifiedAt"
)

func ToField(field string) (Field, error) {
	switch field {
	case "name":
		return Name, nil
	case "modifiedAt":
		return ModifiedAt, nil
	}

	return "", fmt.Errorf("Sorting field: %s is not allowed. Allowed values: name, modifiedAt", field)
}
