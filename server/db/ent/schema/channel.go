package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Channel holds the schema definition for the Channel entity.
type Channel struct {
	ent.Schema
}

// Fields of the Channel.
func (Channel) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id"),
		field.String("name"),
		field.Uint64("kind"),
		field.String("position"),
		field.Bytes("metadata"),
	}
}

// Edges of the Channel.
func (Channel) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("guild", Guild.Type).Ref("channel").Unique(),
		edge.To("message", Message.Type),
	}
}
