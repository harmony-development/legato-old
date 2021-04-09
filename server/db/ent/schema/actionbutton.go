package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// ActionButton holds the schema definition for the ActionButton entity.
type ActionButton struct {
	ent.Schema
}

// Fields of the ActionButton.
func (ActionButton) Fields() []ent.Field {
	return []ent.Field{
		field.String("text"),
		field.String("url"),
	}
}

// Edges of the ActionButton.
func (ActionButton) Edges() []ent.Edge {
	return nil
}
