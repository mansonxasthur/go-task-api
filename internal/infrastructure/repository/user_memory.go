package repository

import (
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
	ErrEmailAlreadyExists = errors.New("email already exists")
	ErrUserNotFound       = errors.New("user not found")
)

func NewUserMemoryRepository() *UserMemoryRepository {
	return &UserMemoryRepository{
		Users:      make(map[int32]user.User),
		emailIndex: make(map[string]int32),
		lastID:     0,
	}
}

func (r *UserMemoryRepository) Create(u *user.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.emailIndex[u.Email.Value]; exists {
		return ErrEmailAlreadyExists
	}

	id := r.lastID + 1

	u.SetID(user.Id(id))
	r.emailIndex[u.Email.Value] = id
	r.Users[id] = *u
	r.lastID = id

	return nil
}

func (r *UserMemoryRepository) FindByID(id int32) (*user.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	u, ok := r.Users[id]
	if !ok {
		return nil, ErrUserNotFound
	}

	return &u, nil
}

func (r *UserMemoryRepository) FindByEmail(email string) (*user.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	id, ok := r.emailIndex[email]
	if !ok {
		return nil, ErrUserNotFound
	}

	return r.FindByID(id)
}

func (r *UserMemoryRepository) All() []*user.User {
	r.mu.RLock()
	defer r.mu.RUnlock()
	
	var users []*user.User

	for _, u := range r.Users {
		users = append(users, &u)
	}

	return users
}
