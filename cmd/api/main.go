package main

import (
	"fmt"
	"net/http"

	handlers "github.com/mansonxasthur/go-task-api/internal/infrastructure/http"
	"github.com/mansonxasthur/go-task-api/internal/infrastructure/repository"
)

type App struct {
}

func main() {
	app := &App{}

	app.Run()
}

func (*App) Run() {
	mux := http.NewServeMux()
	// handling user requests
	userRepo := repository.NewUserMemoryRepository()
	userHandler := handlers.NewUserHandler(userRepo)
	userHandler.Process(mux)

	fmt.Println("Starting server on port 8080...")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		panic(err)
	}
}
