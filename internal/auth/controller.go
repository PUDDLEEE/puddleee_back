package auth

import (
	"github.com/PUDDLEEE/puddleee_back/pkg/interfaces"
	"github.com/PUDDLEEE/puddleee_back/pkg/interfaces/mocks"
	"github.com/gin-gonic/gin"
)

type AuthController struct {
	authService interfaces.IAuthService
}

// Signin godoc
//
//	@Summary	Sign in
//	@Schemes
//	@Description	Sign in to our service.
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			SigninDTO	body		authdto.SigninDTO	true	"Sign-in info"
//	@Success		200			{object}	authdto.SigninOutputDTO
//	@Failure		400			{object}	errors.CustomError
//	@Failure		404			{object}	errors.CustomError
//	@Failure		500			{object}	errors.CustomError
//	@Router			/signin [post]
func (c *AuthController) Signin(ctx *gin.Context) {

}

// Signup godoc
//
//	@Summary	Sign up
//	@Schemes
//	@Description	Sign up to our service.
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			createUserDTO	body		userdto.CreateUserDTO	true	"Sign-up info"
//	@Success		200				{object}	ent.User
//	@Failure		400				{object}	errors.CustomError
//	@Failure		404				{object}	errors.CustomError
//	@Failure		500				{object}	errors.CustomError
//	@Router			/signup [post]
func (c *AuthController) Signup(ctx *gin.Context) {

}

// SendConfirmationMail godoc
//
//	@Summary	Send Confirmation Mail to user.
//	@Schemes
//	@Description	Send Confirmation Mail to user which attached with confirmation code.
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	ent.Verification
//	@Failure		400	{object}	errors.CustomError
//	@Failure		404	{object}	errors.CustomError
//	@Failure		500	{object}	errors.CustomError
//	@Router			/mail [post]
func (c *AuthController) SendConfirmationMail(ctx *gin.Context) {

}

// VerifyConfirmationMail godoc
//
//	@Summary	Verify code from the user.
//	@Schemes
//	@Description	Verify code from the user.
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	ent.Verification
//	@Failure		400	{object}	errors.CustomError
//	@Failure		404	{object}	errors.CustomError
//	@Failure		500	{object}	errors.CustomError
//	@Router			/confirm [post]
func (c *AuthController) VerifyConfirmationMail(ctx *gin.Context) {

}
func NewController(service interfaces.IAuthService) *AuthController {
	if authService, ok := service.(*AuthService); ok {
		return &AuthController{
			authService: authService,
		}
	}

	if mockAuthService, ok := service.(*mocks.IAuthService); ok {
		return &AuthController{
			authService: mockAuthService,
		}
	}

	return nil
}
