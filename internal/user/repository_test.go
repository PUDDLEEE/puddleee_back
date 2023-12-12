package user

import (
	"context"
	"testing"

	"github.com/PUDDLEEE/puddleee_back/ent"
	userdto "github.com/PUDDLEEE/puddleee_back/internal/user/dto"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/require"
)

func TestUserRepository_Create(t *testing.T) {
	tests := []struct {
		name   string
		before func(*testing.T, *UserRepository, *ent.Client, context.Context)
		expect func(*testing.T, *UserRepository, *ent.Client, context.Context)
	}{
		{
			name: "Create/username_already_exists",
			before: func(t *testing.T, repo *UserRepository, client *ent.Client, ctx context.Context) {
				_, err := repo.Create(ctx, client, userdto.CreateUserDTO{
					Username: "test",
					Email:    "123",
					Password: "123",
				})
				require.NoError(t, err)
			},
			expect: func(t *testing.T, repo *UserRepository, client *ent.Client, ctx context.Context) {
				_, err := repo.Create(ctx, client, userdto.CreateUserDTO{
					Username: "test",
					Email:    "1",
					Password: "1",
				})

				require.Error(t, err)
			},
		},
		{
			name: "Create/email_already_exists",
			before: func(t *testing.T, repo *UserRepository, client *ent.Client, ctx context.Context) {
				_, err := repo.Create(ctx, client, userdto.CreateUserDTO{
					Username: "test",
					Email:    "test",
					Password: "123",
				})
				require.NoError(t, err)
			},
			expect: func(t *testing.T, repo *UserRepository, client *ent.Client, ctx context.Context) {
				_, err := repo.Create(ctx, client, userdto.CreateUserDTO{
					Username: "1",
					Email:    "test",
					Password: "1",
				})

				require.Error(t, err)
			},
		},
		{
			name: "Create/username_should_not_empty",
			before: func(t *testing.T, repository *UserRepository, client *ent.Client, ctx context.Context) {

			},
			expect: func(t *testing.T, repo *UserRepository, client *ent.Client, ctx context.Context) {
				_, err := repo.Create(ctx, client, userdto.CreateUserDTO{
					Username: "",
					Email:    "123",
					Password: "123",
				})

				require.Error(t, err)
				require.Contains(t, err.Error(), "username")
			},
		},
		{
			name: "Create/email_should_not_empty",
			before: func(t *testing.T, repository *UserRepository, client *ent.Client, ctx context.Context) {

			},
			expect: func(t *testing.T, repo *UserRepository, client *ent.Client, ctx context.Context) {
				_, err := repo.Create(ctx, client, userdto.CreateUserDTO{
					Username: "123",
					Email:    "",
					Password: "123",
				})

				require.Error(t, err)
				require.Contains(t, err.Error(), "email")
			},
		},
		{
			name: "Create/password_should_not_empty",
			before: func(t *testing.T, repository *UserRepository, client *ent.Client, ctx context.Context) {

			},
			expect: func(t *testing.T, repo *UserRepository, client *ent.Client, ctx context.Context) {
				_, err := repo.Create(ctx, client, userdto.CreateUserDTO{
					Username: "123",
					Email:    "123",
					Password: "",
				})

				require.Error(t, err)
				require.Contains(t, err.Error(), "password")
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
			require.NoError(t, err)
			defer client.Close()
			require.NoError(t, client.Schema.Create(context.Background()))

			repo := &UserRepository{}
			tt.before(t, repo, client, context.Background())
			tt.expect(t, repo, client, context.Background())
		})
	}
}

func TestUserRepository_FindOneById(t *testing.T) {
	tests := []struct {
		name   string
		before func(*testing.T, *UserRepository, *ent.Client, context.Context)
		expect func(*testing.T, *UserRepository, *ent.Client, context.Context)
	}{
		{
			name: "FindOneById/could_not_found",
			before: func(t *testing.T, repo *UserRepository, client *ent.Client, ctx context.Context) {

			},
			expect: func(t *testing.T, repo *UserRepository, client *ent.Client, ctx context.Context) {
				_, err := repo.FindOneById(ctx, client, 1)

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

			repo := &UserRepository{}
			tt.before(t, repo, client, context.Background())
			tt.expect(t, repo, client, context.Background())
		})
	}
}

func TestUserRepository_Update(t *testing.T) {
	tests := []struct {
		name   string
		before func(*testing.T, *UserRepository, *ent.Client, context.Context)
		expect func(*testing.T, *UserRepository, *ent.Client, context.Context)
	}{
		{
			name: "Update/could_not_update",
			before: func(t *testing.T, repo *UserRepository, client *ent.Client, ctx context.Context) {

			},
			expect: func(t *testing.T, repo *UserRepository, client *ent.Client, ctx context.Context) {
				_, err := repo.Update(ctx, client, 1, userdto.UpdateUserDTO{})

				require.Error(t, err)
			},
		},
		{
			name: "Update/could_not_empty_dto",
			before: func(t *testing.T, repo *UserRepository, client *ent.Client, ctx context.Context) {
				_, err := repo.Create(ctx, client, userdto.CreateUserDTO{
					Username: "test",
					Email:    "test",
					Password: "123",
				})
				require.NoError(t, err)
			},
			expect: func(t *testing.T, repo *UserRepository, client *ent.Client, ctx context.Context) {
				_, err := repo.Update(ctx, client, 1, userdto.UpdateUserDTO{})

				require.Error(t, err)
			},
		},
		{
			name: "Update/should_update_username_properly",
			before: func(t *testing.T, repo *UserRepository, client *ent.Client, ctx context.Context) {
				_, err := repo.Create(ctx, client, userdto.CreateUserDTO{
					Username: "test",
					Email:    "test",
					Password: "test",
				})
				require.NoError(t, err)
			},
			expect: func(t *testing.T, repo *UserRepository, client *ent.Client, ctx context.Context) {
				username := "hello"
				usernameUpdated, err := repo.Update(ctx, client, 1, userdto.UpdateUserDTO{
					Username: &username,
				})
				require.NoError(t, err)
				require.Equal(t, usernameUpdated.Username, "hello")
				require.Equal(t, usernameUpdated.Email, "test")
				require.Equal(t, usernameUpdated.Password, "test")
			},
		},
		{
			name: "Update/should_update_email_properly",
			before: func(t *testing.T, repo *UserRepository, client *ent.Client, ctx context.Context) {
				_, err := repo.Create(ctx, client, userdto.CreateUserDTO{
					Username: "test",
					Email:    "test",
					Password: "test",
				})
				require.NoError(t, err)
			},
			expect: func(t *testing.T, repo *UserRepository, client *ent.Client, ctx context.Context) {
				email := "123"
				usernameUpdated, err := repo.Update(ctx, client, 1, userdto.UpdateUserDTO{
					Email: &email,
				})
				require.NoError(t, err)
				require.Equal(t, usernameUpdated.Username, "test")
				require.Equal(t, usernameUpdated.Email, "123")
				require.Equal(t, usernameUpdated.Password, "test")
			},
		},
		{
			name: "Update/should_update_password_properly",
			before: func(t *testing.T, repo *UserRepository, client *ent.Client, ctx context.Context) {
				_, err := repo.Create(ctx, client, userdto.CreateUserDTO{
					Username: "test",
					Email:    "test",
					Password: "test",
				})
				require.NoError(t, err)
			},
			expect: func(t *testing.T, repo *UserRepository, client *ent.Client, ctx context.Context) {
				password := "123"
				usernameUpdated, err := repo.Update(ctx, client, 1, userdto.UpdateUserDTO{
					Password: &password,
				})
				require.NoError(t, err)
				require.Equal(t, usernameUpdated.Username, "test")
				require.Equal(t, usernameUpdated.Email, "test")
				require.Equal(t, usernameUpdated.Password, "123")
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
			require.NoError(t, err)
			defer client.Close()
			require.NoError(t, client.Schema.Create(context.Background()))

			repo := &UserRepository{}
			tt.before(t, repo, client, context.Background())
			tt.expect(t, repo, client, context.Background())
		})
	}
}

func TestUserRepository_Delete(t *testing.T) {
	tests := []struct {
		name   string
		before func(*testing.T, *UserRepository, *ent.Client, context.Context)
		expect func(*testing.T, *UserRepository, *ent.Client, context.Context)
	}{
		{
			name: "Delete/could_not_delete",
			before: func(t *testing.T, repo *UserRepository, client *ent.Client, ctx context.Context) {

			},
			expect: func(t *testing.T, repo *UserRepository, client *ent.Client, ctx context.Context) {
				err := repo.Delete(ctx, client, 1)

				require.Error(t, err)
			},
		},
		{
			name: "Delete/could_delete",
			before: func(t *testing.T, repo *UserRepository, client *ent.Client, ctx context.Context) {
				_, err := repo.Create(ctx, client, userdto.CreateUserDTO{
					Username: "test",
					Email:    "test",
					Password: "test",
				})
				require.NoError(t, err)
			},
			expect: func(t *testing.T, repo *UserRepository, client *ent.Client, ctx context.Context) {
				err := repo.Delete(ctx, client, 1)

				require.NoError(t, err)

				_, err = repo.FindOneById(ctx, client, 1)
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

			repo := &UserRepository{}
			tt.before(t, repo, client, context.Background())
			tt.expect(t, repo, client, context.Background())
		})
	}
}
