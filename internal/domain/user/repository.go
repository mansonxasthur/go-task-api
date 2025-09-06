package user

import "context"

// Repository is the user repository interface
type Repository interface {
	Create(context.Context, *User) (ID, error)
	FindByID(context.Context, ID) (*User, error)
	FindByEmail(context.Context, string) (*User, error)
	All(ctx context.Context) []*User
}
