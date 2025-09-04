package repository

import (
	"context"
	"errors"
	"sync"

	"github.com/mansonxasthur/go-task-api/internal/domain/user"
	"github.com/mansonxasthur/go-task-api/pkg/helpers"
)

type UserMemoryRepository struct {
	mu         sync.RWMutex
	Users      map[user.ID]user.User
	emailIndex map[string]user.ID
	lastID     user.ID
}

var (
	ErrUserNotFound = errors.New("user not found")
)

func NewUserMemoryRepository() *UserMemoryRepository {
	return &UserMemoryRepository{
		Users:      make(map[user.ID]user.User),
		emailIndex: make(map[string]user.ID),
		lastID:     0,
	}
}

func (r *UserMemoryRepository) Create(ctx context.Context, u *user.User) (user.ID, error) {
	if err := ctx.Err(); err != nil {
		return 0, err
	}

	r.mu.Lock()
	defer r.mu.Unlock()

	id := r.lastID + 1

	u.SetID(id)
	r.emailIndex[u.Email.Value] = id
	r.Users[id] = *u
	r.lastID = id

	return u.ID, nil
}

func (r *UserMemoryRepository) FindByID(ctx context.Context, id user.ID) (*user.User, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	r.mu.RLock()
	defer r.mu.RUnlock()

	u, ok := r.Users[id]
	if !ok {
		return nil, ErrUserNotFound
	}

	return &u, nil
}

func (r *UserMemoryRepository) FindByEmail(ctx context.Context, email string) (*user.User, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	r.mu.RLock()
	defer r.mu.RUnlock()
	email = helpers.NormalizeEmail(email)
	id, ok := r.emailIndex[email]
	if !ok {
		return nil, ErrUserNotFound
	}

	return r.FindByID(ctx, id)
}

func (r *UserMemoryRepository) All(ctx context.Context) []*user.User {
	if err := ctx.Err(); err != nil {
		return []*user.User{}
	}

	r.mu.RLock()
	defer r.mu.RUnlock()

	var users []*user.User

	for _, u := range r.Users {
		users = append(users, &u)
	}

	return users
}
