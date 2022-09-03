package sort

import "fmt"

type Order string

const ASC Order = "ASC"
const DESC Order = "DESC"

func ToOrder(field string) (Order, error) {
	switch field {
	case "ASC":
		return ASC, nil
	case "DESC":
		return DESC, nil
	}

	return "", fmt.Errorf("Sorting Order: %s is not allowed. Allowed values: ASC, DESC", field)
}
