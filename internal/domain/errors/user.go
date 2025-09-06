package errors

import "errors"

var (
	// ErrorNameIsRequired is the error for the missing name.
	ErrorNameIsRequired = NewValidationError("name is required")
	// ErrorUserNotFound is the error for not found user.
	ErrorUserNotFound = NewNotFoundError("user")
	// ErrorUserAlreadyExists is the error for the user already exists error.
	ErrorUserAlreadyExists = NewConflictError("user")
	// ErrorEmailIsRequired is the error for missing email.
	ErrorEmailIsRequired = errors.New(`email is required`)
	// ErrorInvalidEmailFormat is the error for invalid email format.
	ErrorInvalidEmailFormat = errors.New(`invalid email format`)
)

// NewCreateUserError creates a new error for user creation
func NewCreateUserError(err error) *DomainError {
	return NewError(
		"error creating user",
		InternalError,
		err,
	)
}
