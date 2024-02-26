package validators

import (
	"reflect"
	"regexp"
)

//nolint:exhaustive
func RegexValidator(regexpValue string, v reflect.Value) error {
	re := regexp.MustCompile(regexpValue)

	switch v.Kind() {
	case reflect.String:
		str := re.MatchString(v.String())
		if !str {
			return ErrInvalidRegexp
		}
	default:
		return ErrInvalidType
	}
	return nil
}
