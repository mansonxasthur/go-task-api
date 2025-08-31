package commands

import (
	"testing"

	"github.com/mansonxasthur/go-task-api/internal/infrastructure/repository"
)

func TestRegisterUserCommand_Success(t *testing.T) {
	repo := repository.NewUserMemoryRepository()
	command := NewRegisterUserCommand(repo)

	name := "Manson"
	email := "mansonx13@gmail.com"

	if err := command.Execute(name, email); err != nil {
		t.Errorf("error executing command: %v", err)
	}

	userCount := len(repo.Users)
	if userCount != 1 {
		t.Errorf("expected 1 user created but got %d", userCount)
	}

	u := repo.Users[1]

	if u.Name != name || u.Email != email {
		t.Errorf("expected user name to be %s and email to be %s but got %s and %s", name, email, u.Name, u.Email)
	}
}

func TestRegisterUserCommand_Validation(t *testing.T) {
	repo := repository.NewUserMemoryRepository()
	command := NewRegisterUserCommand(repo)

	if err := command.Execute("", ""); err == nil {
		t.Errorf("expected validation error but got nil")
	}
}
