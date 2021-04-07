package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// UserMeta holds the schema definition for the UserMeta entity.
type UserMeta struct {
	ent.Schema
}

// Fields of the UserMeta.
func (UserMeta) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").Unique(),
		field.String("meta"),
	}
}

// Edges of the UserMeta.
func (UserMeta) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("metadata").Unique(),
	}
}
