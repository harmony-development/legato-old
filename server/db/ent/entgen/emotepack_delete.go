// Code generated by entc, DO NOT EDIT.

package entgen

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/harmony-development/legato/server/db/ent/entgen/emotepack"
	"github.com/harmony-development/legato/server/db/ent/entgen/predicate"
)

// EmotePackDelete is the builder for deleting a EmotePack entity.
type EmotePackDelete struct {
	config
	hooks    []Hook
	mutation *EmotePackMutation
}

// Where adds a new predicate to the EmotePackDelete builder.
func (epd *EmotePackDelete) Where(ps ...predicate.EmotePack) *EmotePackDelete {
	epd.mutation.predicates = append(epd.mutation.predicates, ps...)
	return epd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (epd *EmotePackDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(epd.hooks) == 0 {
		affected, err = epd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*EmotePackMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			epd.mutation = mutation
			affected, err = epd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(epd.hooks) - 1; i >= 0; i-- {
			mut = epd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, epd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (epd *EmotePackDelete) ExecX(ctx context.Context) int {
	n, err := epd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (epd *EmotePackDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: emotepack.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: emotepack.FieldID,
			},
		},
	}
	if ps := epd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, epd.driver, _spec)
}

// EmotePackDeleteOne is the builder for deleting a single EmotePack entity.
type EmotePackDeleteOne struct {
	epd *EmotePackDelete
}

// Exec executes the deletion query.
func (epdo *EmotePackDeleteOne) Exec(ctx context.Context) error {
	n, err := epdo.epd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{emotepack.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (epdo *EmotePackDeleteOne) ExecX(ctx context.Context) {
	epdo.epd.ExecX(ctx)
}