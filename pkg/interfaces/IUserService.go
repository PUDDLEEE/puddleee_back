package interfaces

import (
	"github.com/PUDDLEEE/puddleee_back/ent"
	"github.com/PUDDLEEE/puddleee_back/internal/user/dto"
)

type IUserService interface {
	CreateUser(dto.CreateUserDTO) (*ent.User, error)
	FindOneUser(int) (*ent.User, error)
	UpdateUser(int, dto.UpdateUserDTO) (*ent.User, error)
	DeleteUser(int) error
}

//mockery --dir=pkg/interfaces --name=IUserService --filename=IUserService.go --output=pkg/interfaces/mocks --outpkg=mocks
