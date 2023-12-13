package auth

import (
	"context"

	"github.com/PUDDLEEE/puddleee_back/ent"
	"github.com/PUDDLEEE/puddleee_back/internal/jwtAuth"
	"github.com/PUDDLEEE/puddleee_back/internal/mail"
	"github.com/PUDDLEEE/puddleee_back/internal/user"
)

func InitializeController(ctx context.Context, client *ent.Client, key string, mailParams mail.MailServiceParams) *AuthController {
	userRepository := NewAuthRepository()
	params := &AuthServiceParameter{}
	userRepo := user.NewUserRepository()

	params.userService = user.NewService(userRepo, ctx, client)
	params.authRepo = userRepository
	params.jwtService = jwtAuth.NewJwtAuthService(key)
	params.mailService = mail.NewMailService(mailParams)
	params.client = client
	params.ctx = ctx

	userService := NewAuthService(params)
	userController := NewController(userService)
	return userController
}
