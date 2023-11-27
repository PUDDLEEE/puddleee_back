package user

import (
	"context"
	"fmt"
	"log"

	"github.com/PUDDLEEE/puddleee_back/ent"
)

type userRepository struct{}

func (u *userRepository) create(ctx context.Context, client *ent.Client) (*ent.User, error) {
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

func NewUserRepository() userRepository {
	return userRepository{}
}
