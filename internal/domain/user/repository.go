package user

import "context"

type Repository interface {
	Create(context.Context, *User) (ID, error)
	FindByID(context.Context, ID) (*User, error)
	FindByEmail(context.Context, string) (*User, error)
	All(ctx context.Context) []*User
}
