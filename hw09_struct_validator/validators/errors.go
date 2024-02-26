package validators

import "errors"

var (
	ErrNotStruct     = errors.New("input value is not a struct")
	ErrInvalidLen    = errors.New("invalid length")
	ErrInvalidMax    = errors.New("invalid max value")
	ErrInvalidIn     = errors.New("invalid in value")
	ErrInvalidMin    = errors.New("invalid min value")
	ErrInvalidType   = errors.New("invalid type")
	ErrInvalidRegexp = errors.New("invalid regexp value")
)
