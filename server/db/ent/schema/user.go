package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Unique(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.
			To("local_user", LocalUser.Type).
			Unique(),
		edge.
			To("foreign_user", ForeignUser.Type).
			Unique(),
		edge.
			To("profile", Profile.Type).
			Unique(),
		edge.
			To("metadata", UserMeta.Type),
		edge.
			To("sessions", Session.Type),
		edge.
			To("message", Message.Type),
		edge.
			To("guild", Guild.Type),
		edge.
			To("emotepack", EmotePack.Type),
		edge.To(
			"createdpacks", EmotePack.Type),
		edge.
			To("listentry", GuildListEntry.Type),
		edge.
			From("role", Role.Type).
			Ref("members"),
	}
}
