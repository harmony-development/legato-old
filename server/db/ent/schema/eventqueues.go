package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Message holds the schema definition for the Message entity.
type Host struct {
	ent.Schema
}

// Fields of the Message.
func (Host) Fields() []ent.Field {
	return []ent.Field{
		field.String("host").Unique(),
		field.Bytes("eventqueue"),
	}
}

// Edges of the Message.
func (Host) Edges() []ent.Edge {
	return []ent.Edge{}
}
