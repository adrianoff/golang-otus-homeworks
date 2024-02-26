package validators

import (
	"fmt"
	"reflect"
	"strconv"
)

//nolint:exhaustive
func MaxValidator(requireMax string, v reflect.Value) error {
	requireMaxVal, err := strconv.Atoi(requireMax)
	if err != nil {
		return fmt.Errorf("invalid Max argument")
	}

	switch v.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		fmt.Println(v.Int(), "v.Int()")
		if v.Int() > int64(requireMaxVal) {
			return ErrInvalidMax
		}
	case reflect.Slice:
		for _, val := range v.Interface().([]int) {
			if val > requireMaxVal {
				return ErrInvalidMax
			}
		}
	default:
		return ErrInvalidType
	}
	return nil
}
