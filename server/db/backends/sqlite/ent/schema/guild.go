package schema

import "entgo.io/ent"

// Guild holds the schema definition for the Guild entity.
type Guild struct {
	ent.Schema
}

// Fields of the Guild.
func (Guild) Fields() []ent.Field {
	return nil
}

// Edges of the Guild.
func (Guild) Edges() []ent.Edge {
	return nil
}
