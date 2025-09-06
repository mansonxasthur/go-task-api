package errors

import "fmt"

type Code string

const (
	ValidationError Code = "VALIDATION_ERROR"
	NotFoundError   Code = "NOT_FOUND"
	ConflictError   Code = "CONFLICT_ERROR"
	InternalError   Code = "INTERNAL_ERROR"
)

type DomainError struct {
	Code    Code
	Message string
	Cause   error
}

func (e *DomainError) Error() string {
	if e.Cause != nil {
		return fmt.Sprintf("%s: %v", e.Message, e.Cause)
	}
	return e.Message
}

func (e *DomainError) Unwrap() error {
	return e.Cause
}

func NewValidationError(message string) *DomainError {
	return &DomainError{
		Code:    ValidationError,
		Message: message,
	}
}

func NewNotFoundError(resource string) *DomainError {
	return &DomainError{
		Code:    NotFoundError,
		Message: fmt.Sprintf("%s not found", resource),
	}
}

func NewConflictError(resource string) *DomainError {
	return &DomainError{
		Code:    ConflictError,
		Message: fmt.Sprintf("%s already exists", resource),
	}
}

func NewErrorWithCause(err *DomainError, cause error) *DomainError {
	err.Cause = cause
	return err
}

func NewError(message string, code Code, err error) *DomainError {
	return &DomainError{
		Code:    code,
		Message: message,
		Cause:   err,
	}
}
