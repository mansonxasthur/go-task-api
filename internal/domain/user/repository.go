package user

import "context"

type Repository interface {
	Create(context.Context, *User) (Id, error)
	FindByID(context.Context, int32) (*User, error)
	FindByEmail(context.Context, string) (*User, error)
	All(ctx context.Context) []*User
}
