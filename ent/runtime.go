// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/PUDDLEEE/puddleee_back/ent/room"
	"github.com/PUDDLEEE/puddleee_back/ent/schema"
	"github.com/PUDDLEEE/puddleee_back/ent/user"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	roomFields := schema.Room{}.Fields()
	_ = roomFields
	// roomDescTitle is the schema descriptor for title field.
	roomDescTitle := roomFields[0].Descriptor()
	// room.TitleValidator is a validator for the "title" field. It is called by the builders before save.
	room.TitleValidator = roomDescTitle.Validators[0].(func(string) error)
	// roomDescIsCompleted is the schema descriptor for is_completed field.
	roomDescIsCompleted := roomFields[1].Descriptor()
	// room.DefaultIsCompleted holds the default value on creation for the is_completed field.
	room.DefaultIsCompleted = roomDescIsCompleted.Default.(bool)
	userMixin := schema.User{}.Mixin()
	userMixinFields0 := userMixin[0].Fields()
	_ = userMixinFields0
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescCreateTime is the schema descriptor for create_time field.
	userDescCreateTime := userMixinFields0[0].Descriptor()
	// user.DefaultCreateTime holds the default value on creation for the create_time field.
	user.DefaultCreateTime = userDescCreateTime.Default.(func() time.Time)
	// userDescUpdateTime is the schema descriptor for update_time field.
	userDescUpdateTime := userMixinFields0[1].Descriptor()
	// user.DefaultUpdateTime holds the default value on creation for the update_time field.
	user.DefaultUpdateTime = userDescUpdateTime.Default.(func() time.Time)
	// user.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	user.UpdateDefaultUpdateTime = userDescUpdateTime.UpdateDefault.(func() time.Time)
	// userDescUsername is the schema descriptor for username field.
	userDescUsername := userFields[0].Descriptor()
	// user.UsernameValidator is a validator for the "username" field. It is called by the builders before save.
	user.UsernameValidator = userDescUsername.Validators[0].(func(string) error)
	// userDescEmail is the schema descriptor for email field.
	userDescEmail := userFields[1].Descriptor()
	// user.EmailValidator is a validator for the "email" field. It is called by the builders before save.
	user.EmailValidator = userDescEmail.Validators[0].(func(string) error)
	// userDescPassword is the schema descriptor for password field.
	userDescPassword := userFields[3].Descriptor()
	// user.PasswordValidator is a validator for the "password" field. It is called by the builders before save.
	user.PasswordValidator = userDescPassword.Validators[0].(func(string) error)
}
