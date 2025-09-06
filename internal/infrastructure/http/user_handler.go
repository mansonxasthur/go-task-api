package http

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/mansonxasthur/go-task-api/internal/application"
	domainErrors "github.com/mansonxasthur/go-task-api/internal/domain/errors"
	"github.com/mansonxasthur/go-task-api/internal/domain/user"
	httphelper "github.com/mansonxasthur/go-task-api/pkg/http"
)

type UserHandler struct {
	userService *application.UserService
}

func NewUserHandler(s *application.UserService) *UserHandler {
	return &UserHandler{
		userService: s,
	}
}

func (h *UserHandler) Handle(mux *http.ServeMux) {
	mux.Handle("POST /users", http.HandlerFunc(h.createUserHandler))
	mux.Handle("GET /users", http.HandlerFunc(h.listUsersHandler))
}

func (h *UserHandler) createUserHandler(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		httphelper.ErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	ctx := r.Context()

	id, err := h.userService.RegisterUser(ctx, req.Name, req.Email)
	if err != nil {
		status := h.mapErrorToStatusCode(err)
		httphelper.ErrorResponse(w, err, status)
		return
	}

	u, err := h.userService.FindByID(ctx, id)

	httphelper.SuccessResponse(w, user.NewUserDtoFromEntity(u), http.StatusCreated)
}

func (h *UserHandler) listUsersHandler(w http.ResponseWriter, r *http.Request) {
	users := h.userService.ListUsers(r.Context())

	httphelper.SuccessResponse(w, user.NewUserDtoList(users), http.StatusOK)
}

func (h *UserHandler) mapErrorToStatusCode(err error) int {
	if errors.Is(err, domainErrors.ErrorNameIsRequired) {
		return http.StatusBadRequest
	}
	if errors.Is(err, domainErrors.ErrUserAlreadyExists) {
		return http.StatusConflict
	}
	if errors.Is(err, domainErrors.ErrUserNotFound) {
		return http.StatusNotFound
	}
	return http.StatusInternalServerError
}
