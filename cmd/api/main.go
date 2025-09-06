package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/mansonxasthur/go-task-api/internal/application"
	handlers "github.com/mansonxasthur/go-task-api/internal/infrastructure/http"
	"github.com/mansonxasthur/go-task-api/internal/infrastructure/repository"
)

// App is the application struct.
type App struct{}

func main() {
	app := &App{}

	app.run()
}

// run runs the application.
func (a *App) run() {
	mux := http.NewServeMux()

	// wiring up
	wire(mux)

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	srv := &http.Server{
		Handler: mux,
		Addr:    ":8080",
	}

	go func() {
		<-ctx.Done()
		fmt.Println("Shutting down...")
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := srv.Shutdown(shutdownCtx); err != nil {
			log.Printf("graceful shutdown failed: %v", err)
		}
	}()

	log.Printf("server started on port %s", srv.Addr)
	if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatalf("server error: %v \n", err)
	}
}

func wire(mux *http.ServeMux) {
	// handling user requests
	userRepo := repository.NewUserMemoryRepository()
	userService := application.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)
	userHandler.Handle(mux)
}
