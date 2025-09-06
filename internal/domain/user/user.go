package user

import (
	"github.com/mansonxasthur/go-task-api/internal/domain/errors"
	"github.com/mansonxasthur/go-task-api/pkg/helpers"
)

type User struct {
	ID    ID
	Name  string
	Email Email
}

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

func (u *User) SetID(id ID) {
	u.ID = id
}

func (u *User) ToResource() map[string]interface{} {
	return map[string]interface{}{
		"id":    u.ID,
		"name":  u.Name,
		"email": u.Email.Value,
	}
}
