package schema

import (
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
		field.UUID(
			"id",
			uuid.UUID{},
		).Unique().Default(uuid.New).Immutable(),
		field.String("email").Unique().MaxLen(150).NotEmpty(),
		field.String("first_name").MaxLen(100).NotEmpty(),
		field.String("last_name").MaxLen(100).NotEmpty(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
