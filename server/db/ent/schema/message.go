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
		field.Time("createdat"),
		field.Time("editedat"),
	}
}

// Edges of the Message.
func (Message) Edges() []ent.Edge {
	return []ent.Edge{
		edge.
			From("user", User.Type).
			Ref("message").
			Unique(),
		edge.
			From("channel", Channel.Type).
			Ref("message").
			Unique(),
		edge.
			To("override", Override.Type).
			Unique(),
		edge.
			To("replies", Message.Type).
			From("parent").
			Unique(),
		edge.
			To("textmessage", TextMessage.Type).
			Unique(),
		edge.
			To("filemessage", FileMessage.Type).
			Unique(),
		edge.
			To("embedmessage", EmbedMessage.Type).
			Unique(),
	}
}

func (Message) Index() []ent.Index {
	return []ent.Index{
		index.Fields("createdat"),
	}
}
