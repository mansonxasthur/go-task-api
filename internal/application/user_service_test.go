package application

import (
	"context"
	"testing"

	domainerrors "github.com/mansonxasthur/go-task-api/internal/domain/errors"
	"github.com/mansonxasthur/go-task-api/internal/domain/user"
	"github.com/mansonxasthur/go-task-api/internal/infrastructure/repository"
)

var names = []string{
	"John",
	"Jane",
	"Josh",
	"Jessica",
	"Jennifer",
}

var emails = []string{
	"john@example.com",
	"jane@example.com",
	"josh@example.com",
	"jessica@example.com",
	"jennifer@example.com",
}

func TestNewUserService(t *testing.T) {
	repo := newMockRepository(context.Background(), makeUsers(0))

	service := NewUserService(repo)
	if service == nil {
		t.Errorf("expected service to be not nil")
		return
	}

	if service.repo != repo {
		t.Errorf("expected service.repo to be equal to repo")
		return
	}
}

func TestUserService_RegisterUser(t *testing.T) {
	ctx := context.Background()

	type Data struct {
		Name  string
		Email string
	}

	testCases := []struct {
		name    string
		data    Data
		wantID  user.ID
		wantErr error
	}{
		{
			name: "success",
			data: Data{
				Name:  "John",
				Email: "john@example.com",
			},
			wantID:  1,
			wantErr: nil,
		},
		{
			name: "name is required error",
			data: Data{
				Name:  "",
				Email: "",
			},
			wantID:  0,
			wantErr: domainerrors.ErrorNameIsRequired,
		},
		{
			name: "email is required error",
			data: Data{
				Name:  "John",
				Email: "",
			},
			wantID:  0,
			wantErr: domainerrors.ErrorEmailIsRequired,
		},
		{
			name: "email is invalid error",
			data: Data{
				Name:  "John",
				Email: "john@example",
			},
			wantID:  0,
			wantErr: domainerrors.ErrorInvalidEmailFormat,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			repo := newMockRepository(ctx, makeUsers(0))
			service := NewUserService(repo)
			id, err := service.RegisterUser(ctx, tt.data.Name, tt.data.Email)

			if err != nil {
				if tt.wantErr == nil {
					t.Errorf("expected no error but got %v", err)
					return
				}
				if tt.wantErr.Error() != err.Error() {
					t.Errorf("expected error %v but got %v", tt.wantErr, err)
					return
				}
			}

			if id != tt.wantID {
				t.Errorf("expected id %d but got %d", tt.wantID, id)
			}
		})
	}
}

func TestUserService_ListUsers(t *testing.T) {
	ctx := context.Background()

	testCases := []struct {
		name         string
		createdUsers int
		wantCount    int
	}{
		{
			name:         "Returns 5 users",
			createdUsers: 5,
			wantCount:    5,
		},
		{
			name:         "Returns 3 users",
			createdUsers: 3,
			wantCount:    3,
		},
		{
			name:         "Returns 0 users",
			createdUsers: 0,
			wantCount:    0,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			repo := newMockRepository(ctx, makeUsers(tt.createdUsers))
			service := NewUserService(repo)

			users := service.ListUsers(ctx)
			if len(users) != tt.wantCount {
				t.Errorf("expected %d users but got %d", tt.wantCount, len(users))
			}
		})
	}
}

func TestUserService_FindByID(t *testing.T) {
	ctx := context.Background()
	repo := newMockRepository(ctx, makeUsers(5))
	service := NewUserService(repo)

	tests := []struct {
		name       string
		id         user.ID
		wantedUser *user.User
		wantID     user.ID
		wantErr    error
	}{
		{
			name: "success",
			id:   user.ID(1),
			wantedUser: &user.User{
				ID:    1,
				Name:  "John",
				Email: user.Email{Value: "john@example.com"},
			},
			wantID: user.ID(1),
		},
		{
			name:    "not found error",
			id:      user.ID(6),
			wantErr: domainerrors.ErrorUserNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u, err := service.FindByID(ctx, tt.id)

			if err != nil {
				if tt.wantErr == nil {
					t.Errorf("expected no error but got %v", err)
					return
				}
				if tt.wantErr.Error() != err.Error() {
					t.Errorf("expected error %v but got %v", tt.wantErr, err)
					return
				}
				if u != nil {
					t.Errorf("expected user to be nil but got %v", u)
					return
				}
				return
			} else {
				if tt.wantErr != nil {
					t.Errorf("expected error %v but got nil", tt.wantErr)
					return
				}
			}

			if u == nil {
				t.Errorf("expected user to be not nil")
				return
			}

			if u.ID != tt.wantID {
				t.Errorf("expected id %d but got %d", tt.wantID, u.ID)
			}

			if u.Name != tt.wantedUser.Name {
				t.Errorf("expected name %s but got %s", tt.wantedUser.Name, u.Name)
			}

			if u.Email.Value != tt.wantedUser.Email.Value {
				t.Errorf("expected email %s but got %s", tt.wantedUser.Email.Value, u.Email.Value)
			}
		})
	}
}

func newMockRepository(ctx context.Context, users []*user.User) user.Repository {
	repo := repository.NewUserMemoryRepository()

	if len(users) == 0 {
		return repo
	}

	for _, u := range users {
		_, err := repo.Create(ctx, u)
		if err != nil {
			panic(err)
		}
	}

	return repo
}

func makeUsers(count int) []*user.User {
	if count == 0 {
		return []*user.User{}
	}

	var users []*user.User

	for i := 0; i < count; i++ {
		u, err := user.NewUser(names[i], emails[i])
		if err != nil {
			panic(err)
		}
		users = append(users, u)
	}

	return users
}
