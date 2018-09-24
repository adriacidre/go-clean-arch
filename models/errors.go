package models

import "errors"

var (
	// ErrInternalServer Internal server error
	ErrInternalServer = errors.New("Internal Server Error")

	// ErrNotFound Not found error
	ErrNotFound = errors.New("Your requested Item is not found")

	// ErrConflict Conflict error
	ErrConflict = errors.New("Your Item already exist")
)
