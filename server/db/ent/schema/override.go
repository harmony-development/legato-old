package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Override holds the schema definition for the Override entity.
type Override struct {
	ent.Schema
}

// Fields of the Override.
func (Override) Fields() []ent.Field {
	return []ent.Field{
		field.String("username"),
		field.String("avatar"),
		field.Int64("reason"),
	}
}

// Edges of the Override.
func (Override) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("message", Message.Type).Ref("override"),
	}
}
