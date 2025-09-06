package errors

var (
	ErrorNameIsRequired  = NewValidationError("name is required")
	ErrUserNotFound      = NewNotFoundError("user")
	ErrUserAlreadyExists = NewConflictError("user")
)

func NewCreateUserError(err error) *DomainError {
	return NewError(
		"error creating user",
		InternalError,
		err,
	)
}
