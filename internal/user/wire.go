//go:build wireinject
// +build wireinject

package user

import (
	"context"

	"github.com/PUDDLEEE/puddleee_back/ent"
	"github.com/google/wire"
)

func InitializeController(ctx context.Context, client *ent.Client) UserController {
	wire.Build(NewController, NewService, NewUserRepository)
	return UserController{}
}
