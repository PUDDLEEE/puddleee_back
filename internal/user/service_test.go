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

func TestUserService_CreateUser(t *testing.T) {
	tests := []struct {
		name   string
		before func(*testing.T, *mocks.IUserRepository, *ent.Client, context.Context)
		expect func(*testing.T, *UserService, *ent.Client, context.Context)
	}{
		{
			name: "CreateUser/create_user_failed",
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
			tt.expect(t, userService, client, context.Background())
		})
	}
}

func TestUserService_FindOneUser(t *testing.T) {
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
			tt.expect(t, userService, client, context.Background())
		})
	}
}

func TestUserService_UpdateUser(t *testing.T) {
	tests := []struct {
		name   string
		before func(*testing.T, *mocks.IUserRepository, *ent.Client, context.Context)
		expect func(*testing.T, *UserService, *ent.Client, context.Context)
	}{
		{
			name: "UpdateUser/update_user_failed",
			before: func(t *testing.T, repo *mocks.IUserRepository, client *ent.Client, ctx context.Context) {
				repo.On("Update",
					ctx,
					client,
					1,
					dto.UpdateUserDTO{},
				).
					Return(nil, errors.New("")).
					Once()
			},
			expect: func(t *testing.T, service *UserService, client *ent.Client, ctx context.Context) {
				user, err := service.UpdateUser(1, dto.UpdateUserDTO{})

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
			tt.expect(t, userService, client, context.Background())
		})
	}
}

func TestUserService_DeleteUser(t *testing.T) {
	tests := []struct {
		name   string
		before func(*testing.T, *mocks.IUserRepository, *ent.Client, context.Context)
		expect func(*testing.T, *UserService, *ent.Client, context.Context)
	}{
		{
			name: "DeleteUser/delete_user_failed",
			before: func(t *testing.T, repo *mocks.IUserRepository, client *ent.Client, ctx context.Context) {
				repo.On("Delete",
					ctx,
					client,
					1,
				).
					Return(errors.New("")).
					Once()
			},
			expect: func(t *testing.T, service *UserService, client *ent.Client, ctx context.Context) {
				err := service.DeleteUser(1)
				require.Error(t, err)
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
			tt.expect(t, userService, client, context.Background())
		})
	}
}
