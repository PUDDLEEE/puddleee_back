package user

import (
	"context"

	"github.com/PUDDLEEE/puddleee_back/ent"
)

func InitializeController(ctx context.Context, client *ent.Client) UserController {
	userRepository := NewUserRepository()
	userService := NewService(userRepository, ctx, client)
	userController := NewController(userService)
	return userController
}
