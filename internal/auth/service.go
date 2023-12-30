package auth

import (
	"context"

	"github.com/PUDDLEEE/puddleee_back/ent"
	authdto "github.com/PUDDLEEE/puddleee_back/internal/auth/dto"
	"github.com/PUDDLEEE/puddleee_back/internal/errors"
	"github.com/PUDDLEEE/puddleee_back/internal/jwtAuth"
	"github.com/PUDDLEEE/puddleee_back/internal/mail"
	"github.com/PUDDLEEE/puddleee_back/internal/user"
	userdto "github.com/PUDDLEEE/puddleee_back/internal/user/dto"
	"github.com/PUDDLEEE/puddleee_back/pkg/interfaces"
	"github.com/PUDDLEEE/puddleee_back/pkg/interfaces/mocks"
)

type AuthService struct {
	authRepo    interfaces.IAuthRepository
	userService interfaces.IUserService
	jwtService  interfaces.IJwtAuthService
	mailService interfaces.IMailService
	ctx         context.Context
	client      *ent.Client
}

type AuthServiceParameter struct {
	AuthService
}

func (a *AuthService) Signin(dto authdto.SigninDTO) (*authdto.SigninOutputDTO, error) {
	user, err := a.userService.FindOneUserByEmail(dto.Email)
	if err != nil {
		return nil, err
	}
	claims := jwtAuth.AuthTokenClaims{
		ID: user.ID,
	}

	accessToken, err := a.jwtService.CreateAccessToken(claims)
	if err != nil {
		return nil, err
	}
	refreshToken, err := a.jwtService.CreateRefreshToken(claims)
	if err != nil {
		return nil, err
	}
	return &authdto.SigninOutputDTO{
		AccessToken:  *accessToken,
		RefreshToken: *refreshToken,
	}, nil
}

func (a *AuthService) Signup(dto userdto.CreateUserDTO) (*ent.User, error) {
	newUser, err := a.userService.CreateUser(dto)
	if err != nil {
		return nil, err
	}
	return newUser, nil
}

func (a *AuthService) CreateEmailVerification(email string) (*ent.Verification, error) {
	newUUID := a.authRepo.GenerateUUIDFromString(email).String()
	newCode := a.authRepo.GenerateCode(6)
	newVerification, err := a.authRepo.Create(a.ctx, a.client, authdto.CreateVerificationDTO{
		UUID: newUUID,
		Code: newCode,
	})
	if err != nil {
		return nil, err
	}
	return newVerification, nil
}

func (a *AuthService) SendEmailVerification(email string, verification *ent.Verification) error {
	err := a.mailService.Send(email, verification.Code)
	if err != nil {
		return err
	}
	return nil
}

func (a *AuthService) VerifyEmailVerification(uuidStr string, code string) error {
	verification, err := a.authRepo.FindOneByUUID(a.ctx, a.client, uuidStr)
	if err != nil {
		return err
	}

	if verification.Code != code {
		return errors.NOT_AUTHORIZED
	}

	return nil
}

func NewAuthService(params *AuthServiceParameter) *AuthService {
	if authRepo, ok := params.authRepo.(*AuthRepository); ok {
		userService, _ := params.userService.(*user.UserService)
		jwtService, _ := params.jwtService.(*jwtAuth.JwtAuthService)
		mailService, _ := params.mailService.(*mail.MailService)
		return &AuthService{
			authRepo:    authRepo,
			userService: userService,
			jwtService:  jwtService,
			mailService: mailService,
			ctx:         params.ctx,
			client:      params.client,
		}
	}

	if mockAuthRepo, ok := params.authRepo.(*mocks.IAuthRepository); ok {
		mockUserService, _ := params.userService.(*mocks.IUserService)
		mockJwtService, _ := params.jwtService.(*mocks.IJwtAuthService)
		mockMailService, _ := params.mailService.(*mocks.IMailService)
		return &AuthService{
			authRepo:    mockAuthRepo,
			userService: mockUserService,
			jwtService:  mockJwtService,
			mailService: mockMailService,
			ctx:         params.ctx,
			client:      params.client,
		}
	}

	return nil
}
