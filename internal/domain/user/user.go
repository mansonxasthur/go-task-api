package user

import (
	"errors"
)

type User struct {
	ID    Id
	Name  string
	Email Email
}

var (
	ErrorNameIsRequired error = errors.New("name is required")
)

func NewUser(name, email string) (*User, error) {
	if name == "" {
		return nil, ErrorNameIsRequired
	}

	emailObj, err := NewEmail(email)
	if err != nil {
		return nil, err
	}

	return &User{
		Name:  name,
		Email: *emailObj,
	}, nil
}

func (u *User) SetID(id Id) {
	u.ID = id
}

func (u *User) ToResource() map[string]interface{} {
	return map[string]interface{}{
		"id":    u.ID,
		"name":  u.Name,
		"email": u.Email.Value,
	}
}
