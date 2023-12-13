package auth

import (
	"context"
	"testing"

	"github.com/PUDDLEEE/puddleee_back/ent"
	authdto "github.com/PUDDLEEE/puddleee_back/internal/auth/dto"
	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/require"
)

func TestAuthRepository_Create(t *testing.T) {
	tests := []struct {
		name   string
		before func(*testing.T, *AuthRepository, *ent.Client, context.Context, string)
		expect func(*testing.T, *AuthRepository, *ent.Client, context.Context, string)
	}{
		{
			name: "Create/uuid_already_exists",
			before: func(t *testing.T, repo *AuthRepository, client *ent.Client, ctx context.Context, uuid string) {
				dto := authdto.CreateVerificationDTO{
					UUID: uuid,
					Code: repo.GenerateCode(6),
				}
				_, err := repo.Create(ctx, client, dto)

				require.NoError(t, err)
			},
			expect: func(t *testing.T, repo *AuthRepository, client *ent.Client, ctx context.Context, uuid string) {
				dto := authdto.CreateVerificationDTO{
					UUID: uuid,
					Code: repo.GenerateCode(6),
				}
				_, err := repo.Create(ctx, client, dto)

				require.Error(t, err)
			},
		},
		{
			name: "Create/create_verification",
			before: func(t *testing.T, repo *AuthRepository, client *ent.Client, ctx context.Context, uuid string) {

			},
			expect: func(t *testing.T, repo *AuthRepository, client *ent.Client, ctx context.Context, uuid string) {
				code := repo.GenerateCode(6)
				dto := authdto.CreateVerificationDTO{
					UUID: uuid,
					Code: code,
				}
				createdVerification, err := repo.Create(ctx, client, dto)

				require.Equal(t, uuid, createdVerification.UUID.String())
				require.Equal(t, code, createdVerification.Code)
				require.NoError(t, err)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
			require.NoError(t, err)
			defer client.Close()
			require.NoError(t, client.Schema.Create(context.Background()))

			uuid := uuid.New().String()
			repo := &AuthRepository{}
			tt.before(t, repo, client, context.Background(), uuid)
			tt.expect(t, repo, client, context.Background(), uuid)
		})
	}
}

func TestAuthRepository_Update(t *testing.T) {
	tests := []struct {
		name   string
		before func(*testing.T, *AuthRepository, *ent.Client, context.Context, string)
		expect func(*testing.T, *AuthRepository, *ent.Client, context.Context, string)
	}{
		{
			name: "Update/verification_not_found",
			before: func(t *testing.T, repo *AuthRepository, client *ent.Client, ctx context.Context, uuidStr string) {

			},
			expect: func(t *testing.T, repo *AuthRepository, client *ent.Client, ctx context.Context, uuidStr string) {
				dto := authdto.UpdateVerificationDTO{
					Code: repo.GenerateCode(6),
				}
				_, err := repo.Update(ctx, client, uuidStr, dto)

				require.Error(t, err)
			},
		},
		{
			name: "Update/update_verification",
			before: func(t *testing.T, repo *AuthRepository, client *ent.Client, ctx context.Context, uuidStr string) {
				dto := authdto.CreateVerificationDTO{
					UUID: uuidStr,
					Code: repo.GenerateCode(6),
				}
				_, err := repo.Create(ctx, client, dto)

				require.NoError(t, err)
			},
			expect: func(t *testing.T, repo *AuthRepository, client *ent.Client, ctx context.Context, uuidStr string) {
				code := repo.GenerateCode(6)
				dto := authdto.UpdateVerificationDTO{
					Code: code,
				}
				updatedVerification, err := repo.Update(ctx, client, uuidStr, dto)
				require.Equal(t, code, updatedVerification.Code)
				require.Equal(t, uuidStr, updatedVerification.UUID.String())
				require.NoError(t, err)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
			require.NoError(t, err)
			defer client.Close()
			require.NoError(t, client.Schema.Create(context.Background()))

			uuid := uuid.New().String()
			repo := &AuthRepository{}
			tt.before(t, repo, client, context.Background(), uuid)
			tt.expect(t, repo, client, context.Background(), uuid)
		})
	}
}

func TestAuthRepository_FindOneByUUID(t *testing.T) {
	tests := []struct {
		name   string
		before func(*testing.T, *AuthRepository, *ent.Client, context.Context, string, string)
		expect func(*testing.T, *AuthRepository, *ent.Client, context.Context, string, string)
	}{
		{
			name: "FindOneByUUID/verification_not_found",
			before: func(t *testing.T, repo *AuthRepository, client *ent.Client, ctx context.Context, uuidStr string, code string) {

			},
			expect: func(t *testing.T, repo *AuthRepository, client *ent.Client, ctx context.Context, uuidStr string, code string) {

				_, err := repo.FindOneByUUID(ctx, client, uuidStr)

				require.Error(t, err)
			},
		},
		{
			name: "FindOneByUUID/verification_found",
			before: func(t *testing.T, repo *AuthRepository, client *ent.Client, ctx context.Context, uuidStr string, code string) {
				dto := authdto.CreateVerificationDTO{
					UUID: uuidStr,
					Code: code,
				}
				_, err := repo.Create(ctx, client, dto)

				require.NoError(t, err)
			},
			expect: func(t *testing.T, repo *AuthRepository, client *ent.Client, ctx context.Context, uuidStr string, code string) {

				verification, err := repo.FindOneByUUID(ctx, client, uuidStr)

				require.Equal(t, code, verification.Code)
				require.Equal(t, uuidStr, verification.UUID.String())
				require.NoError(t, err)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
			require.NoError(t, err)
			defer client.Close()
			require.NoError(t, client.Schema.Create(context.Background()))

			uuid := uuid.New().String()
			repo := &AuthRepository{}
			code := repo.GenerateCode(6)
			tt.before(t, repo, client, context.Background(), uuid, code)
			tt.expect(t, repo, client, context.Background(), uuid, code)
		})
	}
}

func TestAuthRepository_Delete(t *testing.T) {
	tests := []struct {
		name   string
		before func(*testing.T, *AuthRepository, *ent.Client, context.Context, string, string)
		expect func(*testing.T, *AuthRepository, *ent.Client, context.Context, string, string)
	}{
		{
			name: "Delete/verification_not_found",
			before: func(t *testing.T, repo *AuthRepository, client *ent.Client, ctx context.Context, uuidStr string, code string) {

			},
			expect: func(t *testing.T, repo *AuthRepository, client *ent.Client, ctx context.Context, uuidStr string, code string) {

				err := repo.Delete(ctx, client, uuidStr)

				require.Error(t, err)
			},
		},
		{
			name: "Delete/delete_verification",
			before: func(t *testing.T, repo *AuthRepository, client *ent.Client, ctx context.Context, uuidStr string, code string) {
				dto := authdto.CreateVerificationDTO{
					UUID: uuidStr,
					Code: code,
				}
				_, err := repo.Create(ctx, client, dto)
				require.NoError(t, err)

			},
			expect: func(t *testing.T, repo *AuthRepository, client *ent.Client, ctx context.Context, uuidStr string, code string) {

				err := repo.Delete(ctx, client, uuidStr)
				require.NoError(t, err)

				_, err = repo.FindOneByUUID(ctx, client, uuidStr)
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

			uuid := uuid.New().String()
			repo := &AuthRepository{}
			code := repo.GenerateCode(6)
			tt.before(t, repo, client, context.Background(), uuid, code)
			tt.expect(t, repo, client, context.Background(), uuid, code)
		})
	}
}

func TestAuthRepository_GenerateUUIDFromString(t *testing.T) {
	tests := []struct {
		name   string
		expect func(*testing.T, *AuthRepository, string, string)
	}{
		{
			name: "GenerateUUIDFromString/uuid_should_be_diffrent",
			expect: func(t *testing.T, repo *AuthRepository, email string, email2 string) {

				uuid1 := repo.GenerateUUIDFromString(email)
				uuid2 := repo.GenerateUUIDFromString(email2)
				require.NotEqual(t, uuid1, uuid2)
			},
		},
		{
			name: "GenerateUUIDFromString/uuid_should_be_equal",
			expect: func(t *testing.T, repo *AuthRepository, email string, email2 string) {

				uuid1 := repo.GenerateUUIDFromString(email)
				uuid2 := repo.GenerateUUIDFromString(email)
				require.Equal(t, uuid1, uuid2)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			repo := &AuthRepository{}
			email1 := "123@naver.com"
			email2 := "321@naver.com"
			tt.expect(t, repo, email1, email2)
		})
	}
}

func TestAuthRepository_GenerateCode(t *testing.T) {
	tests := []struct {
		name   string
		expect func(*testing.T, *AuthRepository)
	}{
		{
			name: "GenerateCode/code_must_be_generated_by_len",
			expect: func(t *testing.T, repo *AuthRepository) {

				code1 := repo.GenerateCode(6)
				code2 := repo.GenerateCode(3)
				require.Len(t, code1, 6)
				require.Len(t, code2, 3)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			repo := &AuthRepository{}
			tt.expect(t, repo)
		})
	}
}
