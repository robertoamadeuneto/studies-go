package internalerrors

import "errors"

var InternalServerError error = errors.New("Internal server error")
