package commands

import (
	"context"
	"errors"

	domainerrors "github.com/mansonxasthur/go-task-api/internal/domain/errors"
	"github.com/mansonxasthur/go-task-api/internal/domain/user"
	"github.com/mansonxasthur/go-task-api/pkg/helpers"
)

// RegisterUserCommand is the register user command.
type RegisterUserCommand struct {
	repo user.Repository
}

// NewRegisterUserCommand creates a new register user command.
func NewRegisterUserCommand(repo user.Repository) *RegisterUserCommand {
	return &RegisterUserCommand{
		repo: repo,
	}
}

// Execute creates a new user.
func (c *RegisterUserCommand) Execute(ctx context.Context, name, email string) (user.ID, error) {
	email = helpers.NormalizeEmail(email)
	u, err := user.NewUser(name, email)
	if err != nil {
		return 0, err
	}

	existingUser, err := c.repo.FindByEmail(ctx, email)
	if err != nil && !errors.Is(err, domainerrors.ErrorUserNotFound) {
		return 0, domainerrors.NewCreateUserError(err)
	}

	if existingUser != nil {
		return 0, domainerrors.ErrorUserAlreadyExists
	}

	id, err := c.repo.Create(ctx, u)

	if err != nil {
		return 0, domainerrors.NewCreateUserError(err)
	}

	return id, nil
}
