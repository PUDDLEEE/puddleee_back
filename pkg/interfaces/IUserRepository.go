package interfaces

import (
	"context"

	"github.com/PUDDLEEE/puddleee_back/ent"
	userdto "github.com/PUDDLEEE/puddleee_back/internal/user/dto"
)

type IUserRepository interface {
	Create(context.Context, *ent.Client, userdto.CreateUserDTO) (*ent.User, error)
	FindOneById(context.Context, *ent.Client, int) (*ent.User, error)
	FindOneByEmail(context.Context, *ent.Client, string) (*ent.User, error)
	Update(context.Context, *ent.Client, int, userdto.UpdateUserDTO) (*ent.User, error)
	Delete(context.Context, *ent.Client, int) error
}

//mockgen -source=pkg/interfaces/IUserRepository.go -destination=pkg/interfaces/mocks/IUserRepository.go -package=mocks
//mockery --dir=pkg/interfaces --name=IUserRepository --filename=IUserRepository.go --output=pkg/interfaces/mocks --outpkg=mocks
