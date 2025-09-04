package repository

import (
	"context"
	"errors"
	"sync"

	"github.com/mansonxasthur/go-task-api/internal/domain/user"
)

type UserMemoryRepository struct {
	mu         sync.RWMutex
	Users      map[int32]user.User
	emailIndex map[string]int32
	lastID     int32
}

var (
	ErrUserNotFound = errors.New("user not found")
)

func NewUserMemoryRepository() *UserMemoryRepository {
	return &UserMemoryRepository{
		Users:      make(map[int32]user.User),
		emailIndex: make(map[string]int32),
		lastID:     0,
	}
}

func (r *UserMemoryRepository) Create(ctx context.Context, u *user.User) (user.Id, error) {
	if err := ctx.Err(); err != nil {
		return 0, err
	}

	r.mu.Lock()
	defer r.mu.Unlock()

	id := r.lastID + 1

	u.SetID(user.Id(id))
	r.emailIndex[u.Email.Value] = id
	r.Users[id] = *u
	r.lastID = id

	return u.ID, nil
}

func (r *UserMemoryRepository) FindByID(ctx context.Context, id int32) (*user.User, error) {
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
	id, ok := r.emailIndex[email]
	if !ok {
		return nil, nil
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
