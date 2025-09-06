package application

import (
	"context"

	"github.com/mansonxasthur/go-task-api/internal/domain/user"
	"github.com/mansonxasthur/go-task-api/internal/ports/commands"
	"github.com/mansonxasthur/go-task-api/internal/ports/queries"
)

type UserService struct {
	repo user.Repository
}

func NewUserService(repo user.Repository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) RegisterUser(ctx context.Context, name, email string) (user.ID, error) {
	return commands.NewRegisterUserCommand(s.repo).Execute(ctx, name, email)
}

func (s *UserService) ListUsers(ctx context.Context) []*user.User {
	return queries.NewListUsersQuery(s.repo).Execute(ctx)
}

func (s *UserService) FindByID(ctx context.Context, id user.ID) (*user.User, error) {
	return s.repo.FindByID(ctx, id)
}
