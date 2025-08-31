package repository

import (
	"fmt"

	"github.com/mansonxasthur/go-task-api/internal/domain/user"
)

type UserMemoryRepository struct {
	Users  map[int32]user.User
	lastID int32
}

func NewUserMemoryRepository() *UserMemoryRepository {
	users := make(map[int32]user.User)
	return &UserMemoryRepository{
		Users:  users,
		lastID: 0,
	}
}

func (r *UserMemoryRepository) Create(u *user.User) error {
	id := r.lastID + 1

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
