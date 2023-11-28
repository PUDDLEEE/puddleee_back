package user

import (
	"context"
	"fmt"
	"github.com/PUDDLEEE/puddleee_back/ent"
	"github.com/PUDDLEEE/puddleee_back/ent/user"
	"github.com/PUDDLEEE/puddleee_back/internal/user/dto"
)

type UserRepository struct{}

func (u *UserRepository) create(ctx context.Context, client *ent.Client, dto dto.CreateUserDTO) (*ent.User, error) {
	newUser, err := client.User.Create().
		SetUsername(dto.Username).
		SetEmail(dto.Email).
		SetPassword(dto.Password).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating user: %w", err)
	}
	return newUser, nil
}

func (u *UserRepository) findOneById(ctx context.Context, client *ent.Client, id int) (*ent.User, error) {
	existedUser, err := client.User.Query().
		Where(user.ID(id)).
		Select(user.FieldID, user.FieldEmail, user.FieldUsername).
		Only(ctx)
	if err != nil {
		return nil, err
	}
	return existedUser, nil
}

func NewUserRepository() UserRepository {
	return UserRepository{}
}
