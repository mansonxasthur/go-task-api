package commands

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/mansonxasthur/go-task-api/internal/domain/user"
	"github.com/mansonxasthur/go-task-api/internal/infrastructure/repository"
	"github.com/mansonxasthur/go-task-api/pkg/helpers"
)

var (
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

func (c *RegisterUserCommand) Execute(ctx context.Context, name, email string) (user.ID, error) {
	email = helpers.NormalizeEmail(email)
	u, err := user.NewUser(name, email)
	if err != nil {
		return 0, err
	}

	existingUser, err := c.repo.FindByEmail(ctx, email)
	if err != nil && !errors.Is(err, repository.ErrUserNotFound) {
		return 0, errCreatingUser(err)
	}

	if existingUser != nil {
		return 0, ErrUserAlreadyExists
	}

	id, err := c.repo.Create(ctx, u)

	if err != nil {
		return 0, errCreatingUser(err)
	}

	return id, nil
}

func logErrorCreatingUser(err error) {
	log.Printf("error creating user: %v\n", err)
}

func errCreatingUser(err error) error {
	logErrorCreatingUser(err)
	return errors.New(fmt.Sprintf("error creating user: %v", err))
}
