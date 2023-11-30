package user

import (
	"context"

	"github.com/PUDDLEEE/puddleee_back/ent"
	"github.com/PUDDLEEE/puddleee_back/internal/user/dto"
	"github.com/PUDDLEEE/puddleee_back/pkg/interfaces"
	"github.com/PUDDLEEE/puddleee_back/pkg/interfaces/mocks"
)

type UserService struct {
	userRepository interfaces.IUserRepository
	ctx            context.Context
	client         *ent.Client
}

func (s *UserService) createUser(dto dto.CreateUserDTO) (*ent.User, error) {
	user, err := s.userRepository.Create(s.ctx, s.client, dto)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserService) findOneUser(id int) (*ent.User, error) {
	user, err := s.userRepository.FindOneById(s.ctx, s.client, id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func NewService(repo interfaces.IUserRepository, ctx context.Context, client *ent.Client) UserService {
	if userRepository, ok := repo.(*UserRepository); ok {
		return UserService{userRepository: userRepository, ctx: ctx, client: client}
	}

	if userRepository, ok := repo.(*mocks.IUserRepository); ok {
		return UserService{userRepository: userRepository, ctx: ctx, client: client}
	}
	return UserService{}
}
