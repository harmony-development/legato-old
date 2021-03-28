// Code generated by entc, DO NOT EDIT.

package entgen

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/harmony-development/legato/server/db/ent/entgen/guild"
	"github.com/harmony-development/legato/server/db/ent/entgen/invite"
	"github.com/harmony-development/legato/server/db/ent/entgen/predicate"
)

// InviteUpdate is the builder for updating Invite entities.
type InviteUpdate struct {
	config
	hooks    []Hook
	mutation *InviteMutation
}

// Where adds a new predicate for the InviteUpdate builder.
func (iu *InviteUpdate) Where(ps ...predicate.Invite) *InviteUpdate {
	iu.mutation.predicates = append(iu.mutation.predicates, ps...)
	return iu
}

// SetCode sets the "code" field.
func (iu *InviteUpdate) SetCode(s string) *InviteUpdate {
	iu.mutation.SetCode(s)
	return iu
}

// SetUses sets the "uses" field.
func (iu *InviteUpdate) SetUses(i int64) *InviteUpdate {
	iu.mutation.ResetUses()
	iu.mutation.SetUses(i)
	return iu
}

// SetNillableUses sets the "uses" field if the given value is not nil.
func (iu *InviteUpdate) SetNillableUses(i *int64) *InviteUpdate {
	if i != nil {
		iu.SetUses(*i)
	}
	return iu
}

// AddUses adds i to the "uses" field.
func (iu *InviteUpdate) AddUses(i int64) *InviteUpdate {
	iu.mutation.AddUses(i)
	return iu
}

// SetPossibleUses sets the "possible_uses" field.
func (iu *InviteUpdate) SetPossibleUses(i int64) *InviteUpdate {
	iu.mutation.ResetPossibleUses()
	iu.mutation.SetPossibleUses(i)
	return iu
}

// SetNillablePossibleUses sets the "possible_uses" field if the given value is not nil.
func (iu *InviteUpdate) SetNillablePossibleUses(i *int64) *InviteUpdate {
	if i != nil {
		iu.SetPossibleUses(*i)
	}
	return iu
}

// AddPossibleUses adds i to the "possible_uses" field.
func (iu *InviteUpdate) AddPossibleUses(i int64) *InviteUpdate {
	iu.mutation.AddPossibleUses(i)
	return iu
}

// AddGuildIDs adds the "guild" edge to the Guild entity by IDs.
func (iu *InviteUpdate) AddGuildIDs(ids ...uint64) *InviteUpdate {
	iu.mutation.AddGuildIDs(ids...)
	return iu
}

// AddGuild adds the "guild" edges to the Guild entity.
func (iu *InviteUpdate) AddGuild(g ...*Guild) *InviteUpdate {
	ids := make([]uint64, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return iu.AddGuildIDs(ids...)
}

// Mutation returns the InviteMutation object of the builder.
func (iu *InviteUpdate) Mutation() *InviteMutation {
	return iu.mutation
}

// ClearGuild clears all "guild" edges to the Guild entity.
func (iu *InviteUpdate) ClearGuild() *InviteUpdate {
	iu.mutation.ClearGuild()
	return iu
}

// RemoveGuildIDs removes the "guild" edge to Guild entities by IDs.
func (iu *InviteUpdate) RemoveGuildIDs(ids ...uint64) *InviteUpdate {
	iu.mutation.RemoveGuildIDs(ids...)
	return iu
}

// RemoveGuild removes "guild" edges to Guild entities.
func (iu *InviteUpdate) RemoveGuild(g ...*Guild) *InviteUpdate {
	ids := make([]uint64, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return iu.RemoveGuildIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (iu *InviteUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(iu.hooks) == 0 {
		affected, err = iu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*InviteMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			iu.mutation = mutation
			affected, err = iu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(iu.hooks) - 1; i >= 0; i-- {
			mut = iu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, iu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (iu *InviteUpdate) SaveX(ctx context.Context) int {
	affected, err := iu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (iu *InviteUpdate) Exec(ctx context.Context) error {
	_, err := iu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (iu *InviteUpdate) ExecX(ctx context.Context) {
	if err := iu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (iu *InviteUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   invite.Table,
			Columns: invite.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: invite.FieldID,
			},
		},
	}
	if ps := iu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := iu.mutation.Code(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: invite.FieldCode,
		})
	}
	if value, ok := iu.mutation.Uses(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: invite.FieldUses,
		})
	}
	if value, ok := iu.mutation.AddedUses(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: invite.FieldUses,
		})
	}
	if value, ok := iu.mutation.PossibleUses(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: invite.FieldPossibleUses,
		})
	}
	if value, ok := iu.mutation.AddedPossibleUses(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: invite.FieldPossibleUses,
		})
	}
	if iu.mutation.GuildCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   invite.GuildTable,
			Columns: []string{invite.GuildColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: guild.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iu.mutation.RemovedGuildIDs(); len(nodes) > 0 && !iu.mutation.GuildCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   invite.GuildTable,
			Columns: []string{invite.GuildColumn},
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iu.mutation.GuildIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   invite.GuildTable,
			Columns: []string{invite.GuildColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, iu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{invite.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// InviteUpdateOne is the builder for updating a single Invite entity.
type InviteUpdateOne struct {
	config
	hooks    []Hook
	mutation *InviteMutation
}

// SetCode sets the "code" field.
func (iuo *InviteUpdateOne) SetCode(s string) *InviteUpdateOne {
	iuo.mutation.SetCode(s)
	return iuo
}

// SetUses sets the "uses" field.
func (iuo *InviteUpdateOne) SetUses(i int64) *InviteUpdateOne {
	iuo.mutation.ResetUses()
	iuo.mutation.SetUses(i)
	return iuo
}

// SetNillableUses sets the "uses" field if the given value is not nil.
func (iuo *InviteUpdateOne) SetNillableUses(i *int64) *InviteUpdateOne {
	if i != nil {
		iuo.SetUses(*i)
	}
	return iuo
}

// AddUses adds i to the "uses" field.
func (iuo *InviteUpdateOne) AddUses(i int64) *InviteUpdateOne {
	iuo.mutation.AddUses(i)
	return iuo
}

// SetPossibleUses sets the "possible_uses" field.
func (iuo *InviteUpdateOne) SetPossibleUses(i int64) *InviteUpdateOne {
	iuo.mutation.ResetPossibleUses()
	iuo.mutation.SetPossibleUses(i)
	return iuo
}

// SetNillablePossibleUses sets the "possible_uses" field if the given value is not nil.
func (iuo *InviteUpdateOne) SetNillablePossibleUses(i *int64) *InviteUpdateOne {
	if i != nil {
		iuo.SetPossibleUses(*i)
	}
	return iuo
}

// AddPossibleUses adds i to the "possible_uses" field.
func (iuo *InviteUpdateOne) AddPossibleUses(i int64) *InviteUpdateOne {
	iuo.mutation.AddPossibleUses(i)
	return iuo
}

// AddGuildIDs adds the "guild" edge to the Guild entity by IDs.
func (iuo *InviteUpdateOne) AddGuildIDs(ids ...uint64) *InviteUpdateOne {
	iuo.mutation.AddGuildIDs(ids...)
	return iuo
}

// AddGuild adds the "guild" edges to the Guild entity.
func (iuo *InviteUpdateOne) AddGuild(g ...*Guild) *InviteUpdateOne {
	ids := make([]uint64, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return iuo.AddGuildIDs(ids...)
}

// Mutation returns the InviteMutation object of the builder.
func (iuo *InviteUpdateOne) Mutation() *InviteMutation {
	return iuo.mutation
}

// ClearGuild clears all "guild" edges to the Guild entity.
func (iuo *InviteUpdateOne) ClearGuild() *InviteUpdateOne {
	iuo.mutation.ClearGuild()
	return iuo
}

// RemoveGuildIDs removes the "guild" edge to Guild entities by IDs.
func (iuo *InviteUpdateOne) RemoveGuildIDs(ids ...uint64) *InviteUpdateOne {
	iuo.mutation.RemoveGuildIDs(ids...)
	return iuo
}

// RemoveGuild removes "guild" edges to Guild entities.
func (iuo *InviteUpdateOne) RemoveGuild(g ...*Guild) *InviteUpdateOne {
	ids := make([]uint64, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return iuo.RemoveGuildIDs(ids...)
}

// Save executes the query and returns the updated Invite entity.
func (iuo *InviteUpdateOne) Save(ctx context.Context) (*Invite, error) {
	var (
		err  error
		node *Invite
	)
	if len(iuo.hooks) == 0 {
		node, err = iuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*InviteMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			iuo.mutation = mutation
			node, err = iuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(iuo.hooks) - 1; i >= 0; i-- {
			mut = iuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, iuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (iuo *InviteUpdateOne) SaveX(ctx context.Context) *Invite {
	node, err := iuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (iuo *InviteUpdateOne) Exec(ctx context.Context) error {
	_, err := iuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (iuo *InviteUpdateOne) ExecX(ctx context.Context) {
	if err := iuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (iuo *InviteUpdateOne) sqlSave(ctx context.Context) (_node *Invite, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   invite.Table,
			Columns: invite.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: invite.FieldID,
			},
		},
	}
	id, ok := iuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Invite.ID for update")}
	}
	_spec.Node.ID.Value = id
	if ps := iuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := iuo.mutation.Code(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: invite.FieldCode,
		})
	}
	if value, ok := iuo.mutation.Uses(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: invite.FieldUses,
		})
	}
	if value, ok := iuo.mutation.AddedUses(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: invite.FieldUses,
		})
	}
	if value, ok := iuo.mutation.PossibleUses(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: invite.FieldPossibleUses,
		})
	}
	if value, ok := iuo.mutation.AddedPossibleUses(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: invite.FieldPossibleUses,
		})
	}
	if iuo.mutation.GuildCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   invite.GuildTable,
			Columns: []string{invite.GuildColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: guild.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iuo.mutation.RemovedGuildIDs(); len(nodes) > 0 && !iuo.mutation.GuildCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   invite.GuildTable,
			Columns: []string{invite.GuildColumn},
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iuo.mutation.GuildIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   invite.GuildTable,
			Columns: []string{invite.GuildColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Invite{config: iuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, iuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{invite.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
