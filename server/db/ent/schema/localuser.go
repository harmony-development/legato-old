package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/badoux/checkmail"
)

// LocalUser holds the schema definition for the LocalUser entity.
type LocalUser struct {
	ent.Schema
}

// Fields of the LocalUser.
func (LocalUser) Fields() []ent.Field {
	return []ent.Field{
		field.String("email").Unique().Validate(checkmail.ValidateFormat),
		field.Bytes("password"),
	}
}

// Edges of the LocalUser.
func (LocalUser) Edges() []ent.Edge {
	return []ent.Edge{
		edge.
			From("user", User.Type).
			Ref("local_user").
			Unique().
			Required(),
		edge.
			To("sessions", Session.Type),
	}
}
