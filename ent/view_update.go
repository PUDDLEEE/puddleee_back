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
	"github.com/PUDDLEEE/puddleee_back/ent/room"
	"github.com/PUDDLEEE/puddleee_back/ent/view"
)

// ViewUpdate is the builder for updating View entities.
type ViewUpdate struct {
	config
	hooks    []Hook
	mutation *ViewMutation
}

// Where appends a list predicates to the ViewUpdate builder.
func (vu *ViewUpdate) Where(ps ...predicate.View) *ViewUpdate {
	vu.mutation.Where(ps...)
	return vu
}

// SetUpdateTime sets the "update_time" field.
func (vu *ViewUpdate) SetUpdateTime(t time.Time) *ViewUpdate {
	vu.mutation.SetUpdateTime(t)
	return vu
}

// SetFilepath sets the "filepath" field.
func (vu *ViewUpdate) SetFilepath(s string) *ViewUpdate {
	vu.mutation.SetFilepath(s)
	return vu
}

// SetNillableFilepath sets the "filepath" field if the given value is not nil.
func (vu *ViewUpdate) SetNillableFilepath(s *string) *ViewUpdate {
	if s != nil {
		vu.SetFilepath(*s)
	}
	return vu
}

// SetRoomID sets the "room" edge to the Room entity by ID.
func (vu *ViewUpdate) SetRoomID(id int) *ViewUpdate {
	vu.mutation.SetRoomID(id)
	return vu
}

// SetRoom sets the "room" edge to the Room entity.
func (vu *ViewUpdate) SetRoom(r *Room) *ViewUpdate {
	return vu.SetRoomID(r.ID)
}

// Mutation returns the ViewMutation object of the builder.
func (vu *ViewUpdate) Mutation() *ViewMutation {
	return vu.mutation
}

// ClearRoom clears the "room" edge to the Room entity.
func (vu *ViewUpdate) ClearRoom() *ViewUpdate {
	vu.mutation.ClearRoom()
	return vu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (vu *ViewUpdate) Save(ctx context.Context) (int, error) {
	vu.defaults()
	return withHooks(ctx, vu.sqlSave, vu.mutation, vu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (vu *ViewUpdate) SaveX(ctx context.Context) int {
	affected, err := vu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (vu *ViewUpdate) Exec(ctx context.Context) error {
	_, err := vu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (vu *ViewUpdate) ExecX(ctx context.Context) {
	if err := vu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (vu *ViewUpdate) defaults() {
	if _, ok := vu.mutation.UpdateTime(); !ok {
		v := view.UpdateDefaultUpdateTime()
		vu.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (vu *ViewUpdate) check() error {
	if v, ok := vu.mutation.Filepath(); ok {
		if err := view.FilepathValidator(v); err != nil {
			return &ValidationError{Name: "filepath", err: fmt.Errorf(`ent: validator failed for field "View.filepath": %w`, err)}
		}
	}
	if _, ok := vu.mutation.RoomID(); vu.mutation.RoomCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "View.room"`)
	}
	return nil
}

func (vu *ViewUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := vu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(view.Table, view.Columns, sqlgraph.NewFieldSpec(view.FieldID, field.TypeInt))
	if ps := vu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := vu.mutation.UpdateTime(); ok {
		_spec.SetField(view.FieldUpdateTime, field.TypeTime, value)
	}
	if value, ok := vu.mutation.Filepath(); ok {
		_spec.SetField(view.FieldFilepath, field.TypeString, value)
	}
	if vu.mutation.RoomCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   view.RoomTable,
			Columns: []string{view.RoomColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(room.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := vu.mutation.RoomIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   view.RoomTable,
			Columns: []string{view.RoomColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(room.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, vu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{view.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	vu.mutation.done = true
	return n, nil
}

// ViewUpdateOne is the builder for updating a single View entity.
type ViewUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ViewMutation
}

// SetUpdateTime sets the "update_time" field.
func (vuo *ViewUpdateOne) SetUpdateTime(t time.Time) *ViewUpdateOne {
	vuo.mutation.SetUpdateTime(t)
	return vuo
}

// SetFilepath sets the "filepath" field.
func (vuo *ViewUpdateOne) SetFilepath(s string) *ViewUpdateOne {
	vuo.mutation.SetFilepath(s)
	return vuo
}

// SetNillableFilepath sets the "filepath" field if the given value is not nil.
func (vuo *ViewUpdateOne) SetNillableFilepath(s *string) *ViewUpdateOne {
	if s != nil {
		vuo.SetFilepath(*s)
	}
	return vuo
}

// SetRoomID sets the "room" edge to the Room entity by ID.
func (vuo *ViewUpdateOne) SetRoomID(id int) *ViewUpdateOne {
	vuo.mutation.SetRoomID(id)
	return vuo
}

// SetRoom sets the "room" edge to the Room entity.
func (vuo *ViewUpdateOne) SetRoom(r *Room) *ViewUpdateOne {
	return vuo.SetRoomID(r.ID)
}

// Mutation returns the ViewMutation object of the builder.
func (vuo *ViewUpdateOne) Mutation() *ViewMutation {
	return vuo.mutation
}

// ClearRoom clears the "room" edge to the Room entity.
func (vuo *ViewUpdateOne) ClearRoom() *ViewUpdateOne {
	vuo.mutation.ClearRoom()
	return vuo
}

// Where appends a list predicates to the ViewUpdate builder.
func (vuo *ViewUpdateOne) Where(ps ...predicate.View) *ViewUpdateOne {
	vuo.mutation.Where(ps...)
	return vuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (vuo *ViewUpdateOne) Select(field string, fields ...string) *ViewUpdateOne {
	vuo.fields = append([]string{field}, fields...)
	return vuo
}

// Save executes the query and returns the updated View entity.
func (vuo *ViewUpdateOne) Save(ctx context.Context) (*View, error) {
	vuo.defaults()
	return withHooks(ctx, vuo.sqlSave, vuo.mutation, vuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (vuo *ViewUpdateOne) SaveX(ctx context.Context) *View {
	node, err := vuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (vuo *ViewUpdateOne) Exec(ctx context.Context) error {
	_, err := vuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (vuo *ViewUpdateOne) ExecX(ctx context.Context) {
	if err := vuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (vuo *ViewUpdateOne) defaults() {
	if _, ok := vuo.mutation.UpdateTime(); !ok {
		v := view.UpdateDefaultUpdateTime()
		vuo.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (vuo *ViewUpdateOne) check() error {
	if v, ok := vuo.mutation.Filepath(); ok {
		if err := view.FilepathValidator(v); err != nil {
			return &ValidationError{Name: "filepath", err: fmt.Errorf(`ent: validator failed for field "View.filepath": %w`, err)}
		}
	}
	if _, ok := vuo.mutation.RoomID(); vuo.mutation.RoomCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "View.room"`)
	}
	return nil
}

func (vuo *ViewUpdateOne) sqlSave(ctx context.Context) (_node *View, err error) {
	if err := vuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(view.Table, view.Columns, sqlgraph.NewFieldSpec(view.FieldID, field.TypeInt))
	id, ok := vuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "View.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := vuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, view.FieldID)
		for _, f := range fields {
			if !view.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != view.FieldID {
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
		_spec.SetField(view.FieldUpdateTime, field.TypeTime, value)
	}
	if value, ok := vuo.mutation.Filepath(); ok {
		_spec.SetField(view.FieldFilepath, field.TypeString, value)
	}
	if vuo.mutation.RoomCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   view.RoomTable,
			Columns: []string{view.RoomColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(room.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := vuo.mutation.RoomIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   view.RoomTable,
			Columns: []string{view.RoomColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(room.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &View{config: vuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, vuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{view.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	vuo.mutation.done = true
	return _node, nil
}
