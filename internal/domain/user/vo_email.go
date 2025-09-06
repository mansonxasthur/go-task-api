package user

import (
	"fmt"
	"net/mail"
	"regexp"

	domainerrors "github.com/mansonxasthur/go-task-api/internal/domain/errors"
)

var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)

// Email is the email value object.
type Email struct {
	Value string
}

// NewEmail creates a new email value object.
func NewEmail(value string) (*Email, error) {
	if err := validateEmail(value); err != nil {
		return nil, err
	}
	return &Email{
		Value: value,
	}, nil
}

// validateEmail validates the email value.
func validateEmail(v string) error {
	if v == "" {
		return domainerrors.ErrorEmailIsRequired
	}

	if !emailRegex.MatchString(v) {
		return domainerrors.ErrorInvalidEmailFormat
	}

	_, err := mail.ParseAddress(v)
	if err != nil {
		return fmt.Errorf("invalid email address: %w", err)
	}

	return nil
}

// String returns the email value.
func (e *Email) String() string {
	return e.Value
}
