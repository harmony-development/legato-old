package schema

import "entgo.io/ent"

// EmbedMessage holds the schema definition for the EmbedMessage entity.
type EmbedMessage struct {
	ent.Schema
}

// Fields of the EmbedMessage.
func (EmbedMessage) Fields() []ent.Field {
	return nil
}

// Edges of the EmbedMessage.
func (EmbedMessage) Edges() []ent.Edge {
	return nil
}
