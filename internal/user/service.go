package user

import (
	"context"

	"github.com/PUDDLEEE/puddleee_back/ent"
	"github.com/PUDDLEEE/puddleee_back/internal/user/dto"
)

type UserService struct {
	userRepository userRepository
	ctx            context.Context
	client         *ent.Client
}

func (s *UserService) createUser(dto dto.CreateUserDTO) (*ent.User, error) {
	user, err := s.userRepository.create(s.ctx, s.client, dto)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserService) viewOneUser(id int) (*ent.User, error) {
	user, err := s.userRepository.findOneById(s.ctx, s.client, 1)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func NewService(repo userRepository, ctx context.Context, client *ent.Client) UserService {
	return UserService{userRepository: repo, ctx: ctx, client: client}
}
