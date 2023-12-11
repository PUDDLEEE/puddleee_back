package userdto

import "net/mail"

type UpdateUserDTO struct {
	Username *string
	Email    *string
	Password *string
}

func (c *UpdateUserDTO) Validate() error {
	switch {
	case c.Email != nil:
		if _, err := mail.ParseAddress(*c.Email); err != nil {
			return err
		}
	}
	return nil
}
