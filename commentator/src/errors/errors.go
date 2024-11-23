package errors

import "errors"

var (
	ErrNotFound              = errors.New("Not found")
	ErrMissingRequiredParams = errors.New("Missing required parameters")
)
