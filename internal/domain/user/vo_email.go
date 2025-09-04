package user

import (
	"errors"
	"fmt"
	"net/mail"
	"regexp"
)

var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)

var (
	ErrorEmailIsRequired    = errors.New(`email is required`)
	ErrorInvalidEmailFormat = errors.New(`invalid email format`)
)

type Email struct {
	Value string `json:"value"`
}

func NewEmail(value string) (*Email, error) {
	if err := validateEmail(value); err != nil {
		return nil, err
	}
	return &Email{
		Value: value,
	}, nil
}

func validateEmail(v string) error {
	if v == "" {
		return ErrorEmailIsRequired
	}

	if !emailRegex.MatchString(v) {
		return ErrorInvalidEmailFormat
	}

	_, err := mail.ParseAddress(v)
	if err != nil {
		return fmt.Errorf("invalid email address: %w", err)
	}

	return nil
}
