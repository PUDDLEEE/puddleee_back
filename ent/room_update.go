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
	"github.com/PUDDLEEE/puddleee_back/ent/category"
	"github.com/PUDDLEEE/puddleee_back/ent/message"
	"github.com/PUDDLEEE/puddleee_back/ent/predicate"
	"github.com/PUDDLEEE/puddleee_back/ent/room"
	"github.com/PUDDLEEE/puddleee_back/ent/user"
	"github.com/PUDDLEEE/puddleee_back/ent/view"
)

// RoomUpdate is the builder for updating Room entities.
type RoomUpdate struct {
	config
	hooks    []Hook
	mutation *RoomMutation
}

// Where appends a list predicates to the RoomUpdate builder.
func (ru *RoomUpdate) Where(ps ...predicate.Room) *RoomUpdate {
	ru.mutation.Where(ps...)
	return ru
}

// SetUpdateTime sets the "update_time" field.
func (ru *RoomUpdate) SetUpdateTime(t time.Time) *RoomUpdate {
	ru.mutation.SetUpdateTime(t)
	return ru
}

// SetTitle sets the "title" field.
func (ru *RoomUpdate) SetTitle(s string) *RoomUpdate {
	ru.mutation.SetTitle(s)
	return ru
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (ru *RoomUpdate) SetNillableTitle(s *string) *RoomUpdate {
	if s != nil {
		ru.SetTitle(*s)
	}
	return ru
}

// SetIsCompleted sets the "is_completed" field.
func (ru *RoomUpdate) SetIsCompleted(b bool) *RoomUpdate {
	ru.mutation.SetIsCompleted(b)
	return ru
}

// SetNillableIsCompleted sets the "is_completed" field if the given value is not nil.
func (ru *RoomUpdate) SetNillableIsCompleted(b *bool) *RoomUpdate {
	if b != nil {
		ru.SetIsCompleted(*b)
	}
	return ru
}

// SetQuestionerID sets the "questioner" edge to the User entity by ID.
func (ru *RoomUpdate) SetQuestionerID(id int) *RoomUpdate {
	ru.mutation.SetQuestionerID(id)
	return ru
}

// SetNillableQuestionerID sets the "questioner" edge to the User entity by ID if the given value is not nil.
func (ru *RoomUpdate) SetNillableQuestionerID(id *int) *RoomUpdate {
	if id != nil {
		ru = ru.SetQuestionerID(*id)
	}
	return ru
}

// SetQuestioner sets the "questioner" edge to the User entity.
func (ru *RoomUpdate) SetQuestioner(u *User) *RoomUpdate {
	return ru.SetQuestionerID(u.ID)
}

// AddRespondentIDs adds the "respondent" edge to the User entity by IDs.
func (ru *RoomUpdate) AddRespondentIDs(ids ...int) *RoomUpdate {
	ru.mutation.AddRespondentIDs(ids...)
	return ru
}

// AddRespondent adds the "respondent" edges to the User entity.
func (ru *RoomUpdate) AddRespondent(u ...*User) *RoomUpdate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return ru.AddRespondentIDs(ids...)
}

// SetCategoryID sets the "category" edge to the Category entity by ID.
func (ru *RoomUpdate) SetCategoryID(id int) *RoomUpdate {
	ru.mutation.SetCategoryID(id)
	return ru
}

// SetNillableCategoryID sets the "category" edge to the Category entity by ID if the given value is not nil.
func (ru *RoomUpdate) SetNillableCategoryID(id *int) *RoomUpdate {
	if id != nil {
		ru = ru.SetCategoryID(*id)
	}
	return ru
}

// SetCategory sets the "category" edge to the Category entity.
func (ru *RoomUpdate) SetCategory(c *Category) *RoomUpdate {
	return ru.SetCategoryID(c.ID)
}

// AddMessageIDs adds the "messages" edge to the Message entity by IDs.
func (ru *RoomUpdate) AddMessageIDs(ids ...int) *RoomUpdate {
	ru.mutation.AddMessageIDs(ids...)
	return ru
}

// AddMessages adds the "messages" edges to the Message entity.
func (ru *RoomUpdate) AddMessages(m ...*Message) *RoomUpdate {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return ru.AddMessageIDs(ids...)
}

// AddViewIDs adds the "views" edge to the View entity by IDs.
func (ru *RoomUpdate) AddViewIDs(ids ...int) *RoomUpdate {
	ru.mutation.AddViewIDs(ids...)
	return ru
}

// AddViews adds the "views" edges to the View entity.
func (ru *RoomUpdate) AddViews(v ...*View) *RoomUpdate {
	ids := make([]int, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return ru.AddViewIDs(ids...)
}

// Mutation returns the RoomMutation object of the builder.
func (ru *RoomUpdate) Mutation() *RoomMutation {
	return ru.mutation
}

// ClearQuestioner clears the "questioner" edge to the User entity.
func (ru *RoomUpdate) ClearQuestioner() *RoomUpdate {
	ru.mutation.ClearQuestioner()
	return ru
}

// ClearRespondent clears all "respondent" edges to the User entity.
func (ru *RoomUpdate) ClearRespondent() *RoomUpdate {
	ru.mutation.ClearRespondent()
	return ru
}

// RemoveRespondentIDs removes the "respondent" edge to User entities by IDs.
func (ru *RoomUpdate) RemoveRespondentIDs(ids ...int) *RoomUpdate {
	ru.mutation.RemoveRespondentIDs(ids...)
	return ru
}

// RemoveRespondent removes "respondent" edges to User entities.
func (ru *RoomUpdate) RemoveRespondent(u ...*User) *RoomUpdate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return ru.RemoveRespondentIDs(ids...)
}

// ClearCategory clears the "category" edge to the Category entity.
func (ru *RoomUpdate) ClearCategory() *RoomUpdate {
	ru.mutation.ClearCategory()
	return ru
}

// ClearMessages clears all "messages" edges to the Message entity.
func (ru *RoomUpdate) ClearMessages() *RoomUpdate {
	ru.mutation.ClearMessages()
	return ru
}

// RemoveMessageIDs removes the "messages" edge to Message entities by IDs.
func (ru *RoomUpdate) RemoveMessageIDs(ids ...int) *RoomUpdate {
	ru.mutation.RemoveMessageIDs(ids...)
	return ru
}

// RemoveMessages removes "messages" edges to Message entities.
func (ru *RoomUpdate) RemoveMessages(m ...*Message) *RoomUpdate {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return ru.RemoveMessageIDs(ids...)
}

// ClearViews clears all "views" edges to the View entity.
func (ru *RoomUpdate) ClearViews() *RoomUpdate {
	ru.mutation.ClearViews()
	return ru
}

// RemoveViewIDs removes the "views" edge to View entities by IDs.
func (ru *RoomUpdate) RemoveViewIDs(ids ...int) *RoomUpdate {
	ru.mutation.RemoveViewIDs(ids...)
	return ru
}

// RemoveViews removes "views" edges to View entities.
func (ru *RoomUpdate) RemoveViews(v ...*View) *RoomUpdate {
	ids := make([]int, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return ru.RemoveViewIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ru *RoomUpdate) Save(ctx context.Context) (int, error) {
	ru.defaults()
	return withHooks(ctx, ru.sqlSave, ru.mutation, ru.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ru *RoomUpdate) SaveX(ctx context.Context) int {
	affected, err := ru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ru *RoomUpdate) Exec(ctx context.Context) error {
	_, err := ru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ru *RoomUpdate) ExecX(ctx context.Context) {
	if err := ru.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ru *RoomUpdate) defaults() {
	if _, ok := ru.mutation.UpdateTime(); !ok {
		v := room.UpdateDefaultUpdateTime()
		ru.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ru *RoomUpdate) check() error {
	if v, ok := ru.mutation.Title(); ok {
		if err := room.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "Room.title": %w`, err)}
		}
	}
	return nil
}

func (ru *RoomUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := ru.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(room.Table, room.Columns, sqlgraph.NewFieldSpec(room.FieldID, field.TypeInt))
	if ps := ru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ru.mutation.UpdateTime(); ok {
		_spec.SetField(room.FieldUpdateTime, field.TypeTime, value)
	}
	if value, ok := ru.mutation.Title(); ok {
		_spec.SetField(room.FieldTitle, field.TypeString, value)
	}
	if value, ok := ru.mutation.IsCompleted(); ok {
		_spec.SetField(room.FieldIsCompleted, field.TypeBool, value)
	}
	if ru.mutation.QuestionerCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.QuestionerIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ru.mutation.RespondentCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.RemovedRespondentIDs(); len(nodes) > 0 && !ru.mutation.RespondentCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.RespondentIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ru.mutation.CategoryCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.CategoryIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ru.mutation.MessagesCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.RemovedMessagesIDs(); len(nodes) > 0 && !ru.mutation.MessagesCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.MessagesIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ru.mutation.ViewsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.RemovedViewsIDs(); len(nodes) > 0 && !ru.mutation.ViewsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.ViewsIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{room.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ru.mutation.done = true
	return n, nil
}

// RoomUpdateOne is the builder for updating a single Room entity.
type RoomUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *RoomMutation
}

// SetUpdateTime sets the "update_time" field.
func (ruo *RoomUpdateOne) SetUpdateTime(t time.Time) *RoomUpdateOne {
	ruo.mutation.SetUpdateTime(t)
	return ruo
}

// SetTitle sets the "title" field.
func (ruo *RoomUpdateOne) SetTitle(s string) *RoomUpdateOne {
	ruo.mutation.SetTitle(s)
	return ruo
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (ruo *RoomUpdateOne) SetNillableTitle(s *string) *RoomUpdateOne {
	if s != nil {
		ruo.SetTitle(*s)
	}
	return ruo
}

// SetIsCompleted sets the "is_completed" field.
func (ruo *RoomUpdateOne) SetIsCompleted(b bool) *RoomUpdateOne {
	ruo.mutation.SetIsCompleted(b)
	return ruo
}

// SetNillableIsCompleted sets the "is_completed" field if the given value is not nil.
func (ruo *RoomUpdateOne) SetNillableIsCompleted(b *bool) *RoomUpdateOne {
	if b != nil {
		ruo.SetIsCompleted(*b)
	}
	return ruo
}

// SetQuestionerID sets the "questioner" edge to the User entity by ID.
func (ruo *RoomUpdateOne) SetQuestionerID(id int) *RoomUpdateOne {
	ruo.mutation.SetQuestionerID(id)
	return ruo
}

// SetNillableQuestionerID sets the "questioner" edge to the User entity by ID if the given value is not nil.
func (ruo *RoomUpdateOne) SetNillableQuestionerID(id *int) *RoomUpdateOne {
	if id != nil {
		ruo = ruo.SetQuestionerID(*id)
	}
	return ruo
}

// SetQuestioner sets the "questioner" edge to the User entity.
func (ruo *RoomUpdateOne) SetQuestioner(u *User) *RoomUpdateOne {
	return ruo.SetQuestionerID(u.ID)
}

// AddRespondentIDs adds the "respondent" edge to the User entity by IDs.
func (ruo *RoomUpdateOne) AddRespondentIDs(ids ...int) *RoomUpdateOne {
	ruo.mutation.AddRespondentIDs(ids...)
	return ruo
}

// AddRespondent adds the "respondent" edges to the User entity.
func (ruo *RoomUpdateOne) AddRespondent(u ...*User) *RoomUpdateOne {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return ruo.AddRespondentIDs(ids...)
}

// SetCategoryID sets the "category" edge to the Category entity by ID.
func (ruo *RoomUpdateOne) SetCategoryID(id int) *RoomUpdateOne {
	ruo.mutation.SetCategoryID(id)
	return ruo
}

// SetNillableCategoryID sets the "category" edge to the Category entity by ID if the given value is not nil.
func (ruo *RoomUpdateOne) SetNillableCategoryID(id *int) *RoomUpdateOne {
	if id != nil {
		ruo = ruo.SetCategoryID(*id)
	}
	return ruo
}

// SetCategory sets the "category" edge to the Category entity.
func (ruo *RoomUpdateOne) SetCategory(c *Category) *RoomUpdateOne {
	return ruo.SetCategoryID(c.ID)
}

// AddMessageIDs adds the "messages" edge to the Message entity by IDs.
func (ruo *RoomUpdateOne) AddMessageIDs(ids ...int) *RoomUpdateOne {
	ruo.mutation.AddMessageIDs(ids...)
	return ruo
}

// AddMessages adds the "messages" edges to the Message entity.
func (ruo *RoomUpdateOne) AddMessages(m ...*Message) *RoomUpdateOne {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return ruo.AddMessageIDs(ids...)
}

// AddViewIDs adds the "views" edge to the View entity by IDs.
func (ruo *RoomUpdateOne) AddViewIDs(ids ...int) *RoomUpdateOne {
	ruo.mutation.AddViewIDs(ids...)
	return ruo
}

// AddViews adds the "views" edges to the View entity.
func (ruo *RoomUpdateOne) AddViews(v ...*View) *RoomUpdateOne {
	ids := make([]int, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return ruo.AddViewIDs(ids...)
}

// Mutation returns the RoomMutation object of the builder.
func (ruo *RoomUpdateOne) Mutation() *RoomMutation {
	return ruo.mutation
}

// ClearQuestioner clears the "questioner" edge to the User entity.
func (ruo *RoomUpdateOne) ClearQuestioner() *RoomUpdateOne {
	ruo.mutation.ClearQuestioner()
	return ruo
}

// ClearRespondent clears all "respondent" edges to the User entity.
func (ruo *RoomUpdateOne) ClearRespondent() *RoomUpdateOne {
	ruo.mutation.ClearRespondent()
	return ruo
}

// RemoveRespondentIDs removes the "respondent" edge to User entities by IDs.
func (ruo *RoomUpdateOne) RemoveRespondentIDs(ids ...int) *RoomUpdateOne {
	ruo.mutation.RemoveRespondentIDs(ids...)
	return ruo
}

// RemoveRespondent removes "respondent" edges to User entities.
func (ruo *RoomUpdateOne) RemoveRespondent(u ...*User) *RoomUpdateOne {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return ruo.RemoveRespondentIDs(ids...)
}

// ClearCategory clears the "category" edge to the Category entity.
func (ruo *RoomUpdateOne) ClearCategory() *RoomUpdateOne {
	ruo.mutation.ClearCategory()
	return ruo
}

// ClearMessages clears all "messages" edges to the Message entity.
func (ruo *RoomUpdateOne) ClearMessages() *RoomUpdateOne {
	ruo.mutation.ClearMessages()
	return ruo
}

// RemoveMessageIDs removes the "messages" edge to Message entities by IDs.
func (ruo *RoomUpdateOne) RemoveMessageIDs(ids ...int) *RoomUpdateOne {
	ruo.mutation.RemoveMessageIDs(ids...)
	return ruo
}

// RemoveMessages removes "messages" edges to Message entities.
func (ruo *RoomUpdateOne) RemoveMessages(m ...*Message) *RoomUpdateOne {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return ruo.RemoveMessageIDs(ids...)
}

// ClearViews clears all "views" edges to the View entity.
func (ruo *RoomUpdateOne) ClearViews() *RoomUpdateOne {
	ruo.mutation.ClearViews()
	return ruo
}

// RemoveViewIDs removes the "views" edge to View entities by IDs.
func (ruo *RoomUpdateOne) RemoveViewIDs(ids ...int) *RoomUpdateOne {
	ruo.mutation.RemoveViewIDs(ids...)
	return ruo
}

// RemoveViews removes "views" edges to View entities.
func (ruo *RoomUpdateOne) RemoveViews(v ...*View) *RoomUpdateOne {
	ids := make([]int, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return ruo.RemoveViewIDs(ids...)
}

// Where appends a list predicates to the RoomUpdate builder.
func (ruo *RoomUpdateOne) Where(ps ...predicate.Room) *RoomUpdateOne {
	ruo.mutation.Where(ps...)
	return ruo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ruo *RoomUpdateOne) Select(field string, fields ...string) *RoomUpdateOne {
	ruo.fields = append([]string{field}, fields...)
	return ruo
}

// Save executes the query and returns the updated Room entity.
func (ruo *RoomUpdateOne) Save(ctx context.Context) (*Room, error) {
	ruo.defaults()
	return withHooks(ctx, ruo.sqlSave, ruo.mutation, ruo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ruo *RoomUpdateOne) SaveX(ctx context.Context) *Room {
	node, err := ruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ruo *RoomUpdateOne) Exec(ctx context.Context) error {
	_, err := ruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ruo *RoomUpdateOne) ExecX(ctx context.Context) {
	if err := ruo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ruo *RoomUpdateOne) defaults() {
	if _, ok := ruo.mutation.UpdateTime(); !ok {
		v := room.UpdateDefaultUpdateTime()
		ruo.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ruo *RoomUpdateOne) check() error {
	if v, ok := ruo.mutation.Title(); ok {
		if err := room.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "Room.title": %w`, err)}
		}
	}
	return nil
}

func (ruo *RoomUpdateOne) sqlSave(ctx context.Context) (_node *Room, err error) {
	if err := ruo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(room.Table, room.Columns, sqlgraph.NewFieldSpec(room.FieldID, field.TypeInt))
	id, ok := ruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Room.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, room.FieldID)
		for _, f := range fields {
			if !room.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != room.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ruo.mutation.UpdateTime(); ok {
		_spec.SetField(room.FieldUpdateTime, field.TypeTime, value)
	}
	if value, ok := ruo.mutation.Title(); ok {
		_spec.SetField(room.FieldTitle, field.TypeString, value)
	}
	if value, ok := ruo.mutation.IsCompleted(); ok {
		_spec.SetField(room.FieldIsCompleted, field.TypeBool, value)
	}
	if ruo.mutation.QuestionerCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.QuestionerIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ruo.mutation.RespondentCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.RemovedRespondentIDs(); len(nodes) > 0 && !ruo.mutation.RespondentCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.RespondentIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ruo.mutation.CategoryCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.CategoryIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ruo.mutation.MessagesCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.RemovedMessagesIDs(); len(nodes) > 0 && !ruo.mutation.MessagesCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.MessagesIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ruo.mutation.ViewsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.RemovedViewsIDs(); len(nodes) > 0 && !ruo.mutation.ViewsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.ViewsIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Room{config: ruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{room.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	ruo.mutation.done = true
	return _node, nil
}
