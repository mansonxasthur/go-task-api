package user

import (
	"fmt"
)

type User struct {
	ID    Id
	Name  string
	Email Email
}

func NewUser(name, email string) (*User, error) {
	if name == "" {
		return nil, fmt.Errorf("name and email are required")
	}

	emailObj, err := NewEmail(email)
	if err != nil {
		return nil, fmt.Errorf("invalid email: %w", err)
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
		"id":    u.ID.Value,
		"name":  u.Name,
		"email": u.Email.Value,
	}
}
