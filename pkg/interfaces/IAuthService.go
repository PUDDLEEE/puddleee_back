package interfaces

import (
	"github.com/PUDDLEEE/puddleee_back/ent"
	authdto "github.com/PUDDLEEE/puddleee_back/internal/auth/dto"
	userdto "github.com/PUDDLEEE/puddleee_back/internal/user/dto"
)

type IAuthService interface {
	Signin(authdto.SigninDTO) (*authdto.SigninOutputDTO, error)
	Signup(userdto.CreateUserDTO) (*ent.User, error)
	CreateEmailVerification(string) (*ent.Verification, error)
	SendEmailVerification(*ent.Verification) error
	VerifyEmailVerification(string, string) error
}
