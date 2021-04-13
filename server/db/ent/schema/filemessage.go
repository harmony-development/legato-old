package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
)

// FileMessage holds the schema definition for the FileMessage entity.
type FileMessage struct {
	ent.Schema
}

// Fields of the FileMessage.
func (FileMessage) Fields() []ent.Field {
	return nil
}

// Edges of the FileMessage.
func (FileMessage) Edges() []ent.Edge {
	return []ent.Edge{
		edge.
			To("file", File.Type),
	}
}
