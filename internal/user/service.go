package user

import (
	"context"

	"github.com/PUDDLEEE/puddleee_back/ent"
)

type UserService struct {
	userRepository userRepository
	ctx            context.Context
	client         *ent.Client
}

func (s *UserService) createUser() (*ent.User, error) {
	user, err := s.userRepository.Create(s.ctx, s.client)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserService) viewOneUser(id int) (*ent.User, error) {
	user, err := s.userRepository.FindOneById(s.ctx, s.client, 1)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func NewService(repo userRepository, ctx context.Context, client *ent.Client) UserService {
	return UserService{userRepository: repo}
}
