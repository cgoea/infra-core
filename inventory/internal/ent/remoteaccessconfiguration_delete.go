// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/open-edge-platform/infra-core/inventory/v2/internal/ent/predicate"
	"github.com/open-edge-platform/infra-core/inventory/v2/internal/ent/remoteaccessconfiguration"
)

// RemoteAccessConfigurationDelete is the builder for deleting a RemoteAccessConfiguration entity.
type RemoteAccessConfigurationDelete struct {
	config
	hooks    []Hook
	mutation *RemoteAccessConfigurationMutation
}

// Where appends a list predicates to the RemoteAccessConfigurationDelete builder.
func (racd *RemoteAccessConfigurationDelete) Where(ps ...predicate.RemoteAccessConfiguration) *RemoteAccessConfigurationDelete {
	racd.mutation.Where(ps...)
	return racd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (racd *RemoteAccessConfigurationDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, racd.sqlExec, racd.mutation, racd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (racd *RemoteAccessConfigurationDelete) ExecX(ctx context.Context) int {
	n, err := racd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (racd *RemoteAccessConfigurationDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(remoteaccessconfiguration.Table, sqlgraph.NewFieldSpec(remoteaccessconfiguration.FieldID, field.TypeInt))
	if ps := racd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, racd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	racd.mutation.done = true
	return affected, err
}

// RemoteAccessConfigurationDeleteOne is the builder for deleting a single RemoteAccessConfiguration entity.
type RemoteAccessConfigurationDeleteOne struct {
	racd *RemoteAccessConfigurationDelete
}

// Where appends a list predicates to the RemoteAccessConfigurationDelete builder.
func (racdo *RemoteAccessConfigurationDeleteOne) Where(ps ...predicate.RemoteAccessConfiguration) *RemoteAccessConfigurationDeleteOne {
	racdo.racd.mutation.Where(ps...)
	return racdo
}

// Exec executes the deletion query.
func (racdo *RemoteAccessConfigurationDeleteOne) Exec(ctx context.Context) error {
	n, err := racdo.racd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{remoteaccessconfiguration.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (racdo *RemoteAccessConfigurationDeleteOne) ExecX(ctx context.Context) {
	if err := racdo.Exec(ctx); err != nil {
		panic(err)
	}
}
