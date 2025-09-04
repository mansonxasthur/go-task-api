package commands

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/mansonxasthur/go-task-api/internal/domain/user"
)

var (
	ErrCreatingUser      = errors.New("error creating user")
	ErrUserAlreadyExists = errors.New("user already exists")
)

type RegisterUserCommand struct {
	repo user.Repository
}

func NewRegisterUserCommand(repo user.Repository) *RegisterUserCommand {
	return &RegisterUserCommand{
		repo: repo,
	}
}

func (c *RegisterUserCommand) Execute(ctx context.Context, name, email string) (*user.User, error) {
	email = strings.TrimSpace(strings.ToLower(email))
	u, err := user.NewUser(name, email)
	if err != nil {
		return nil, err
	}

	existingUser, err := c.repo.FindByEmail(ctx, email)
	if err != nil {
		logErrorCreatingUser(err)
		return nil, fmt.Errorf("error finding user by email: %w", err)
	}

	if existingUser != nil {
		return nil, ErrUserAlreadyExists
	}

	_, err = c.repo.Create(ctx, u)

	if err != nil {
		logErrorCreatingUser(err)
		return nil, ErrCreatingUser
	}

	return u, nil
}

func logErrorCreatingUser(err error) {
	log.Printf("error creating user: %v\n", err)
}
