package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// EmbedField holds the schema definition for the EmbedField entity.
type EmbedField struct {
	ent.Schema
}

// Fields of the EmbedField.
func (EmbedField) Fields() []ent.Field {
	return []ent.Field{
		field.String("title"),
		field.String("subtitle"),
		field.String("body"),
		field.String("image_url"),
		field.Int8("presentation"),
	}
}

// Edges of the EmbedField.
func (EmbedField) Edges() []ent.Edge {
	return []ent.Edge{
		edge.
			To("embed_action", EmbedAction.Type),
		edge.
			From("embed_message", EmbedMessage.Type).
			Ref("embed_field").
			Unique(),
	}
}
