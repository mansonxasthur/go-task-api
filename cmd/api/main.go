package main

import (
	"fmt"

	"github.com/mansonxasthur/go-task-api/internal/domain/user"
	"github.com/mansonxasthur/go-task-api/internal/infrastructure/repository"
	commands "github.com/mansonxasthur/go-task-api/internal/usecase/commnds"
)

func main() {
	u := &user.User{
		Name:  "Manson",
		Email: "mansonx13@gmail.com",
	}

	repo := repository.NewUserMemoryRepository()
	command := commands.NewRegisterUserCommand(repo)
	command.Execute(u.Name, u.Email)
	fmt.Println(repo.Users)
}
