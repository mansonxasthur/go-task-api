package queries

import (
	"context"

	"github.com/mansonxasthur/go-task-api/internal/domain/user"
)

type ListUsersQuery struct {
	repo user.Repository
}

func NewListUsersQuery(repo user.Repository) *ListUsersQuery {
	return &ListUsersQuery{
		repo: repo,
	}
}

func (q *ListUsersQuery) Execute(ctx context.Context) []*user.User {
	return q.repo.All(ctx)
}
