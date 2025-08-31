package commands

import (
	"fmt"

	"github.com/mansonxasthur/go-task-api/internal/domain/user"
)

type RegisterUserCommand struct {
	repo user.Repository
}

func NewRegisterUserCommand(repo user.Repository) *RegisterUserCommand {
	return &RegisterUserCommand{
		repo,
	}
}

func (c *RegisterUserCommand) Execute(name, email string) error {
	if name == "" || email == "" {
		return fmt.Errorf("name and email are required")
	}

	u := user.User{
		Name:  name,
		Email: email,
	}

	err := c.repo.Create(&u)

	if err != nil {
		return fmt.Errorf("error creating user: %w", err)
	}

	return nil
}
