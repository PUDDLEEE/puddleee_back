// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/PUDDLEEE/puddleee_back/ent/predicate"
	"github.com/PUDDLEEE/puddleee_back/ent/verification"
	"github.com/google/uuid"
)

// VerificationUpdate is the builder for updating Verification entities.
type VerificationUpdate struct {
	config
	hooks    []Hook
	mutation *VerificationMutation
}

// Where appends a list predicates to the VerificationUpdate builder.
func (vu *VerificationUpdate) Where(ps ...predicate.Verification) *VerificationUpdate {
	vu.mutation.Where(ps...)
	return vu
}

// SetUpdateTime sets the "update_time" field.
func (vu *VerificationUpdate) SetUpdateTime(t time.Time) *VerificationUpdate {
	vu.mutation.SetUpdateTime(t)
	return vu
}

// SetUUID sets the "uuid" field.
func (vu *VerificationUpdate) SetUUID(u uuid.UUID) *VerificationUpdate {
	vu.mutation.SetUUID(u)
	return vu
}

// SetNillableUUID sets the "uuid" field if the given value is not nil.
func (vu *VerificationUpdate) SetNillableUUID(u *uuid.UUID) *VerificationUpdate {
	if u != nil {
		vu.SetUUID(*u)
	}
	return vu
}

// SetCode sets the "code" field.
func (vu *VerificationUpdate) SetCode(s string) *VerificationUpdate {
	vu.mutation.SetCode(s)
	return vu
}

// SetNillableCode sets the "code" field if the given value is not nil.
func (vu *VerificationUpdate) SetNillableCode(s *string) *VerificationUpdate {
	if s != nil {
		vu.SetCode(*s)
	}
	return vu
}

// SetExpiredAt sets the "expired_at" field.
func (vu *VerificationUpdate) SetExpiredAt(t time.Time) *VerificationUpdate {
	vu.mutation.SetExpiredAt(t)
	return vu
}

// SetNillableExpiredAt sets the "expired_at" field if the given value is not nil.
func (vu *VerificationUpdate) SetNillableExpiredAt(t *time.Time) *VerificationUpdate {
	if t != nil {
		vu.SetExpiredAt(*t)
	}
	return vu
}

// Mutation returns the VerificationMutation object of the builder.
func (vu *VerificationUpdate) Mutation() *VerificationMutation {
	return vu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (vu *VerificationUpdate) Save(ctx context.Context) (int, error) {
	vu.defaults()
	return withHooks(ctx, vu.sqlSave, vu.mutation, vu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (vu *VerificationUpdate) SaveX(ctx context.Context) int {
	affected, err := vu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (vu *VerificationUpdate) Exec(ctx context.Context) error {
	_, err := vu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (vu *VerificationUpdate) ExecX(ctx context.Context) {
	if err := vu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (vu *VerificationUpdate) defaults() {
	if _, ok := vu.mutation.UpdateTime(); !ok {
		v := verification.UpdateDefaultUpdateTime()
		vu.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (vu *VerificationUpdate) check() error {
	if v, ok := vu.mutation.Code(); ok {
		if err := verification.CodeValidator(v); err != nil {
			return &ValidationError{Name: "code", err: fmt.Errorf(`ent: validator failed for field "Verification.code": %w`, err)}
		}
	}
	return nil
}

func (vu *VerificationUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := vu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(verification.Table, verification.Columns, sqlgraph.NewFieldSpec(verification.FieldID, field.TypeInt))
	if ps := vu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := vu.mutation.UpdateTime(); ok {
		_spec.SetField(verification.FieldUpdateTime, field.TypeTime, value)
	}
	if value, ok := vu.mutation.UUID(); ok {
		_spec.SetField(verification.FieldUUID, field.TypeUUID, value)
	}
	if value, ok := vu.mutation.Code(); ok {
		_spec.SetField(verification.FieldCode, field.TypeString, value)
	}
	if value, ok := vu.mutation.ExpiredAt(); ok {
		_spec.SetField(verification.FieldExpiredAt, field.TypeTime, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, vu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{verification.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	vu.mutation.done = true
	return n, nil
}

// VerificationUpdateOne is the builder for updating a single Verification entity.
type VerificationUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *VerificationMutation
}

// SetUpdateTime sets the "update_time" field.
func (vuo *VerificationUpdateOne) SetUpdateTime(t time.Time) *VerificationUpdateOne {
	vuo.mutation.SetUpdateTime(t)
	return vuo
}

// SetUUID sets the "uuid" field.
func (vuo *VerificationUpdateOne) SetUUID(u uuid.UUID) *VerificationUpdateOne {
	vuo.mutation.SetUUID(u)
	return vuo
}

// SetNillableUUID sets the "uuid" field if the given value is not nil.
func (vuo *VerificationUpdateOne) SetNillableUUID(u *uuid.UUID) *VerificationUpdateOne {
	if u != nil {
		vuo.SetUUID(*u)
	}
	return vuo
}

// SetCode sets the "code" field.
func (vuo *VerificationUpdateOne) SetCode(s string) *VerificationUpdateOne {
	vuo.mutation.SetCode(s)
	return vuo
}

// SetNillableCode sets the "code" field if the given value is not nil.
func (vuo *VerificationUpdateOne) SetNillableCode(s *string) *VerificationUpdateOne {
	if s != nil {
		vuo.SetCode(*s)
	}
	return vuo
}

// SetExpiredAt sets the "expired_at" field.
func (vuo *VerificationUpdateOne) SetExpiredAt(t time.Time) *VerificationUpdateOne {
	vuo.mutation.SetExpiredAt(t)
	return vuo
}

// SetNillableExpiredAt sets the "expired_at" field if the given value is not nil.
func (vuo *VerificationUpdateOne) SetNillableExpiredAt(t *time.Time) *VerificationUpdateOne {
	if t != nil {
		vuo.SetExpiredAt(*t)
	}
	return vuo
}

// Mutation returns the VerificationMutation object of the builder.
func (vuo *VerificationUpdateOne) Mutation() *VerificationMutation {
	return vuo.mutation
}

// Where appends a list predicates to the VerificationUpdate builder.
func (vuo *VerificationUpdateOne) Where(ps ...predicate.Verification) *VerificationUpdateOne {
	vuo.mutation.Where(ps...)
	return vuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (vuo *VerificationUpdateOne) Select(field string, fields ...string) *VerificationUpdateOne {
	vuo.fields = append([]string{field}, fields...)
	return vuo
}

// Save executes the query and returns the updated Verification entity.
func (vuo *VerificationUpdateOne) Save(ctx context.Context) (*Verification, error) {
	vuo.defaults()
	return withHooks(ctx, vuo.sqlSave, vuo.mutation, vuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (vuo *VerificationUpdateOne) SaveX(ctx context.Context) *Verification {
	node, err := vuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (vuo *VerificationUpdateOne) Exec(ctx context.Context) error {
	_, err := vuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (vuo *VerificationUpdateOne) ExecX(ctx context.Context) {
	if err := vuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (vuo *VerificationUpdateOne) defaults() {
	if _, ok := vuo.mutation.UpdateTime(); !ok {
		v := verification.UpdateDefaultUpdateTime()
		vuo.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (vuo *VerificationUpdateOne) check() error {
	if v, ok := vuo.mutation.Code(); ok {
		if err := verification.CodeValidator(v); err != nil {
			return &ValidationError{Name: "code", err: fmt.Errorf(`ent: validator failed for field "Verification.code": %w`, err)}
		}
	}
	return nil
}

func (vuo *VerificationUpdateOne) sqlSave(ctx context.Context) (_node *Verification, err error) {
	if err := vuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(verification.Table, verification.Columns, sqlgraph.NewFieldSpec(verification.FieldID, field.TypeInt))
	id, ok := vuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Verification.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := vuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, verification.FieldID)
		for _, f := range fields {
			if !verification.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != verification.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := vuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := vuo.mutation.UpdateTime(); ok {
		_spec.SetField(verification.FieldUpdateTime, field.TypeTime, value)
	}
	if value, ok := vuo.mutation.UUID(); ok {
		_spec.SetField(verification.FieldUUID, field.TypeUUID, value)
	}
	if value, ok := vuo.mutation.Code(); ok {
		_spec.SetField(verification.FieldCode, field.TypeString, value)
	}
	if value, ok := vuo.mutation.ExpiredAt(); ok {
		_spec.SetField(verification.FieldExpiredAt, field.TypeTime, value)
	}
	_node = &Verification{config: vuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, vuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{verification.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	vuo.mutation.done = true
	return _node, nil
}