// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"gonebook/internal/ent/contact"
	"gonebook/internal/ent/predicate"

	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
)

// ContactDelete is the builder for deleting a Contact entity.
type ContactDelete struct {
	config
	hooks      []Hook
	mutation   *ContactMutation
	predicates []predicate.Contact
}

// Where adds a new predicate to the delete builder.
func (cd *ContactDelete) Where(ps ...predicate.Contact) *ContactDelete {
	cd.predicates = append(cd.predicates, ps...)
	return cd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (cd *ContactDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(cd.hooks) == 0 {
		affected, err = cd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ContactMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			cd.mutation = mutation
			affected, err = cd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(cd.hooks) - 1; i >= 0; i-- {
			mut = cd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (cd *ContactDelete) ExecX(ctx context.Context) int {
	n, err := cd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (cd *ContactDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: contact.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: contact.FieldID,
			},
		},
	}
	if ps := cd.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, cd.driver, _spec)
}

// ContactDeleteOne is the builder for deleting a single Contact entity.
type ContactDeleteOne struct {
	cd *ContactDelete
}

// Exec executes the deletion query.
func (cdo *ContactDeleteOne) Exec(ctx context.Context) error {
	n, err := cdo.cd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{contact.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (cdo *ContactDeleteOne) ExecX(ctx context.Context) {
	cdo.cd.ExecX(ctx)
}