package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// ForeignUser holds the schema definition for the ForeignUser entity.
type ForeignUser struct {
	ent.Schema
}

// Fields of the ForeignUser.
func (ForeignUser) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("foreignid"),
		field.String("host"),
	}
}

// Edges of the ForeignUser.
func (ForeignUser) Edges() []ent.Edge {
	return []ent.Edge{
		edge.
			From("user", User.Type).
			Ref("foreign_user").
			Unique().
			Required(),
	}
}
