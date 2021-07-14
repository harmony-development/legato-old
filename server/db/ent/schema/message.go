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
		field.
			Uint64("messageID").
			Unique(),
		field.
			Time("createdat").
			Default(
				func() time.Time {
					return time.Now()
				},
			),
		field.
			Time("editedat").
			Optional(),
		field.Bytes("metadata").GoType(&harmonytypesv1.Metadata{}).Optional(),
		field.Bytes("override").GoType(&harmonytypesv1.Override{}).Optional(),
		field.Bytes("content").GoType(&harmonytypesv1.Content{}),
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
	}
}

func (Message) Index() []ent.Index {
	return []ent.Index{
		index.Fields("createdat"),
	}
}
