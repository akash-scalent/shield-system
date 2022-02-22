package utils

import (
	"fmt"
)

func CheckIfEmpty(nameOfField string, field interface{}) error {
	switch v := field.(type) {
	case int:
		if v == 0 {
			return fmt.Errorf("%s cannot be 0", nameOfField)
		}
	case string:
		if v == "" {
			return fmt.Errorf("%s cannot be empty", nameOfField)
		}
	case []interface{}:
		if len(v)==0 {
			return fmt.Errorf("%s cannot be empty", nameOfField)
		}

	}
	return nil
}
