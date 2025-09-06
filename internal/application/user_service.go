package application

import (
	"context"

	"github.com/mansonxasthur/go-task-api/internal/domain/user"
	"github.com/mansonxasthur/go-task-api/internal/ports/commands"
	"github.com/mansonxasthur/go-task-api/internal/ports/queries"
)

// Service is the user service interface
type Service interface {
	RegisterUser(ctx context.Context, name, email string) (user.ID, error)
	ListUsers(ctx context.Context) []*user.User
	FindByID(ctx context.Context, id user.ID) (*user.User, error)
}

// UserService is the user service implementation
type UserService struct {
	repo user.Repository
}

// NewUserService creates a new user service
func NewUserService(repo user.Repository) *UserService {
	return &UserService{
		repo: repo,
	}
}

// RegisterUser creates a new user
func (s *UserService) RegisterUser(ctx context.Context, name, email string) (user.ID, error) {
	return commands.NewRegisterUserCommand(s.repo).Execute(ctx, name, email)
}

// ListUsers lists all users
func (s *UserService) ListUsers(ctx context.Context) []*user.User {
	return queries.NewListUsersQuery(s.repo).Execute(ctx)
}

// FindByID finds a user by id
func (s *UserService) FindByID(ctx context.Context, id user.ID) (*user.User, error) {
	return s.repo.FindByID(ctx, id)
}
