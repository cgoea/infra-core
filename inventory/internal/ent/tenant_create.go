// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/open-edge-platform/infra-core/inventory/v2/internal/ent/tenant"
)

// TenantCreate is the builder for creating a Tenant entity.
type TenantCreate struct {
	config
	mutation *TenantMutation
	hooks    []Hook
}

// SetResourceID sets the "resource_id" field.
func (tc *TenantCreate) SetResourceID(s string) *TenantCreate {
	tc.mutation.SetResourceID(s)
	return tc
}

// SetCurrentState sets the "current_state" field.
func (tc *TenantCreate) SetCurrentState(ts tenant.CurrentState) *TenantCreate {
	tc.mutation.SetCurrentState(ts)
	return tc
}

// SetNillableCurrentState sets the "current_state" field if the given value is not nil.
func (tc *TenantCreate) SetNillableCurrentState(ts *tenant.CurrentState) *TenantCreate {
	if ts != nil {
		tc.SetCurrentState(*ts)
	}
	return tc
}

// SetDesiredState sets the "desired_state" field.
func (tc *TenantCreate) SetDesiredState(ts tenant.DesiredState) *TenantCreate {
	tc.mutation.SetDesiredState(ts)
	return tc
}

// SetWatcherOsmanager sets the "watcher_osmanager" field.
func (tc *TenantCreate) SetWatcherOsmanager(b bool) *TenantCreate {
	tc.mutation.SetWatcherOsmanager(b)
	return tc
}

// SetNillableWatcherOsmanager sets the "watcher_osmanager" field if the given value is not nil.
func (tc *TenantCreate) SetNillableWatcherOsmanager(b *bool) *TenantCreate {
	if b != nil {
		tc.SetWatcherOsmanager(*b)
	}
	return tc
}

// SetTenantID sets the "tenant_id" field.
func (tc *TenantCreate) SetTenantID(s string) *TenantCreate {
	tc.mutation.SetTenantID(s)
	return tc
}

// SetCreatedAt sets the "created_at" field.
func (tc *TenantCreate) SetCreatedAt(s string) *TenantCreate {
	tc.mutation.SetCreatedAt(s)
	return tc
}

// SetUpdatedAt sets the "updated_at" field.
func (tc *TenantCreate) SetUpdatedAt(s string) *TenantCreate {
	tc.mutation.SetUpdatedAt(s)
	return tc
}

// Mutation returns the TenantMutation object of the builder.
func (tc *TenantCreate) Mutation() *TenantMutation {
	return tc.mutation
}

// Save creates the Tenant in the database.
func (tc *TenantCreate) Save(ctx context.Context) (*Tenant, error) {
	return withHooks(ctx, tc.sqlSave, tc.mutation, tc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (tc *TenantCreate) SaveX(ctx context.Context) *Tenant {
	v, err := tc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tc *TenantCreate) Exec(ctx context.Context) error {
	_, err := tc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tc *TenantCreate) ExecX(ctx context.Context) {
	if err := tc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tc *TenantCreate) check() error {
	if _, ok := tc.mutation.ResourceID(); !ok {
		return &ValidationError{Name: "resource_id", err: errors.New(`ent: missing required field "Tenant.resource_id"`)}
	}
	if v, ok := tc.mutation.CurrentState(); ok {
		if err := tenant.CurrentStateValidator(v); err != nil {
			return &ValidationError{Name: "current_state", err: fmt.Errorf(`ent: validator failed for field "Tenant.current_state": %w`, err)}
		}
	}
	if _, ok := tc.mutation.DesiredState(); !ok {
		return &ValidationError{Name: "desired_state", err: errors.New(`ent: missing required field "Tenant.desired_state"`)}
	}
	if v, ok := tc.mutation.DesiredState(); ok {
		if err := tenant.DesiredStateValidator(v); err != nil {
			return &ValidationError{Name: "desired_state", err: fmt.Errorf(`ent: validator failed for field "Tenant.desired_state": %w`, err)}
		}
	}
	if _, ok := tc.mutation.TenantID(); !ok {
		return &ValidationError{Name: "tenant_id", err: errors.New(`ent: missing required field "Tenant.tenant_id"`)}
	}
	if _, ok := tc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Tenant.created_at"`)}
	}
	if _, ok := tc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Tenant.updated_at"`)}
	}
	return nil
}

func (tc *TenantCreate) sqlSave(ctx context.Context) (*Tenant, error) {
	if err := tc.check(); err != nil {
		return nil, err
	}
	_node, _spec := tc.createSpec()
	if err := sqlgraph.CreateNode(ctx, tc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	tc.mutation.id = &_node.ID
	tc.mutation.done = true
	return _node, nil
}

func (tc *TenantCreate) createSpec() (*Tenant, *sqlgraph.CreateSpec) {
	var (
		_node = &Tenant{config: tc.config}
		_spec = sqlgraph.NewCreateSpec(tenant.Table, sqlgraph.NewFieldSpec(tenant.FieldID, field.TypeInt))
	)
	if value, ok := tc.mutation.ResourceID(); ok {
		_spec.SetField(tenant.FieldResourceID, field.TypeString, value)
		_node.ResourceID = value
	}
	if value, ok := tc.mutation.CurrentState(); ok {
		_spec.SetField(tenant.FieldCurrentState, field.TypeEnum, value)
		_node.CurrentState = value
	}
	if value, ok := tc.mutation.DesiredState(); ok {
		_spec.SetField(tenant.FieldDesiredState, field.TypeEnum, value)
		_node.DesiredState = value
	}
	if value, ok := tc.mutation.WatcherOsmanager(); ok {
		_spec.SetField(tenant.FieldWatcherOsmanager, field.TypeBool, value)
		_node.WatcherOsmanager = value
	}
	if value, ok := tc.mutation.TenantID(); ok {
		_spec.SetField(tenant.FieldTenantID, field.TypeString, value)
		_node.TenantID = value
	}
	if value, ok := tc.mutation.CreatedAt(); ok {
		_spec.SetField(tenant.FieldCreatedAt, field.TypeString, value)
		_node.CreatedAt = value
	}
	if value, ok := tc.mutation.UpdatedAt(); ok {
		_spec.SetField(tenant.FieldUpdatedAt, field.TypeString, value)
		_node.UpdatedAt = value
	}
	return _node, _spec
}

// TenantCreateBulk is the builder for creating many Tenant entities in bulk.
type TenantCreateBulk struct {
	config
	err      error
	builders []*TenantCreate
}

// Save creates the Tenant entities in the database.
func (tcb *TenantCreateBulk) Save(ctx context.Context) ([]*Tenant, error) {
	if tcb.err != nil {
		return nil, tcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(tcb.builders))
	nodes := make([]*Tenant, len(tcb.builders))
	mutators := make([]Mutator, len(tcb.builders))
	for i := range tcb.builders {
		func(i int, root context.Context) {
			builder := tcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TenantMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, tcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, tcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, tcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (tcb *TenantCreateBulk) SaveX(ctx context.Context) []*Tenant {
	v, err := tcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tcb *TenantCreateBulk) Exec(ctx context.Context) error {
	_, err := tcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tcb *TenantCreateBulk) ExecX(ctx context.Context) {
	if err := tcb.Exec(ctx); err != nil {
		panic(err)
	}
}
