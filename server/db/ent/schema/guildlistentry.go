package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// GuildListEntry holds the schema definition for the GuildListEntry entity.
type GuildListEntry struct {
	ent.Schema
}

// Fields of the GuildListEntry.
func (GuildListEntry) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id"),
		field.String("host"),
		field.String("position"),
	}
}

// Edges of the GuildListEntry.
func (GuildListEntry) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("listentry").Unique().Required(),
	}
}
