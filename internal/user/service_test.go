package user

import (
	"context"
	"errors"
	"testing"

	"github.com/PUDDLEEE/puddleee_back/ent"
	"github.com/PUDDLEEE/puddleee_back/internal/user/dto"
	"github.com/PUDDLEEE/puddleee_back/pkg/interfaces/mocks"
	"github.com/stretchr/testify/require"
)

func TestUserService_createUser(t *testing.T) {
	tests := []struct {
		name   string
		before func(*testing.T, *mocks.IUserRepository, *ent.Client, context.Context)
		expect func(*testing.T, *UserService, *ent.Client, context.Context)
	}{
		{
			name: "createUser/create_user_failed",
			before: func(t *testing.T, repo *mocks.IUserRepository, client *ent.Client, ctx context.Context) {
				repo.On("Create",
					ctx,
					client,
					dto.CreateUserDTO{},
				).
					Return(nil, errors.New("")).
					Once()
			},
			expect: func(t *testing.T, service *UserService, client *ent.Client, ctx context.Context) {
				user, err := service.CreateUser(dto.CreateUserDTO{})

				require.Error(t, err)
				require.Nil(t, user)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
			require.NoError(t, err)
			defer client.Close()
			require.NoError(t, client.Schema.Create(context.Background()))
			repo := mocks.NewIUserRepository(t)

			userService := NewService(repo, context.Background(), client)
			tt.before(t, repo, client, context.Background())
			tt.expect(t, &userService, client, context.Background())
		})
	}
}

func TestUserService_findOneUser(t *testing.T) {
	tests := []struct {
		name   string
		before func(*testing.T, *mocks.IUserRepository, *ent.Client, context.Context)
		expect func(*testing.T, *UserService, *ent.Client, context.Context)
	}{
		{
			name: "findOneUser/find_one_user_failed",
			before: func(t *testing.T, repo *mocks.IUserRepository, client *ent.Client, ctx context.Context) {
				repo.On("FindOneById",
					ctx,
					client,
					1,
				).
					Return(nil, errors.New("")).
					Once()
			},
			expect: func(t *testing.T, service *UserService, client *ent.Client, ctx context.Context) {
				user, err := service.FindOneUser(1)

				require.Error(t, err)
				require.Nil(t, user)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
			require.NoError(t, err)
			defer client.Close()
			require.NoError(t, client.Schema.Create(context.Background()))
			repo := mocks.NewIUserRepository(t)

			userService := NewService(repo, context.Background(), client)
			tt.before(t, repo, client, context.Background())
			tt.expect(t, &userService, client, context.Background())
		})
	}
}
