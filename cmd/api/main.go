package main

import (
	"context"
	"fmt"
	"net/http"

	handlers "github.com/mansonxasthur/go-task-api/internal/infrastructure/http"
	"github.com/mansonxasthur/go-task-api/internal/infrastructure/repository"
)

type App struct {
	ctx context.Context
}

func main() {
	app := &App{
		ctx: context.Background(),
	}

	app.Run()
}

func (a *App) Run() {
	mux := http.NewServeMux()
	// handling user requests
	userRepo := repository.NewUserMemoryRepository()
	userHandler := handlers.NewUserHandler(a.ctx, userRepo)
	userHandler.Process(mux)

	fmt.Println("Starting server on port 8080...")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		panic(err)
	}
}
