package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// ActionInput holds the schema definition for the ActionInput entity.
type ActionInput struct {
	ent.Schema
}

// Fields of the ActionInput.
func (ActionInput) Fields() []ent.Field {
	return []ent.Field{
		field.String("label"),
		field.Bool("wide"),
	}
}

// Edges of the ActionInput.
func (ActionInput) Edges() []ent.Edge {
	return nil
}
