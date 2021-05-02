package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// FileHash holds the schema definition for the FileHash entity.
type FileHash struct {
	ent.Schema
}

// Fields of the FileHash.
func (FileHash) Fields() []ent.Field {
	return []ent.Field{
		field.Bytes("hash"),
		field.String("fileid"),
	}
}

func (FileHash) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("hash").Unique(),
	}
}

// Edges of the FileHash.
func (FileHash) Edges() []ent.Edge {
	return nil
}
