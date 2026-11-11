package errorspractice

import "errors"

// Sentinel errors that callers can check for with errors.Is.
var (
	ErrEmpty   = errors.New("amount is empty")
	ErrInvalid = errors.New("invalid amount")
)
