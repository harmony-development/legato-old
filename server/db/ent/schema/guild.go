package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Guild holds the schema definition for the Guild entity.
type Guild struct {
	ent.Schema
}

// Fields of the Guild.
func (Guild) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Unique(),
		field.Uint64("owner"),
		field.String("name"),
		field.String("picture"),
		field.Bytes("metadata"),
	}
}

// Edges of the Guild.
func (Guild) Edges() []ent.Edge {
	return []ent.Edge{
		edge.
			To("invite", Invite.Type),
		edge.
			To("bans", User.Type),
		edge.
			To("channel", Channel.Type),
		edge.
			From("user", User.Type).Ref("guild").Required(),
	}
}
