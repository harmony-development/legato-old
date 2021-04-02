package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// TextMessage holds the schema definition for the TextMessage entity.
type TextMessage struct {
	ent.Schema
}

// Fields of the TextMessage.
func (TextMessage) Fields() []ent.Field {
	return []ent.Field{
		field.String("content"),
	}
}

// Edges of the TextMessage.
func (TextMessage) Edges() []ent.Edge {
	return []ent.Edge{
		edge.
			From("message", Message.Type).
			Ref("textmessage").
			Unique(),
	}
}
