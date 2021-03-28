package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Message holds the schema definition for the Message entity.
type Message struct {
	ent.Schema
}

// Fields of the Message.
func (Message) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Unique(),
		field.Uint64("author"),
		field.Time("createdat"),
		field.Time("editedat"),
		field.Uint64("replyto").Default(0),
	}
}

// Edges of the Message.
func (Message) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("message").Unique(),
		edge.To("override", Override.Type),
	}
}

func (Message) Index() []ent.Index {
	return []ent.Index{
		index.Fields("createdat"),
	}
}