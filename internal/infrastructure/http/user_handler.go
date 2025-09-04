package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/mansonxasthur/go-task-api/internal/domain/user"
	"github.com/mansonxasthur/go-task-api/internal/usecase/commands"
	"github.com/mansonxasthur/go-task-api/internal/usecase/queries"
)

type UserHandler struct {
	repo user.Repository
}

func NewUserHandler(repo user.Repository) *UserHandler {
	return &UserHandler{
		repo: repo,
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

	ctx := r.Context()

	id, err := commands.NewRegisterUserCommand(h.repo).Execute(ctx, req.Name, req.Email)

	if err != nil {
		ErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	u, err := h.repo.FindByID(ctx, id)

	SuccessResponse(w, user.NewUserDtoFromEntity(u), http.StatusCreated)
}

func (h *UserHandler) getUsersHandler(w http.ResponseWriter, r *http.Request) {
	users := queries.NewListUsersQuery(h.repo).Execute(r.Context())

	SuccessResponse(w, user.NewUserDtoList(users), http.StatusOK)
}
