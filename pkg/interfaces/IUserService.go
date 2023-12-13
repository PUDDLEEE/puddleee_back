package interfaces

import (
	"github.com/PUDDLEEE/puddleee_back/ent"
	userdto "github.com/PUDDLEEE/puddleee_back/internal/user/dto"
)

type IUserService interface {
	CreateUser(userdto.CreateUserDTO) (*ent.User, error)
	FindOneUserById(int) (*ent.User, error)
	FindOneUserByEmail(string) (*ent.User, error)
	UpdateUser(int, userdto.UpdateUserDTO) (*ent.User, error)
	DeleteUser(int) error
}

//mockery --dir=pkg/interfaces --name=IUserService --filename=IUserService.go --output=pkg/interfaces/mocks --outpkg=mocks
