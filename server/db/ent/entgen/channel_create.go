// Code generated by entc, DO NOT EDIT.

package entgen

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/harmony-development/legato/server/db/ent/entgen/channel"
	"github.com/harmony-development/legato/server/db/ent/entgen/guild"
	"github.com/harmony-development/legato/server/db/ent/entgen/message"
	"github.com/harmony-development/legato/server/db/ent/entgen/permissionnode"
	"github.com/harmony-development/legato/server/db/ent/entgen/role"
)

// ChannelCreate is the builder for creating a Channel entity.
type ChannelCreate struct {
	config
	mutation *ChannelMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (cc *ChannelCreate) SetName(s string) *ChannelCreate {
	cc.mutation.SetName(s)
	return cc
}

// SetKind sets the "kind" field.
func (cc *ChannelCreate) SetKind(u uint64) *ChannelCreate {
	cc.mutation.SetKind(u)
	return cc
}

// SetPosition sets the "position" field.
func (cc *ChannelCreate) SetPosition(s string) *ChannelCreate {
	cc.mutation.SetPosition(s)
	return cc
}

// SetMetadata sets the "metadata" field.
func (cc *ChannelCreate) SetMetadata(b []byte) *ChannelCreate {
	cc.mutation.SetMetadata(b)
	return cc
}

// SetID sets the "id" field.
func (cc *ChannelCreate) SetID(u uint64) *ChannelCreate {
	cc.mutation.SetID(u)
	return cc
}

// SetGuildID sets the "guild" edge to the Guild entity by ID.
func (cc *ChannelCreate) SetGuildID(id uint64) *ChannelCreate {
	cc.mutation.SetGuildID(id)
	return cc
}

// SetNillableGuildID sets the "guild" edge to the Guild entity by ID if the given value is not nil.
func (cc *ChannelCreate) SetNillableGuildID(id *uint64) *ChannelCreate {
	if id != nil {
		cc = cc.SetGuildID(*id)
	}
	return cc
}

// SetGuild sets the "guild" edge to the Guild entity.
func (cc *ChannelCreate) SetGuild(g *Guild) *ChannelCreate {
	return cc.SetGuildID(g.ID)
}

// AddMessageIDs adds the "message" edge to the Message entity by IDs.
func (cc *ChannelCreate) AddMessageIDs(ids ...uint64) *ChannelCreate {
	cc.mutation.AddMessageIDs(ids...)
	return cc
}

// AddMessage adds the "message" edges to the Message entity.
func (cc *ChannelCreate) AddMessage(m ...*Message) *ChannelCreate {
	ids := make([]uint64, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return cc.AddMessageIDs(ids...)
}

// AddRoleIDs adds the "role" edge to the Role entity by IDs.
func (cc *ChannelCreate) AddRoleIDs(ids ...uint64) *ChannelCreate {
	cc.mutation.AddRoleIDs(ids...)
	return cc
}

// AddRole adds the "role" edges to the Role entity.
func (cc *ChannelCreate) AddRole(r ...*Role) *ChannelCreate {
	ids := make([]uint64, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return cc.AddRoleIDs(ids...)
}

// AddPermissionNodeIDs adds the "permission_node" edge to the PermissionNode entity by IDs.
func (cc *ChannelCreate) AddPermissionNodeIDs(ids ...int) *ChannelCreate {
	cc.mutation.AddPermissionNodeIDs(ids...)
	return cc
}

// AddPermissionNode adds the "permission_node" edges to the PermissionNode entity.
func (cc *ChannelCreate) AddPermissionNode(p ...*PermissionNode) *ChannelCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return cc.AddPermissionNodeIDs(ids...)
}

// Mutation returns the ChannelMutation object of the builder.
func (cc *ChannelCreate) Mutation() *ChannelMutation {
	return cc.mutation
}

// Save creates the Channel in the database.
func (cc *ChannelCreate) Save(ctx context.Context) (*Channel, error) {
	var (
		err  error
		node *Channel
	)
	if len(cc.hooks) == 0 {
		if err = cc.check(); err != nil {
			return nil, err
		}
		node, err = cc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ChannelMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = cc.check(); err != nil {
				return nil, err
			}
			cc.mutation = mutation
			node, err = cc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(cc.hooks) - 1; i >= 0; i-- {
			mut = cc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (cc *ChannelCreate) SaveX(ctx context.Context) *Channel {
	v, err := cc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// check runs all checks and user-defined validators on the builder.
func (cc *ChannelCreate) check() error {
	if _, ok := cc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New("entgen: missing required field \"name\"")}
	}
	if _, ok := cc.mutation.Kind(); !ok {
		return &ValidationError{Name: "kind", err: errors.New("entgen: missing required field \"kind\"")}
	}
	if _, ok := cc.mutation.Position(); !ok {
		return &ValidationError{Name: "position", err: errors.New("entgen: missing required field \"position\"")}
	}
	if _, ok := cc.mutation.Metadata(); !ok {
		return &ValidationError{Name: "metadata", err: errors.New("entgen: missing required field \"metadata\"")}
	}
	return nil
}

func (cc *ChannelCreate) sqlSave(ctx context.Context) (*Channel, error) {
	_node, _spec := cc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cc.driver, _spec); err != nil {
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

func (cc *ChannelCreate) createSpec() (*Channel, *sqlgraph.CreateSpec) {
	var (
		_node = &Channel{config: cc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: channel.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: channel.FieldID,
			},
		}
	)
	if id, ok := cc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := cc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: channel.FieldName,
		})
		_node.Name = value
	}
	if value, ok := cc.mutation.Kind(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: channel.FieldKind,
		})
		_node.Kind = value
	}
	if value, ok := cc.mutation.Position(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: channel.FieldPosition,
		})
		_node.Position = value
	}
	if value, ok := cc.mutation.Metadata(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBytes,
			Value:  value,
			Column: channel.FieldMetadata,
		})
		_node.Metadata = value
	}
	if nodes := cc.mutation.GuildIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   channel.GuildTable,
			Columns: []string{channel.GuildColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: guild.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.guild_channel = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cc.mutation.MessageIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   channel.MessageTable,
			Columns: []string{channel.MessageColumn},
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
	if nodes := cc.mutation.RoleIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   channel.RoleTable,
			Columns: []string{channel.RoleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: role.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cc.mutation.PermissionNodeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   channel.PermissionNodeTable,
			Columns: []string{channel.PermissionNodeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: permissionnode.FieldID,
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

// ChannelCreateBulk is the builder for creating many Channel entities in bulk.
type ChannelCreateBulk struct {
	config
	builders []*ChannelCreate
}

// Save creates the Channel entities in the database.
func (ccb *ChannelCreateBulk) Save(ctx context.Context) ([]*Channel, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ccb.builders))
	nodes := make([]*Channel, len(ccb.builders))
	mutators := make([]Mutator, len(ccb.builders))
	for i := range ccb.builders {
		func(i int, root context.Context) {
			builder := ccb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ChannelMutation)
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
					_, err = mutators[i+1].Mutate(root, ccb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ccb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ccb *ChannelCreateBulk) SaveX(ctx context.Context) []*Channel {
	v, err := ccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
