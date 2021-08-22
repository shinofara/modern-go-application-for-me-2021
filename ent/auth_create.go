// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"mygo/ent/auth"
	"mygo/ent/user"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// AuthCreate is the builder for creating a Auth entity.
type AuthCreate struct {
	config
	mutation *AuthMutation
	hooks    []Hook
}

// SetEmail sets the "email" field.
func (ac *AuthCreate) SetEmail(s string) *AuthCreate {
	ac.mutation.SetEmail(s)
	return ac
}

// SetPassword sets the "password" field.
func (ac *AuthCreate) SetPassword(s string) *AuthCreate {
	ac.mutation.SetPassword(s)
	return ac
}

// SetUserID sets the "user" edge to the User entity by ID.
func (ac *AuthCreate) SetUserID(id int) *AuthCreate {
	ac.mutation.SetUserID(id)
	return ac
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (ac *AuthCreate) SetNillableUserID(id *int) *AuthCreate {
	if id != nil {
		ac = ac.SetUserID(*id)
	}
	return ac
}

// SetUser sets the "user" edge to the User entity.
func (ac *AuthCreate) SetUser(u *User) *AuthCreate {
	return ac.SetUserID(u.ID)
}

// Mutation returns the AuthMutation object of the builder.
func (ac *AuthCreate) Mutation() *AuthMutation {
	return ac.mutation
}

// Save creates the Auth in the database.
func (ac *AuthCreate) Save(ctx context.Context) (*Auth, error) {
	var (
		err  error
		node *Auth
	)
	if len(ac.hooks) == 0 {
		if err = ac.check(); err != nil {
			return nil, err
		}
		node, err = ac.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AuthMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ac.check(); err != nil {
				return nil, err
			}
			ac.mutation = mutation
			if node, err = ac.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(ac.hooks) - 1; i >= 0; i-- {
			if ac.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ac.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ac.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (ac *AuthCreate) SaveX(ctx context.Context) *Auth {
	v, err := ac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ac *AuthCreate) Exec(ctx context.Context) error {
	_, err := ac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ac *AuthCreate) ExecX(ctx context.Context) {
	if err := ac.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ac *AuthCreate) check() error {
	if _, ok := ac.mutation.Email(); !ok {
		return &ValidationError{Name: "email", err: errors.New(`ent: missing required field "email"`)}
	}
	if v, ok := ac.mutation.Email(); ok {
		if err := auth.EmailValidator(v); err != nil {
			return &ValidationError{Name: "email", err: fmt.Errorf(`ent: validator failed for field "email": %w`, err)}
		}
	}
	if _, ok := ac.mutation.Password(); !ok {
		return &ValidationError{Name: "password", err: errors.New(`ent: missing required field "password"`)}
	}
	if v, ok := ac.mutation.Password(); ok {
		if err := auth.PasswordValidator(v); err != nil {
			return &ValidationError{Name: "password", err: fmt.Errorf(`ent: validator failed for field "password": %w`, err)}
		}
	}
	return nil
}

func (ac *AuthCreate) sqlSave(ctx context.Context) (*Auth, error) {
	_node, _spec := ac.createSpec()
	if err := sqlgraph.CreateNode(ctx, ac.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (ac *AuthCreate) createSpec() (*Auth, *sqlgraph.CreateSpec) {
	var (
		_node = &Auth{config: ac.config}
		_spec = &sqlgraph.CreateSpec{
			Table: auth.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: auth.FieldID,
			},
		}
	)
	if value, ok := ac.mutation.Email(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: auth.FieldEmail,
		})
		_node.Email = value
	}
	if value, ok := ac.mutation.Password(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: auth.FieldPassword,
		})
		_node.Password = value
	}
	if nodes := ac.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   auth.UserTable,
			Columns: []string{auth.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
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

// AuthCreateBulk is the builder for creating many Auth entities in bulk.
type AuthCreateBulk struct {
	config
	builders []*AuthCreate
}

// Save creates the Auth entities in the database.
func (acb *AuthCreateBulk) Save(ctx context.Context) ([]*Auth, error) {
	specs := make([]*sqlgraph.CreateSpec, len(acb.builders))
	nodes := make([]*Auth, len(acb.builders))
	mutators := make([]Mutator, len(acb.builders))
	for i := range acb.builders {
		func(i int, root context.Context) {
			builder := acb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AuthMutation)
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
					_, err = mutators[i+1].Mutate(root, acb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, acb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
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
		if _, err := mutators[0].Mutate(ctx, acb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (acb *AuthCreateBulk) SaveX(ctx context.Context) []*Auth {
	v, err := acb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (acb *AuthCreateBulk) Exec(ctx context.Context) error {
	_, err := acb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (acb *AuthCreateBulk) ExecX(ctx context.Context) {
	if err := acb.Exec(ctx); err != nil {
		panic(err)
	}
}
