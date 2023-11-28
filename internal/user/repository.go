package user

import (
	"context"
	"fmt"
	"log"

	"github.com/PUDDLEEE/puddleee_back/ent"
	"github.com/PUDDLEEE/puddleee_back/ent/user"
	"github.com/PUDDLEEE/puddleee_back/internal/user/dto"
)

type userRepository struct{}

func (u *userRepository) create(ctx context.Context, client *ent.Client, dto dto.CreateUserDTO) (*ent.User, error) {
	user, err := client.User.Create().
		SetUsername(dto.Username).
		SetEmail(dto.Email).
		SetPassword(dto.Password).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating user: %w", err)
	}
	log.Println("user was created: ", user)
	return user, nil
}

func (u *userRepository) findOneById(ctx context.Context, client *ent.Client, id int) (*ent.User, error) {
	user, err := client.User.Query().
		Where(user.ID(id)).
		Only(ctx)
	if err != nil {
		return nil, err
	}
	log.Println("user found: ", user)
	return user, nil
}

func NewUserRepository() userRepository {
	return userRepository{}
}
