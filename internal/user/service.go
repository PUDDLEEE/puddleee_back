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

func (s *UserService) CreateUser(dto dto.CreateUserDTO) (*ent.User, error) {
	user, err := s.userRepository.Create(s.ctx, s.client, dto)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserService) FindOneUser(id int) (*ent.User, error) {
	user, err := s.userRepository.FindOneById(s.ctx, s.client, id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserService) UpdateUser(id int, dto dto.UpdateUserDTO) (*ent.User, error) {
	user, err := s.userRepository.Update(s.ctx, s.client, id, dto)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserService) DeleteUser(id int) error {
	err := s.userRepository.Delete(s.ctx, s.client, id)
	if err != nil {
		return err
	}
	return nil
}

func NewService(repo interfaces.IUserRepository, ctx context.Context, client *ent.Client) *UserService {
	if userRepository, ok := repo.(*UserRepository); ok {
		return &UserService{userRepository: userRepository, ctx: ctx, client: client}
	}

	if userRepository, ok := repo.(*mocks.IUserRepository); ok {
		return &UserService{userRepository: userRepository, ctx: ctx, client: client}
	}
	return nil
}
