package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// EmbedAction holds the schema definition for the EmbedAction entity.
type EmbedAction struct {
	ent.Schema
}

// Fields of the EmbedAction.
func (EmbedAction) Fields() []ent.Field {
	return []ent.Field{
		field.String("action_id"),
		field.Int8("action_type"),
	}
}

// Edges of the EmbedAction.
func (EmbedAction) Edges() []ent.Edge {
	return []ent.Edge{
		edge.
			To("button", ActionButton.Type).
			Unique(),
		edge.
			To("dropdown", ActionDropdown.Type).
			Unique(),
		edge.
			To("input", ActionInput.Type).
			Unique(),
	}
}
