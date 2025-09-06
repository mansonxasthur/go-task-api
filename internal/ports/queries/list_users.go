package queries

import (
	"context"

	"github.com/mansonxasthur/go-task-api/internal/domain/user"
)

// ListUsersQuery is the list users query.
type ListUsersQuery struct {
	repo user.Repository
}

// NewListUsersQuery creates a new list users query.
func NewListUsersQuery(repo user.Repository) *ListUsersQuery {
	return &ListUsersQuery{
		repo: repo,
	}
}

// Execute returns a list of users.
func (q *ListUsersQuery) Execute(ctx context.Context) []*user.User {
	return q.repo.All(ctx)
}
