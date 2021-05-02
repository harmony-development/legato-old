package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// EmotePack holds the schema definition for the EmotePack entity.
type EmotePack struct {
	ent.Schema
}

// Fields of the EmotePack.
func (EmotePack) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Unique(),
		field.String("name"),
	}
}

// Edges of the EmotePack.
func (EmotePack) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("emotepack").Unique(),
		edge.From("owner", User.Type).Ref("createdpacks").Unique(),
		edge.To("emote", Emote.Type),
	}
}
