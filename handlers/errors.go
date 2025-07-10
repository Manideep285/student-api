package handlers

import "errors"

// Error definitions for validation
var (
	ErrNameRequired  = errors.New("name is required")
	ErrInvalidAge    = errors.New("age must be a positive number")
	ErrEmailRequired = errors.New("email is required")
)
