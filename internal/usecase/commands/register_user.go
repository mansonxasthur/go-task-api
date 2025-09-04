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
		repo: repo,
	}
}

func (c *RegisterUserCommand) Execute(name, email string) (*user.User, error) {
	u, err := user.NewUser(name, email)
	if err != nil {
		return nil, fmt.Errorf("error creating user: %w", err)
	}

	err = c.repo.Create(u)

	if err != nil {
		return nil, fmt.Errorf("error creating user: %w", err)
	}

	return u, nil
}
