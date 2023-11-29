package dto

import "net/mail"

type CreateUserDTO struct {
	Username string
	Email    string
	Password string
}

func (c *CreateUserDTO) Validate() error {
	if _, err := mail.ParseAddress(c.Email); err != nil {
		return err
	}
	return nil
}
