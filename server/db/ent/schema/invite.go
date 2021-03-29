package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Invite holds the schema definition for the Invite entity.
type Invite struct {
	ent.Schema
}

// Fields of the Invites.
func (Invite) Fields() []ent.Field {
	return []ent.Field{
		field.String("code").Unique(),
		field.Int64("uses").Default(0),
		field.Int64("possible_uses").Default(-1),
	}
}

// Edges of the Invites.
func (Invite) Edges() []ent.Edge {
	return []ent.Edge{
		edge.
			From("guild", Guild.Type).
			Ref("invite").
			Unique(),
	}
}
