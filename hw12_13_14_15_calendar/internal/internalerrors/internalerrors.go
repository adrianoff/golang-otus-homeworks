package internalerrors

import "errors"

var ErrNotExistID = errors.New("the event with this id is not found")
