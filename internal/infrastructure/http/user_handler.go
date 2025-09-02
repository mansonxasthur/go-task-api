package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/mansonxasthur/go-task-api/internal/domain/user"
	commands "github.com/mansonxasthur/go-task-api/internal/usecase/commands"
	"github.com/mansonxasthur/go-task-api/internal/usecase/queries"
)

type UserHandler struct {
	repo user.Repository
	ctx  context.Context
}

func NewUserHandler(ctx context.Context, repo user.Repository) *UserHandler {
	return &UserHandler{
		repo: repo,
		ctx:  ctx,
	}
}

func (h *UserHandler) Process(mux *http.ServeMux) {
	mux.Handle("POST /users", http.HandlerFunc(h.createUserHandler))
	mux.Handle("GET /users", http.HandlerFunc(h.getUsersHandler))
}

func (h *UserHandler) createUserHandler(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		ErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	u, err := commands.NewRegisterUserCommand(h.repo).Execute(req.Name, req.Email)

	if err != nil {
		ErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	SuccessResponse(w, u.ToResource(), http.StatusCreated)
}

func (h *UserHandler) getUsersHandler(w http.ResponseWriter, r *http.Request) {
	users := queries.NewListUsersQuery(h.repo).Execute()
	var usersResource []map[string]interface{}
	for _, u := range users {
		usersResource = append(usersResource, u.ToResource())
	}

	SuccessResponse(w, usersResource, http.StatusOK)
}
