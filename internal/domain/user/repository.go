package user

type Repository interface {
	Create(user *User) error
	FindByID(id int32) (*User, error)
}
