package user

import (
	"context"
	"errors"

	userdto "github.com/PUDDLEEE/puddleee_back/internal/user/dto"

	"github.com/PUDDLEEE/puddleee_back/ent"
	"github.com/PUDDLEEE/puddleee_back/ent/user"
)

type UserRepository struct{}

func (u *UserRepository) Create(ctx context.Context, client *ent.Client, dto userdto.CreateUserDTO) (*ent.User, error) {
	newUser, err := client.User.Create().
		SetUsername(dto.Username).
		SetEmail(dto.Email).
		SetPassword(dto.Password).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return newUser, nil
}

func (u *UserRepository) FindOneById(ctx context.Context, client *ent.Client, id int) (*ent.User, error) {
	existedUser, err := client.User.Query().
		Where(user.ID(id)).
		Select(user.FieldID, user.FieldEmail, user.FieldUsername).
		Only(ctx)
	if err != nil {
		return nil, err
	}
	return existedUser, nil
}

func (u *UserRepository) FindOneByEmail(ctx context.Context, client *ent.Client, email string) (*ent.User, error) {
	existedUser, err := client.User.Query().Where(user.Email(email)).
		Select(user.FieldID, user.FieldEmail, user.FieldUsername).
		Only(ctx)
	if err != nil {
		return nil, err
	}
	return existedUser, nil
}
func (u *UserRepository) Update(ctx context.Context, client *ent.Client, id int, dto userdto.UpdateUserDTO) (*ent.User, error) {
	if dto.Username == nil && dto.Email == nil && dto.Password == nil {
		err := errors.New("at least one field should contain data")
		return nil, err
	}
	updatedUser, err := client.User.
		UpdateOneID(id).
		SetNillableUsername(dto.Username).
		SetNillableEmail(dto.Email).
		SetNillablePassword(dto.Password).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return updatedUser, nil
}

func (u *UserRepository) Delete(ctx context.Context, client *ent.Client, id int) error {
	err := client.User.
		DeleteOneID(id).
		Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}
