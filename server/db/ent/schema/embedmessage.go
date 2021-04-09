package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// EmbedMessage holds the schema definition for the EmbedMessage entity.
type EmbedMessage struct {
	ent.Schema
}

// Fields of the EmbedMessage.
func (EmbedMessage) Fields() []ent.Field {
	return []ent.Field{
		field.String("title"),
		field.String("body"),
		field.Int64("color"),
		field.String("header_text"),
		field.String("header_subtext"),
		field.String("header_url"),
		field.String("header_icon"),
		field.String("footer_text"),
		field.String("footer_subtext"),
		field.String("footer_url"),
		field.String("footer_icon"),
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
