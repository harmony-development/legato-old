package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	harmonytypesv1 "github.com/harmony-development/legato/gen/harmonytypes/v1"
)

// Message holds the schema definition for the Message entity.
type Message struct {
	ent.Schema
}

// Fields of the Message.
func (Message) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Unique(),
		field.Time("createdat").Default(func() time.Time {
			return time.Now()
		}),
		field.Time("editedat").Optional(),
		field.JSON("actions", []*harmonytypesv1.Action{}).Optional(),
		field.JSON("metadata", &harmonytypesv1.Metadata{}).Optional(),
		field.Bytes("overrides").Optional(),
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
			To("replies", Message.Type).
			From("parent").
			Unique(),
		edge.
			To("text_message", TextMessage.Type).
			Unique(),
		edge.
			To("file_message", FileMessage.Type).
			Unique(),
		edge.
			To("embed_message", EmbedMessage.Type).
			Unique(),
	}
}

func (Message) Index() []ent.Index {
	return []ent.Index{
		index.Fields("createdat"),
	}
}
