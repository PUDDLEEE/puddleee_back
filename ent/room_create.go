// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/PUDDLEEE/puddleee_back/ent/category"
	"github.com/PUDDLEEE/puddleee_back/ent/message"
	"github.com/PUDDLEEE/puddleee_back/ent/room"
	"github.com/PUDDLEEE/puddleee_back/ent/user"
	"github.com/PUDDLEEE/puddleee_back/ent/view"
)

// RoomCreate is the builder for creating a Room entity.
type RoomCreate struct {
	config
	mutation *RoomMutation
	hooks    []Hook
}

// SetCreateTime sets the "create_time" field.
func (rc *RoomCreate) SetCreateTime(t time.Time) *RoomCreate {
	rc.mutation.SetCreateTime(t)
	return rc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (rc *RoomCreate) SetNillableCreateTime(t *time.Time) *RoomCreate {
	if t != nil {
		rc.SetCreateTime(*t)
	}
	return rc
}

// SetUpdateTime sets the "update_time" field.
func (rc *RoomCreate) SetUpdateTime(t time.Time) *RoomCreate {
	rc.mutation.SetUpdateTime(t)
	return rc
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (rc *RoomCreate) SetNillableUpdateTime(t *time.Time) *RoomCreate {
	if t != nil {
		rc.SetUpdateTime(*t)
	}
	return rc
}

// SetTitle sets the "title" field.
func (rc *RoomCreate) SetTitle(s string) *RoomCreate {
	rc.mutation.SetTitle(s)
	return rc
}

// SetIsCompleted sets the "is_completed" field.
func (rc *RoomCreate) SetIsCompleted(b bool) *RoomCreate {
	rc.mutation.SetIsCompleted(b)
	return rc
}

// SetNillableIsCompleted sets the "is_completed" field if the given value is not nil.
func (rc *RoomCreate) SetNillableIsCompleted(b *bool) *RoomCreate {
	if b != nil {
		rc.SetIsCompleted(*b)
	}
	return rc
}

// SetQuestionerID sets the "questioner" edge to the User entity by ID.
func (rc *RoomCreate) SetQuestionerID(id int) *RoomCreate {
	rc.mutation.SetQuestionerID(id)
	return rc
}

// SetNillableQuestionerID sets the "questioner" edge to the User entity by ID if the given value is not nil.
func (rc *RoomCreate) SetNillableQuestionerID(id *int) *RoomCreate {
	if id != nil {
		rc = rc.SetQuestionerID(*id)
	}
	return rc
}

// SetQuestioner sets the "questioner" edge to the User entity.
func (rc *RoomCreate) SetQuestioner(u *User) *RoomCreate {
	return rc.SetQuestionerID(u.ID)
}

// AddRespondentIDs adds the "respondent" edge to the User entity by IDs.
func (rc *RoomCreate) AddRespondentIDs(ids ...int) *RoomCreate {
	rc.mutation.AddRespondentIDs(ids...)
	return rc
}

// AddRespondent adds the "respondent" edges to the User entity.
func (rc *RoomCreate) AddRespondent(u ...*User) *RoomCreate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return rc.AddRespondentIDs(ids...)
}

// SetCategoryID sets the "category" edge to the Category entity by ID.
func (rc *RoomCreate) SetCategoryID(id int) *RoomCreate {
	rc.mutation.SetCategoryID(id)
	return rc
}

// SetNillableCategoryID sets the "category" edge to the Category entity by ID if the given value is not nil.
func (rc *RoomCreate) SetNillableCategoryID(id *int) *RoomCreate {
	if id != nil {
		rc = rc.SetCategoryID(*id)
	}
	return rc
}

// SetCategory sets the "category" edge to the Category entity.
func (rc *RoomCreate) SetCategory(c *Category) *RoomCreate {
	return rc.SetCategoryID(c.ID)
}

// AddMessageIDs adds the "messages" edge to the Message entity by IDs.
func (rc *RoomCreate) AddMessageIDs(ids ...int) *RoomCreate {
	rc.mutation.AddMessageIDs(ids...)
	return rc
}

// AddMessages adds the "messages" edges to the Message entity.
func (rc *RoomCreate) AddMessages(m ...*Message) *RoomCreate {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return rc.AddMessageIDs(ids...)
}

// AddViewIDs adds the "views" edge to the View entity by IDs.
func (rc *RoomCreate) AddViewIDs(ids ...int) *RoomCreate {
	rc.mutation.AddViewIDs(ids...)
	return rc
}

// AddViews adds the "views" edges to the View entity.
func (rc *RoomCreate) AddViews(v ...*View) *RoomCreate {
	ids := make([]int, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return rc.AddViewIDs(ids...)
}

// Mutation returns the RoomMutation object of the builder.
func (rc *RoomCreate) Mutation() *RoomMutation {
	return rc.mutation
}

// Save creates the Room in the database.
func (rc *RoomCreate) Save(ctx context.Context) (*Room, error) {
	rc.defaults()
	return withHooks(ctx, rc.sqlSave, rc.mutation, rc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (rc *RoomCreate) SaveX(ctx context.Context) *Room {
	v, err := rc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rc *RoomCreate) Exec(ctx context.Context) error {
	_, err := rc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rc *RoomCreate) ExecX(ctx context.Context) {
	if err := rc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (rc *RoomCreate) defaults() {
	if _, ok := rc.mutation.CreateTime(); !ok {
		v := room.DefaultCreateTime()
		rc.mutation.SetCreateTime(v)
	}
	if _, ok := rc.mutation.UpdateTime(); !ok {
		v := room.DefaultUpdateTime()
		rc.mutation.SetUpdateTime(v)
	}
	if _, ok := rc.mutation.IsCompleted(); !ok {
		v := room.DefaultIsCompleted
		rc.mutation.SetIsCompleted(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (rc *RoomCreate) check() error {
	if _, ok := rc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New(`ent: missing required field "Room.create_time"`)}
	}
	if _, ok := rc.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New(`ent: missing required field "Room.update_time"`)}
	}
	if _, ok := rc.mutation.Title(); !ok {
		return &ValidationError{Name: "title", err: errors.New(`ent: missing required field "Room.title"`)}
	}
	if v, ok := rc.mutation.Title(); ok {
		if err := room.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "Room.title": %w`, err)}
		}
	}
	if _, ok := rc.mutation.IsCompleted(); !ok {
		return &ValidationError{Name: "is_completed", err: errors.New(`ent: missing required field "Room.is_completed"`)}
	}
	return nil
}

func (rc *RoomCreate) sqlSave(ctx context.Context) (*Room, error) {
	if err := rc.check(); err != nil {
		return nil, err
	}
	_node, _spec := rc.createSpec()
	if err := sqlgraph.CreateNode(ctx, rc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	rc.mutation.id = &_node.ID
	rc.mutation.done = true
	return _node, nil
}

func (rc *RoomCreate) createSpec() (*Room, *sqlgraph.CreateSpec) {
	var (
		_node = &Room{config: rc.config}
		_spec = sqlgraph.NewCreateSpec(room.Table, sqlgraph.NewFieldSpec(room.FieldID, field.TypeInt))
	)
	if value, ok := rc.mutation.CreateTime(); ok {
		_spec.SetField(room.FieldCreateTime, field.TypeTime, value)
		_node.CreateTime = value
	}
	if value, ok := rc.mutation.UpdateTime(); ok {
		_spec.SetField(room.FieldUpdateTime, field.TypeTime, value)
		_node.UpdateTime = value
	}
	if value, ok := rc.mutation.Title(); ok {
		_spec.SetField(room.FieldTitle, field.TypeString, value)
		_node.Title = value
	}
	if value, ok := rc.mutation.IsCompleted(); ok {
		_spec.SetField(room.FieldIsCompleted, field.TypeBool, value)
		_node.IsCompleted = value
	}
	if nodes := rc.mutation.QuestionerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   room.QuestionerTable,
			Columns: []string{room.QuestionerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.user_own_rooms = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := rc.mutation.RespondentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   room.RespondentTable,
			Columns: room.RespondentPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := rc.mutation.CategoryIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   room.CategoryTable,
			Columns: []string{room.CategoryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(category.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.category_rooms = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := rc.mutation.MessagesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   room.MessagesTable,
			Columns: room.MessagesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(message.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := rc.mutation.ViewsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   room.ViewsTable,
			Columns: []string{room.ViewsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(view.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// RoomCreateBulk is the builder for creating many Room entities in bulk.
type RoomCreateBulk struct {
	config
	err      error
	builders []*RoomCreate
}

// Save creates the Room entities in the database.
func (rcb *RoomCreateBulk) Save(ctx context.Context) ([]*Room, error) {
	if rcb.err != nil {
		return nil, rcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(rcb.builders))
	nodes := make([]*Room, len(rcb.builders))
	mutators := make([]Mutator, len(rcb.builders))
	for i := range rcb.builders {
		func(i int, root context.Context) {
			builder := rcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*RoomMutation)
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
					_, err = mutators[i+1].Mutate(root, rcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, rcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, rcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (rcb *RoomCreateBulk) SaveX(ctx context.Context) []*Room {
	v, err := rcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rcb *RoomCreateBulk) Exec(ctx context.Context) error {
	_, err := rcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rcb *RoomCreateBulk) ExecX(ctx context.Context) {
	if err := rcb.Exec(ctx); err != nil {
		panic(err)
	}
}