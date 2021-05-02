package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	harmonytypesv1 "github.com/harmony-development/legato/gen/harmonytypes/v1"
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
		field.JSON("metadata", &harmonytypesv1.Metadata{}),
	}
}

// Edges of the Channel.
func (Channel) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("guild", Guild.Type).Ref("channel").Unique(),
		edge.To("message", Message.Type),
		edge.To("role", Role.Type),
		edge.To("permission_node", PermissionNode.Type),
	}
}
