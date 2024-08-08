package hw09structvalidator

import (
	"fmt"
	"log"
	"reflect"
	"strings"

	"github.com/adrianoff/golang-otus-homework/hw09_struct_validator/validators"
)

type ValidationError struct {
	Field string
	Err   error
}

type ValidationErrors []ValidationError

func (v ValidationErrors) Error() string {
	if len(v) == 0 {
		return ""
	}

	var sb strings.Builder
	sb.WriteString("validation failed: ")
	for i, err := range v {
		if i > 0 {
			sb.WriteString(", ")
		}
		sb.WriteString(fmt.Sprintf("%s: %s", err.Field, err.Err))
	}
	return sb.String()
}

func Validate(v interface{}) error {
	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Struct {
		return validators.ErrNotStruct
	}

	var errs ValidationErrors
	rt := rv.Type()
	for i := 0; i < rv.NumField(); i++ {
		field := rv.Field(i)

		tag := rt.Field(i).Tag.Get("validate")
		if tag == "" {
			continue
		}

		if validationErrs := FieldValidate(field, tag); validationErrs != nil {
			for _, err := range validationErrs {
				errs = append(errs, ValidationError{
					Field: rt.Field(i).Name,
					Err:   err,
				})
			}
		}
	}

	if len(errs) > 0 {
		return errs
	}
	return nil
}

func FieldValidate(v reflect.Value, tag string) []error {
	rules := strings.Split(tag, "|")
	var errs []error

	for _, ruleField := range rules {
		if len(ruleField) < 2 {
			continue
		}

		args := strings.Split(ruleField, ":")[1:][0]
		rule := strings.Split(ruleField, ":")[0]

		switch rule {
		case "len":
			err := validators.LenValidator(args, v)
			if err != nil {
				errs = append(errs, err)
			}
		case "regexp":
			err := validators.RegexValidator(args, v)
			if err != nil {
				errs = append(errs, err)
			}
		case "in":
			err := validators.InValidator(args, v)
			if err != nil {
				errs = append(errs, err)
			}
		case "min":
			err := validators.MinValidator(args, v)
			if err != nil {
				errs = append(errs, err)
			}
		case "max":
			err := validators.MaxValidator(args, v)
			if err != nil {
				errs = append(errs, err)
			}
		default:
			log.Fatal("Invalid rule")
		}
	}
	return errs
}
