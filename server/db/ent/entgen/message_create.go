// Code generated by entc, DO NOT EDIT.

package entgen

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	v1 "github.com/harmony-development/legato/gen/harmonytypes/v1"
	"github.com/harmony-development/legato/server/db/ent/entgen/channel"
	"github.com/harmony-development/legato/server/db/ent/entgen/embedmessage"
	"github.com/harmony-development/legato/server/db/ent/entgen/filemessage"
	"github.com/harmony-development/legato/server/db/ent/entgen/message"
	"github.com/harmony-development/legato/server/db/ent/entgen/textmessage"
	"github.com/harmony-development/legato/server/db/ent/entgen/user"
)

// MessageCreate is the builder for creating a Message entity.
type MessageCreate struct {
	config
	mutation *MessageMutation
	hooks    []Hook
}

// SetCreatedat sets the "createdat" field.
func (mc *MessageCreate) SetCreatedat(t time.Time) *MessageCreate {
	mc.mutation.SetCreatedat(t)
	return mc
}

// SetNillableCreatedat sets the "createdat" field if the given value is not nil.
func (mc *MessageCreate) SetNillableCreatedat(t *time.Time) *MessageCreate {
	if t != nil {
		mc.SetCreatedat(*t)
	}
	return mc
}

// SetEditedat sets the "editedat" field.
func (mc *MessageCreate) SetEditedat(t time.Time) *MessageCreate {
	mc.mutation.SetEditedat(t)
	return mc
}

// SetNillableEditedat sets the "editedat" field if the given value is not nil.
func (mc *MessageCreate) SetNillableEditedat(t *time.Time) *MessageCreate {
	if t != nil {
		mc.SetEditedat(*t)
	}
	return mc
}

// SetMetadata sets the "metadata" field.
func (mc *MessageCreate) SetMetadata(v *v1.Metadata) *MessageCreate {
	mc.mutation.SetMetadata(v)
	return mc
}

// SetOverride sets the "override" field.
func (mc *MessageCreate) SetOverride(v *v1.Override) *MessageCreate {
	mc.mutation.SetOverride(v)
	return mc
}

// SetID sets the "id" field.
func (mc *MessageCreate) SetID(u uint64) *MessageCreate {
	mc.mutation.SetID(u)
	return mc
}

// SetUserID sets the "user" edge to the User entity by ID.
func (mc *MessageCreate) SetUserID(id uint64) *MessageCreate {
	mc.mutation.SetUserID(id)
	return mc
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (mc *MessageCreate) SetNillableUserID(id *uint64) *MessageCreate {
	if id != nil {
		mc = mc.SetUserID(*id)
	}
	return mc
}

// SetUser sets the "user" edge to the User entity.
func (mc *MessageCreate) SetUser(u *User) *MessageCreate {
	return mc.SetUserID(u.ID)
}

// SetChannelID sets the "channel" edge to the Channel entity by ID.
func (mc *MessageCreate) SetChannelID(id uint64) *MessageCreate {
	mc.mutation.SetChannelID(id)
	return mc
}

// SetNillableChannelID sets the "channel" edge to the Channel entity by ID if the given value is not nil.
func (mc *MessageCreate) SetNillableChannelID(id *uint64) *MessageCreate {
	if id != nil {
		mc = mc.SetChannelID(*id)
	}
	return mc
}

// SetChannel sets the "channel" edge to the Channel entity.
func (mc *MessageCreate) SetChannel(c *Channel) *MessageCreate {
	return mc.SetChannelID(c.ID)
}

// SetParentID sets the "parent" edge to the Message entity by ID.
func (mc *MessageCreate) SetParentID(id uint64) *MessageCreate {
	mc.mutation.SetParentID(id)
	return mc
}

// SetNillableParentID sets the "parent" edge to the Message entity by ID if the given value is not nil.
func (mc *MessageCreate) SetNillableParentID(id *uint64) *MessageCreate {
	if id != nil {
		mc = mc.SetParentID(*id)
	}
	return mc
}

// SetParent sets the "parent" edge to the Message entity.
func (mc *MessageCreate) SetParent(m *Message) *MessageCreate {
	return mc.SetParentID(m.ID)
}

// AddReplyIDs adds the "replies" edge to the Message entity by IDs.
func (mc *MessageCreate) AddReplyIDs(ids ...uint64) *MessageCreate {
	mc.mutation.AddReplyIDs(ids...)
	return mc
}

// AddReplies adds the "replies" edges to the Message entity.
func (mc *MessageCreate) AddReplies(m ...*Message) *MessageCreate {
	ids := make([]uint64, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return mc.AddReplyIDs(ids...)
}

// SetTextMessageID sets the "text_message" edge to the TextMessage entity by ID.
func (mc *MessageCreate) SetTextMessageID(id int) *MessageCreate {
	mc.mutation.SetTextMessageID(id)
	return mc
}

// SetNillableTextMessageID sets the "text_message" edge to the TextMessage entity by ID if the given value is not nil.
func (mc *MessageCreate) SetNillableTextMessageID(id *int) *MessageCreate {
	if id != nil {
		mc = mc.SetTextMessageID(*id)
	}
	return mc
}

// SetTextMessage sets the "text_message" edge to the TextMessage entity.
func (mc *MessageCreate) SetTextMessage(t *TextMessage) *MessageCreate {
	return mc.SetTextMessageID(t.ID)
}

// SetFileMessageID sets the "file_message" edge to the FileMessage entity by ID.
func (mc *MessageCreate) SetFileMessageID(id int) *MessageCreate {
	mc.mutation.SetFileMessageID(id)
	return mc
}

// SetNillableFileMessageID sets the "file_message" edge to the FileMessage entity by ID if the given value is not nil.
func (mc *MessageCreate) SetNillableFileMessageID(id *int) *MessageCreate {
	if id != nil {
		mc = mc.SetFileMessageID(*id)
	}
	return mc
}

// SetFileMessage sets the "file_message" edge to the FileMessage entity.
func (mc *MessageCreate) SetFileMessage(f *FileMessage) *MessageCreate {
	return mc.SetFileMessageID(f.ID)
}

// SetEmbedMessageID sets the "embed_message" edge to the EmbedMessage entity by ID.
func (mc *MessageCreate) SetEmbedMessageID(id int) *MessageCreate {
	mc.mutation.SetEmbedMessageID(id)
	return mc
}

// SetNillableEmbedMessageID sets the "embed_message" edge to the EmbedMessage entity by ID if the given value is not nil.
func (mc *MessageCreate) SetNillableEmbedMessageID(id *int) *MessageCreate {
	if id != nil {
		mc = mc.SetEmbedMessageID(*id)
	}
	return mc
}

// SetEmbedMessage sets the "embed_message" edge to the EmbedMessage entity.
func (mc *MessageCreate) SetEmbedMessage(e *EmbedMessage) *MessageCreate {
	return mc.SetEmbedMessageID(e.ID)
}

// Mutation returns the MessageMutation object of the builder.
func (mc *MessageCreate) Mutation() *MessageMutation {
	return mc.mutation
}

// Save creates the Message in the database.
func (mc *MessageCreate) Save(ctx context.Context) (*Message, error) {
	var (
		err  error
		node *Message
	)
	mc.defaults()
	if len(mc.hooks) == 0 {
		if err = mc.check(); err != nil {
			return nil, err
		}
		node, err = mc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MessageMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = mc.check(); err != nil {
				return nil, err
			}
			mc.mutation = mutation
			node, err = mc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(mc.hooks) - 1; i >= 0; i-- {
			mut = mc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, mc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (mc *MessageCreate) SaveX(ctx context.Context) *Message {
	v, err := mc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (mc *MessageCreate) defaults() {
	if _, ok := mc.mutation.Createdat(); !ok {
		v := message.DefaultCreatedat()
		mc.mutation.SetCreatedat(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mc *MessageCreate) check() error {
	if _, ok := mc.mutation.Createdat(); !ok {
		return &ValidationError{Name: "createdat", err: errors.New("entgen: missing required field \"createdat\"")}
	}
	return nil
}

func (mc *MessageCreate) sqlSave(ctx context.Context) (*Message, error) {
	_node, _spec := mc.createSpec()
	if err := sqlgraph.CreateNode(ctx, mc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	if _node.ID == 0 {
		id := _spec.ID.Value.(int64)
		_node.ID = uint64(id)
	}
	return _node, nil
}

func (mc *MessageCreate) createSpec() (*Message, *sqlgraph.CreateSpec) {
	var (
		_node = &Message{config: mc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: message.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: message.FieldID,
			},
		}
	)
	if id, ok := mc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := mc.mutation.Createdat(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: message.FieldCreatedat,
		})
		_node.Createdat = value
	}
	if value, ok := mc.mutation.Editedat(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: message.FieldEditedat,
		})
		_node.Editedat = value
	}
	if value, ok := mc.mutation.Metadata(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: message.FieldMetadata,
		})
		_node.Metadata = value
	}
	if value, ok := mc.mutation.Override(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: message.FieldOverride,
		})
		_node.Override = value
	}
	if nodes := mc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   message.UserTable,
			Columns: []string{message.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.user_message = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := mc.mutation.ChannelIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   message.ChannelTable,
			Columns: []string{message.ChannelColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: channel.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.channel_message = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := mc.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   message.ParentTable,
			Columns: []string{message.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: message.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.message_replies = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := mc.mutation.RepliesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   message.RepliesTable,
			Columns: []string{message.RepliesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: message.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := mc.mutation.TextMessageIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   message.TextMessageTable,
			Columns: []string{message.TextMessageColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: textmessage.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := mc.mutation.FileMessageIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   message.FileMessageTable,
			Columns: []string{message.FileMessageColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: filemessage.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.message_file_message = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := mc.mutation.EmbedMessageIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   message.EmbedMessageTable,
			Columns: []string{message.EmbedMessageColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: embedmessage.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// MessageCreateBulk is the builder for creating many Message entities in bulk.
type MessageCreateBulk struct {
	config
	builders []*MessageCreate
}

// Save creates the Message entities in the database.
func (mcb *MessageCreateBulk) Save(ctx context.Context) ([]*Message, error) {
	specs := make([]*sqlgraph.CreateSpec, len(mcb.builders))
	nodes := make([]*Message, len(mcb.builders))
	mutators := make([]Mutator, len(mcb.builders))
	for i := range mcb.builders {
		func(i int, root context.Context) {
			builder := mcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*MessageMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, mcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, mcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
						if cerr, ok := isSQLConstraintError(err); ok {
							err = cerr
						}
					}
				}
				mutation.done = true
				if err != nil {
					return nil, err
				}
				if nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = uint64(id)
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, mcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (mcb *MessageCreateBulk) SaveX(ctx context.Context) []*Message {
	v, err := mcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
