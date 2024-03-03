package validators

import (
	"fmt"
	"reflect"
	"slices"
	"strconv"
)

//nolint:exhaustive
func MaxValidator(requireMax string, v reflect.Value) error {
	requireMaxVal, err := strconv.Atoi(requireMax)
	if err != nil {
		return fmt.Errorf("invalid Max argument")
	}

	IntKinds := []reflect.Kind{reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64}
	switch {
	case slices.Contains(IntKinds, v.Kind()):
		fmt.Println(v.Int(), "v.Int()")
		if v.Int() > int64(requireMaxVal) {
			return ErrInvalidMax
		}
	case v.Kind() == reflect.Slice && v.Len() > 0 && slices.Contains(IntKinds, v.Index(0).Kind()):
		for _, val := range v.Interface().([]int64) {
			if val > int64(requireMaxVal) {
				return ErrInvalidMax
			}
		}
	default:
		return ErrInvalidType
	}
	return nil
}
