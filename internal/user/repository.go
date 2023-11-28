package user

import (
	"context"
	"fmt"
	"log"

	"github.com/PUDDLEEE/puddleee_back/ent"
	"github.com/PUDDLEEE/puddleee_back/ent/user"
)

type userRepository struct{}

func (u *userRepository) Create(ctx context.Context, client *ent.Client) (*ent.User, error) {
	user, err := client.User.
		Create().
		SetUsername("Something").
		SetEmail("123@naver.com").
		SetPassword("123").
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating user: %w", err)
	}
	log.Println("user was created: ", user)
	return user, nil
}

func (u *userRepository) FindOneById(ctx context.Context, client *ent.Client, id int) (*ent.User, error) {
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
