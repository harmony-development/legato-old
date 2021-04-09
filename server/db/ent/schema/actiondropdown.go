package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// ActionDropdown holds the schema definition for the ActionDropdown entity.
type ActionDropdown struct {
	ent.Schema
}

// Fields of the ActionDropdown.
func (ActionDropdown) Fields() []ent.Field {
	return []ent.Field{
		field.String("text"),
		field.Strings("options"),
	}
}

// Edges of the ActionDropdown.
func (ActionDropdown) Edges() []ent.Edge {
	return nil
}
