// Package errors encapsulates domain errors.
package errors

import "fmt"

// Code represents an error code.
type Code string

const (
	// ValidationError is the error code for validation errors.
	ValidationError Code = "VALIDATION_ERROR"
	// NotFoundError is the error code for "not found" errors.
	NotFoundError Code = "NOT_FOUND"
	// ConflictError is the error code for "conflict" errors.
	ConflictError Code = "CONFLICT_ERROR"
	// InternalError is the error code for internal errors.
	InternalError Code = "INTERNAL_ERROR"
)

// DomainError represents a domain error.
type DomainError struct {
	Code    Code
	Message string
	Cause   error
}

// Error returns the error message.
func (e *DomainError) Error() string {
	if e.Cause != nil {
		return fmt.Sprintf("%s: %v", e.Message, e.Cause)
	}
	return e.Message
}

// Unwrap returns the underlying error.
func (e *DomainError) Unwrap() error {
	return e.Cause
}

// NewValidationError creates a new validation error.
func NewValidationError(message string) *DomainError {
	return &DomainError{
		Code:    ValidationError,
		Message: message,
	}
}

// NewNotFoundError creates a new "not found" error.
func NewNotFoundError(resource string) *DomainError {
	return &DomainError{
		Code:    NotFoundError,
		Message: fmt.Sprintf("%s not found", resource),
	}
}

// NewConflictError creates a new "conflict" error.
func NewConflictError(resource string) *DomainError {
	return &DomainError{
		Code:    ConflictError,
		Message: fmt.Sprintf("%s already exists", resource),
	}
}

// NewError creates a new domain error.
func NewError(message string, code Code, err error) *DomainError {
	return &DomainError{
		Code:    code,
		Message: message,
		Cause:   err,
	}
}
