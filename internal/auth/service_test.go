package auth

import (
	"context"
	"errors"
	"testing"

	"github.com/PUDDLEEE/puddleee_back/ent"
	authdto "github.com/PUDDLEEE/puddleee_back/internal/auth/dto"
	customError "github.com/PUDDLEEE/puddleee_back/internal/errors"
	"github.com/PUDDLEEE/puddleee_back/internal/jwtAuth"
	userdto "github.com/PUDDLEEE/puddleee_back/internal/user/dto"
	"github.com/PUDDLEEE/puddleee_back/pkg/interfaces/mocks"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestAuthService_CreateEmailVerification(t *testing.T) {
	tests := []struct {
		name   string
		before func(*testing.T, *mocks.IAuthRepository, *ent.Client, context.Context, uuid.UUID, string)
		expect func(*testing.T, *AuthService, *ent.Client, context.Context, uuid.UUID, string)
	}{
		{
			name: "CreateEmailVerification/create_verification_failed",
			before: func(t *testing.T, repo *mocks.IAuthRepository, client *ent.Client, ctx context.Context, uuid uuid.UUID, code string) {

				repo.On("GenerateUUIDFromString", "").Return(uuid).Once()
				repo.On("GenerateCode", 6).Return(code).Once()
				repo.On("Create", ctx, client, authdto.CreateVerificationDTO{
					UUID: uuid.String(),
					Code: code,
				}).Return(nil, errors.New("")).Once()
			},
			expect: func(t *testing.T, service *AuthService, client *ent.Client, ctx context.Context, uuid uuid.UUID, code string) {
				_, err := service.CreateEmailVerification("")
				require.Error(t, err)
			},
		},
		{
			name: "CreateEmailVerification/create_verification_success",
			before: func(t *testing.T, repo *mocks.IAuthRepository, client *ent.Client, ctx context.Context, uuid uuid.UUID, code string) {

				repo.On("GenerateUUIDFromString", "").Return(uuid).Once()
				repo.On("GenerateCode", 6).Return(code).Once()
				repo.On("Create", ctx, client, authdto.CreateVerificationDTO{
					UUID: uuid.String(),
					Code: code,
				}).Return(&ent.Verification{UUID: uuid, Code: code}, nil).Once()
			},
			expect: func(t *testing.T, service *AuthService, client *ent.Client, ctx context.Context, uuid uuid.UUID, code string) {
				verification, err := service.CreateEmailVerification("")

				require.Equal(t, code, verification.Code)
				require.Equal(t, uuid.String(), verification.UUID.String())
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

			authRepo := mocks.NewIAuthRepository(t)
			userService := mocks.NewIUserService(t)
			jwtService := mocks.NewIJwtAuthService(t)

			params := &AuthServiceParameter{}
			params.authRepo = authRepo
			params.userService = userService
			params.jwtService = jwtService
			params.ctx = context.Background()
			params.client = client

			authService := NewAuthService(params)

			newUUID := uuid.New()
			code := "123456"
			tt.before(t, authRepo, client, context.Background(), newUUID, code)
			tt.expect(t, authService, client, context.Background(), newUUID, code)
		})
	}
}

func TestAuthService_VerifyEmailVerification(t *testing.T) {
	tests := []struct {
		name   string
		before func(*testing.T, *mocks.IAuthRepository, *ent.Client, context.Context)
		expect func(*testing.T, *AuthService, *ent.Client, context.Context)
	}{
		{
			name: "VerifyEmailVerification/found_verification_failed",
			before: func(t *testing.T, repo *mocks.IAuthRepository, client *ent.Client, ctx context.Context) {
				repo.On("FindOneByUUID", ctx, client, "").Return(nil, errors.New("")).Once()
			},
			expect: func(t *testing.T, service *AuthService, client *ent.Client, ctx context.Context) {
				err := service.VerifyEmailVerification("", "")
				require.Error(t, err)
			},
		},
		{
			name: "VerifyEmailVerification/verification_code_not_match",
			before: func(t *testing.T, repo *mocks.IAuthRepository, client *ent.Client, ctx context.Context) {
				repo.On("FindOneByUUID", ctx, client, "").Return(&ent.Verification{Code: "123456"}, nil).Once()
			},
			expect: func(t *testing.T, service *AuthService, client *ent.Client, ctx context.Context) {
				err := service.VerifyEmailVerification("", "123")
				require.ErrorIs(t, err, customError.NOT_AUTHORIZED)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
			require.NoError(t, err)
			defer client.Close()
			require.NoError(t, client.Schema.Create(context.Background()))

			authRepo := mocks.NewIAuthRepository(t)
			userService := mocks.NewIUserService(t)
			jwtService := mocks.NewIJwtAuthService(t)

			params := &AuthServiceParameter{}
			params.authRepo = authRepo
			params.userService = userService
			params.jwtService = jwtService
			params.ctx = context.Background()
			params.client = client

			authService := NewAuthService(params)

			tt.before(t, authRepo, client, context.Background())
			tt.expect(t, authService, client, context.Background())
		})
	}
}

func TestAuthService_Signup(t *testing.T) {
	tests := []struct {
		name   string
		before func(*testing.T, *mocks.IUserService, *ent.Client, context.Context)
		expect func(*testing.T, *AuthService, *ent.Client, context.Context)
	}{
		{
			name: "Signup/create_user_failed",
			before: func(t *testing.T, service *mocks.IUserService, client *ent.Client, ctx context.Context) {
				service.On("CreateUser", userdto.CreateUserDTO{}).Return(nil, errors.New("")).Once()
			},
			expect: func(t *testing.T, service *AuthService, client *ent.Client, ctx context.Context) {
				_, err := service.Signup(userdto.CreateUserDTO{})
				require.Error(t, err)
			},
		},
		{
			name: "Signup/create_user_success",
			before: func(t *testing.T, service *mocks.IUserService, client *ent.Client, ctx context.Context) {
				service.On("CreateUser", userdto.CreateUserDTO{}).Return(&ent.User{
					Username: "1",
					Email:    "1",
					Password: "1",
				}, nil).Once()
			},
			expect: func(t *testing.T, service *AuthService, client *ent.Client, ctx context.Context) {
				user, err := service.Signup(userdto.CreateUserDTO{})
				require.NotEmpty(t, user)
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

			authRepo := mocks.NewIAuthRepository(t)
			userService := mocks.NewIUserService(t)
			jwtService := mocks.NewIJwtAuthService(t)

			params := &AuthServiceParameter{}
			params.authRepo = authRepo
			params.userService = userService
			params.jwtService = jwtService
			params.ctx = context.Background()
			params.client = client

			authService := NewAuthService(params)

			tt.before(t, userService, client, context.Background())
			tt.expect(t, authService, client, context.Background())
		})
	}
}

func TestAuthService_Signin(t *testing.T) {
	tests := []struct {
		name   string
		before func(*testing.T, *mocks.IUserService, *mocks.IJwtAuthService, *ent.Client, context.Context)
		expect func(*testing.T, *AuthService, *ent.Client, context.Context)
	}{
		{
			name: "Signin/find_user_failed",
			before: func(t *testing.T, userService *mocks.IUserService, jwtService *mocks.IJwtAuthService, client *ent.Client, ctx context.Context) {
				userService.On("FindOneUserByEmail", "").Return(nil, errors.New("")).Once()
			},
			expect: func(t *testing.T, service *AuthService, client *ent.Client, ctx context.Context) {
				_, err := service.Signin(authdto.SigninDTO{})
				require.Error(t, err)
			},
		},
		{
			name: "Signin/create_access_token_failed",
			before: func(t *testing.T, userService *mocks.IUserService, jwtService *mocks.IJwtAuthService, client *ent.Client, ctx context.Context) {
				userService.On("FindOneUserByEmail", "").Return(&ent.User{}, nil).Once()
				jwtService.On("CreateAccessToken", jwtAuth.AuthTokenClaims{}).Return(nil, errors.New("")).Once()
			},
			expect: func(t *testing.T, service *AuthService, client *ent.Client, ctx context.Context) {
				_, err := service.Signin(authdto.SigninDTO{})
				require.Error(t, err)
			},
		},
		{
			name: "Signin/create_refresh_token_failed",
			before: func(t *testing.T, userService *mocks.IUserService, jwtService *mocks.IJwtAuthService, client *ent.Client, ctx context.Context) {
				someTokenString := ""
				userService.On("FindOneUserByEmail", "").Return(&ent.User{}, nil).Once()
				jwtService.On("CreateAccessToken", jwtAuth.AuthTokenClaims{}).Return(&someTokenString, nil).Once()
				jwtService.On("CreateRefreshToken", jwtAuth.AuthTokenClaims{}).Return(nil, errors.New("")).Once()
			},
			expect: func(t *testing.T, service *AuthService, client *ent.Client, ctx context.Context) {
				_, err := service.Signin(authdto.SigninDTO{})
				require.Error(t, err)
			},
		},
		{
			name: "Signin/success_signin",
			before: func(t *testing.T, userService *mocks.IUserService, jwtService *mocks.IJwtAuthService, client *ent.Client, ctx context.Context) {
				someTokenString := "1"
				userService.On("FindOneUserByEmail", "").Return(&ent.User{}, nil).Once()
				jwtService.On("CreateAccessToken", jwtAuth.AuthTokenClaims{}).Return(&someTokenString, nil).Once()
				jwtService.On("CreateRefreshToken", jwtAuth.AuthTokenClaims{}).Return(&someTokenString, nil).Once()
			},
			expect: func(t *testing.T, service *AuthService, client *ent.Client, ctx context.Context) {
				tokens, err := service.Signin(authdto.SigninDTO{})
				require.Equal(t, "1", tokens.AccessToken)
				require.Equal(t, "1", tokens.RefreshToken)
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

			authRepo := mocks.NewIAuthRepository(t)
			userService := mocks.NewIUserService(t)
			jwtService := mocks.NewIJwtAuthService(t)

			params := &AuthServiceParameter{}
			params.authRepo = authRepo
			params.userService = userService
			params.jwtService = jwtService
			params.ctx = context.Background()
			params.client = client

			authService := NewAuthService(params)

			tt.before(t, userService, jwtService, client, context.Background())
			tt.expect(t, authService, client, context.Background())
		})
	}
}
