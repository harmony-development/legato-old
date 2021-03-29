package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Emote holds the schema definition for the Emote entity.
type Emote struct {
	ent.Schema
}

// Fields of the Emote.
func (Emote) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
	}
}

// Edges of the Emote.
func (Emote) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("emotepack", EmotePack.Type).Ref("emote").Unique(),
		edge.To("file", File.Type).Unique(),
	}
}
