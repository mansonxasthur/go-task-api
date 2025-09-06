package user

// Dto represents a User DTO.
type Dto struct {
	ID    int32  `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// NewUserDtoFromEntity creates a new Dto from a User entity.
func NewUserDtoFromEntity(u *User) *Dto {
	return &Dto{
		ID:    int32(u.ID),
		Name:  u.Name,
		Email: u.Email.String(),
	}
}

// NewUserDtoList creates a new Dto list from a User list.
func NewUserDtoList(users []*User) []*Dto {
	userList := make([]*Dto, len(users))
	for i, u := range users {
		userList[i] = NewUserDtoFromEntity(u)
	}

	return userList
}
