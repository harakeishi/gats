package gats

import (
	"errors"
	"fmt"
)

func ToString(a any) (string, error) {
	var result string
	switch a.(type) {
	case int, int8, int16, int32, int64:
		result = fmt.Sprintf("%d", a)
	case uint, uint8, uint16, uint32, uint64:
		result = fmt.Sprintf("%d", a)
	case float32, float64:
		result = fmt.Sprintf("%g", a)
	case complex64, complex128:
		result = fmt.Sprintf("%g", a)
	case []byte:
		result = fmt.Sprintf("%s", a)
	case bool:
		result = fmt.Sprintf("%t", a)
	case string:
		result = fmt.Sprintf("%s", a)
	case nil:
		result = "nil"
	default:
		return "", errors.New("No suitable type found")
	}
	return result, nil
}
