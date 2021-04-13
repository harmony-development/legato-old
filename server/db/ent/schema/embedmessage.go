package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	harmonytypesv1 "github.com/harmony-development/legato/gen/harmonytypes/v1"
)

// EmbedMessage holds the schema definition for the EmbedMessage entity.
type EmbedMessage struct {
	ent.Schema
}

// Fields of the EmbedMessage.
func (EmbedMessage) Fields() []ent.Field {
	return []ent.Field{
		field.JSON("data", &harmonytypesv1.Embed{}),
	}
}

// Edges of the EmbedMessage.
func (EmbedMessage) Edges() []ent.Edge {
	return []ent.Edge{
		edge.
			To("embed_field", EmbedField.Type),
		edge.
			From("message", Message.Type).
			Ref("embed_message").
			Unique(),
	}
}
