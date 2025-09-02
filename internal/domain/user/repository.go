package user

type Repository interface {
	Create(*User) error
	FindByID(int32) (*User, error)
	FindByEmail(string) (*User, error)
	All() []*User
}
