// Code generated by entc, DO NOT EDIT.

package entgen

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/harmony-development/legato/server/db/ent/entgen/file"
	"github.com/harmony-development/legato/server/db/ent/entgen/filemessage"
)

// FileMessageCreate is the builder for creating a FileMessage entity.
type FileMessageCreate struct {
	config
	mutation *FileMessageMutation
	hooks    []Hook
}

// AddFileIDs adds the "file" edge to the File entity by IDs.
func (fmc *FileMessageCreate) AddFileIDs(ids ...string) *FileMessageCreate {
	fmc.mutation.AddFileIDs(ids...)
	return fmc
}

// AddFile adds the "file" edges to the File entity.
func (fmc *FileMessageCreate) AddFile(f ...*File) *FileMessageCreate {
	ids := make([]string, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return fmc.AddFileIDs(ids...)
}

// Mutation returns the FileMessageMutation object of the builder.
func (fmc *FileMessageCreate) Mutation() *FileMessageMutation {
	return fmc.mutation
}

// Save creates the FileMessage in the database.
func (fmc *FileMessageCreate) Save(ctx context.Context) (*FileMessage, error) {
	var (
		err  error
		node *FileMessage
	)
	if len(fmc.hooks) == 0 {
		if err = fmc.check(); err != nil {
			return nil, err
		}
		node, err = fmc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*FileMessageMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = fmc.check(); err != nil {
				return nil, err
			}
			fmc.mutation = mutation
			node, err = fmc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(fmc.hooks) - 1; i >= 0; i-- {
			mut = fmc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, fmc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (fmc *FileMessageCreate) SaveX(ctx context.Context) *FileMessage {
	v, err := fmc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// check runs all checks and user-defined validators on the builder.
func (fmc *FileMessageCreate) check() error {
	return nil
}

func (fmc *FileMessageCreate) sqlSave(ctx context.Context) (*FileMessage, error) {
	_node, _spec := fmc.createSpec()
	if err := sqlgraph.CreateNode(ctx, fmc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (fmc *FileMessageCreate) createSpec() (*FileMessage, *sqlgraph.CreateSpec) {
	var (
		_node = &FileMessage{config: fmc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: filemessage.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: filemessage.FieldID,
			},
		}
	)
	if nodes := fmc.mutation.FileIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   filemessage.FileTable,
			Columns: []string{filemessage.FileColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: file.FieldID,
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

// FileMessageCreateBulk is the builder for creating many FileMessage entities in bulk.
type FileMessageCreateBulk struct {
	config
	builders []*FileMessageCreate
}

// Save creates the FileMessage entities in the database.
func (fmcb *FileMessageCreateBulk) Save(ctx context.Context) ([]*FileMessage, error) {
	specs := make([]*sqlgraph.CreateSpec, len(fmcb.builders))
	nodes := make([]*FileMessage, len(fmcb.builders))
	mutators := make([]Mutator, len(fmcb.builders))
	for i := range fmcb.builders {
		func(i int, root context.Context) {
			builder := fmcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*FileMessageMutation)
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
					_, err = mutators[i+1].Mutate(root, fmcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, fmcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
						if cerr, ok := isSQLConstraintError(err); ok {
							err = cerr
						}
					}
				}
				mutation.done = true
				if err != nil {
					return nil, err
				}
				id := specs[i].ID.Value.(int64)
				nodes[i].ID = int(id)
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, fmcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (fmcb *FileMessageCreateBulk) SaveX(ctx context.Context) []*FileMessage {
	v, err := fmcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
