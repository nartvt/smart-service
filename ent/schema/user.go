package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Unique(),
		field.String("username").NotEmpty().Unique(),
		field.String("email").NotEmpty().Unique(),
		field.String("avatar").Optional(),
		field.String("birth_day").Optional(),
		field.String("first_name").Optional(),
		field.String("last_name").Optional(),
		field.String("phone_number").Optional(),
		field.Int("gender").Optional(),
		field.String("status"),
		field.String("address").Optional(),
		field.String("street").Optional(),
		field.String("district_id").Optional(),
		field.String("city_id").Optional(),
		field.String("country_code").Optional(),
		field.String("hash_password").NotEmpty(),
		field.Bool("verified").Default(false),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
