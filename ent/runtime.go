// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/google/uuid"
	"github.com/nartvt/smart-service/ent/address"
	"github.com/nartvt/smart-service/ent/schema"
	"github.com/nartvt/smart-service/ent/user"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	addressFields := schema.Address{}.Fields()
	_ = addressFields
	// addressDescCreatedAt is the schema descriptor for created_at field.
	addressDescCreatedAt := addressFields[5].Descriptor()
	// address.DefaultCreatedAt holds the default value on creation for the created_at field.
	address.DefaultCreatedAt = addressDescCreatedAt.Default.(func() time.Time)
	// addressDescUpdatedAt is the schema descriptor for updated_at field.
	addressDescUpdatedAt := addressFields[6].Descriptor()
	// address.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	address.DefaultUpdatedAt = addressDescUpdatedAt.Default.(func() time.Time)
	// address.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	address.UpdateDefaultUpdatedAt = addressDescUpdatedAt.UpdateDefault.(func() time.Time)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescUsername is the schema descriptor for username field.
	userDescUsername := userFields[1].Descriptor()
	// user.UsernameValidator is a validator for the "username" field. It is called by the builders before save.
	user.UsernameValidator = userDescUsername.Validators[0].(func(string) error)
	// userDescEmail is the schema descriptor for email field.
	userDescEmail := userFields[2].Descriptor()
	// user.EmailValidator is a validator for the "email" field. It is called by the builders before save.
	user.EmailValidator = userDescEmail.Validators[0].(func(string) error)
	// userDescHashPassword is the schema descriptor for hash_password field.
	userDescHashPassword := userFields[15].Descriptor()
	// user.HashPasswordValidator is a validator for the "hash_password" field. It is called by the builders before save.
	user.HashPasswordValidator = userDescHashPassword.Validators[0].(func(string) error)
	// userDescVerified is the schema descriptor for verified field.
	userDescVerified := userFields[16].Descriptor()
	// user.DefaultVerified holds the default value on creation for the verified field.
	user.DefaultVerified = userDescVerified.Default.(bool)
	// userDescCreatedAt is the schema descriptor for created_at field.
	userDescCreatedAt := userFields[17].Descriptor()
	// user.DefaultCreatedAt holds the default value on creation for the created_at field.
	user.DefaultCreatedAt = userDescCreatedAt.Default.(func() time.Time)
	// userDescUpdatedAt is the schema descriptor for updated_at field.
	userDescUpdatedAt := userFields[18].Descriptor()
	// user.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	user.DefaultUpdatedAt = userDescUpdatedAt.Default.(func() time.Time)
	// user.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	user.UpdateDefaultUpdatedAt = userDescUpdatedAt.UpdateDefault.(func() time.Time)
	// userDescID is the schema descriptor for id field.
	userDescID := userFields[0].Descriptor()
	// user.DefaultID holds the default value on creation for the id field.
	user.DefaultID = userDescID.Default.(func() uuid.UUID)
}
