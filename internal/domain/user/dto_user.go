package user

type Dto struct {
	ID    int32  `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func NewUserDtoFromEntity(u *User) *Dto {
	return &Dto{
		ID:    int32(u.ID),
		Name:  u.Name,
		Email: u.Email.String(),
	}
}

func NewUserDtoList(users []*User) []*Dto {
	userList := make([]*Dto, len(users))
	for i, u := range users {
		userList[i] = NewUserDtoFromEntity(u)
	}

	return userList
}
