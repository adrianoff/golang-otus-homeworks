package validators

import (
	"reflect"
	"strconv"
	"strings"

	"golang.org/x/exp/slices"
)

//gocognit:ignore
func InValidator(inValues string, v reflect.Value) error {
	IntKinds := []reflect.Kind{reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64}
	switch {
	case v.Kind() == reflect.String:
		for _, val := range strings.Split(inValues, ",") {
			if v.String() == val {
				return nil
			}
		}

	case slices.Contains(IntKinds, v.Kind()):
		inVal := strings.Split(inValues, ",")

		for _, val := range inVal {
			valInt, err := strconv.Atoi(val)
			if err != nil {
				return err
			}

			if v.Int() == int64(valInt) {
				return nil
			}
		}
	case v.Kind() == reflect.Slice && v.Len() > 0 && v.Index(0).Kind() == reflect.String:
		inVal := strings.Split(inValues, ",")
		for _, val := range v.Interface().([]string) {
			if !slices.Contains(inVal, val) {
				return ErrInvalidIn
			}
		}
		return nil
	case v.Kind() == reflect.Slice && v.Len() > 0 && slices.Contains(IntKinds, v.Index(0).Kind()):
		inVal := strings.Split(inValues, ",")
		for _, val := range v.Interface().([]int64) {
			if !slices.Contains(inVal, strconv.FormatInt(val, 10)) {
				return ErrInvalidIn
			}
		}
		return nil

	default:
		return ErrInvalidType
	}

	return ErrInvalidIn
}
