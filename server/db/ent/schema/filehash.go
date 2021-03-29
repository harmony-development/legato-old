package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// FileHash holds the schema definition for the FileHash entity.
type FileHash struct {
	ent.Schema
}

// Fields of the FileHash.
func (FileHash) Fields() []ent.Field {
	return []ent.Field{
		field.String("hash").Unique(),
	}
}

// Edges of the FileHash.
func (FileHash) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("file", File.Type).Unique(),
	}
}
