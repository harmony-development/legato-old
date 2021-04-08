package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// PermissionNode holds the schema definition for the PermissionNode entity.
type PermissionNode struct {
	ent.Schema
}

// Fields of the PermissionNode.
func (PermissionNode) Fields() []ent.Field {
	return []ent.Field{
		field.String("node"),
		field.Bool("allow"),
	}
}

// Edges of the PermissionNode.
func (PermissionNode) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("role", Role.Type).Ref("permission_node").Unique(),
		edge.From("guild", Guild.Type).Ref("permission_node").Unique(),
		edge.From("channel", Channel.Type).Ref("permission_node").Unique(),
	}
}
