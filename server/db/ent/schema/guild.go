package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	harmonytypesv1 "github.com/harmony-development/legato/gen/harmonytypes/v1"
)

// Guild holds the schema definition for the Guild entity.
type Guild struct {
	ent.Schema
}

// Fields of the Guild.
func (Guild) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Unique(),
		field.String("name"),
		field.String("picture"),
		field.Bytes("metadata").GoType(&harmonytypesv1.Metadata{}),
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
			To("role", Role.Type),
		edge.
			To("permission_node", PermissionNode.Type),
		edge.
			To("owner", User.Type).Required().Unique(),
		edge.
			From("user", User.Type).Ref("guild"),
	}
}
