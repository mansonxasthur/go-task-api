package commands

import (
	"context"
	"errors"
	"fmt"
	"testing"

	domainerrors "github.com/mansonxasthur/go-task-api/internal/domain/errors"
	"github.com/mansonxasthur/go-task-api/internal/infrastructure/repository"
)

func TestRegisterUserCommand_Success(t *testing.T) {
	type user struct {
		Name  string
		Email string
	}

	var cases = []user{
		{
			Name:  "Manson",
			Email: "mansonx13@gmail.com",
		},
		{
			Name:  "Xasthur",
			Email: "mansonxasthur@gmail.com",
		},
	}

	repo := repository.NewUserMemoryRepository()
	command := NewRegisterUserCommand(repo)
	ctx := context.Background()
	var count int32

	for i, u := range cases {
		t.Run(fmt.Sprintf("case-%d", i), func(t *testing.T) {
			id, err := command.Execute(ctx, u.Name, u.Email)
			if err != nil {
				t.Errorf("error executing command: %v", err)
			}

			user, err := repo.FindByID(ctx, id)
			if err != nil {
				t.Errorf("error finding user: %v", err)
			}

			count++

			userCount := len(repo.Users)
			if userCount != int(count) {
				t.Errorf("expected %d user created but got %d", count, userCount)
			}

			foundUser, err := repo.FindByID(ctx, user.ID)
			if err != nil {
				t.Errorf("error finding user: %v", err)
			}

			if foundUser.Name != user.Name || foundUser.Email.Value != user.Email.Value {
				t.Errorf("expected user name to be %s and email to be %s but got %s and %s", user.Name, user.Email.Value, foundUser.Name, foundUser.Email)
			}
		})
	}
}

func TestRegisterUserCommand_Validation(t *testing.T) {
	repo := repository.NewUserMemoryRepository()
	command := NewRegisterUserCommand(repo)
	ctx := context.Background()

	_, err := command.Execute(ctx, "", "")
	if err == nil {
		t.Errorf("expected validation error but got nil")
		return
	}

	if !errors.Is(err, domainerrors.ErrorNameIsRequired) {
		t.Errorf("expected error to be %v but got %v", domainerrors.ErrorNameIsRequired, err)
	}
}
