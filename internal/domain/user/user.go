package user

import (
	"github.com/mansonxasthur/go-task-api/internal/domain/errors"
	"github.com/mansonxasthur/go-task-api/pkg/helpers"
)

// ID represents a user ID.
type ID int32

// User is the user domain entity.
type User struct {
	ID    ID
	Name  string
	Email Email
}

// NewUser creates a new user domain entity.
func NewUser(name, email string) (*User, error) {
	if name == "" {
		return nil, errors.ErrorNameIsRequired
	}
	email = helpers.NormalizeEmail(email)
	emailObj, err := NewEmail(email)
	if err != nil {
		return nil, err
	}

	return &User{
		Name:  name,
		Email: *emailObj,
	}, nil
}

// SetID sets the user ID.
func (u *User) SetID(id ID) {
	u.ID = id
}
