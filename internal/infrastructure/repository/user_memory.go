package repository

import (
	"fmt"

	"github.com/mansonxasthur/go-task-api/internal/domain/user"
)

type UserMemoryRepository struct {
	Users      map[int32]user.User
	emailIndex map[string]int32
	lastID     int32
}

func NewUserMemoryRepository() *UserMemoryRepository {
	return &UserMemoryRepository{
		Users:      make(map[int32]user.User),
		emailIndex: make(map[string]int32),
		lastID:     0,
	}
}

func (r *UserMemoryRepository) Create(u *user.User) error {
	if _, exists := r.emailIndex[u.Email.Value]; exists {
		return fmt.Errorf("email already exists")
	}
	id := r.lastID + 1

	u.SetID(user.Id{Value: id})
	r.emailIndex[u.Email.Value] = id
	r.Users[id] = *u
	r.lastID = id

	return nil
}

func (r *UserMemoryRepository) FindByID(id int32) (*user.User, error) {
	u, ok := r.Users[id]
	if !ok {
		return nil, fmt.Errorf("user not found")
	}

	return &u, nil
}

func (r *UserMemoryRepository) FindByEmail(email string) (*user.User, error) {
	id, ok := r.emailIndex[email]
	if !ok {
		return nil, fmt.Errorf("user not found")
	}

	return r.FindByID(id)
}

func (r *UserMemoryRepository) All() []*user.User {
	var users []*user.User

	for _, u := range r.Users {
		users = append(users, &u)
	}

	return users
}
