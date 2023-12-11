// Code generated by ent, DO NOT EDIT.

package view

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/PUDDLEEE/puddleee_back/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.View {
	return predicate.View(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.View {
	return predicate.View(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.View {
	return predicate.View(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.View {
	return predicate.View(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.View {
	return predicate.View(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.View {
	return predicate.View(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.View {
	return predicate.View(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.View {
	return predicate.View(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.View {
	return predicate.View(sql.FieldLTE(FieldID, id))
}

// CreateTime applies equality check predicate on the "create_time" field. It's identical to CreateTimeEQ.
func CreateTime(v time.Time) predicate.View {
	return predicate.View(sql.FieldEQ(FieldCreateTime, v))
}

// UpdateTime applies equality check predicate on the "update_time" field. It's identical to UpdateTimeEQ.
func UpdateTime(v time.Time) predicate.View {
	return predicate.View(sql.FieldEQ(FieldUpdateTime, v))
}

// Filepath applies equality check predicate on the "filepath" field. It's identical to FilepathEQ.
func Filepath(v string) predicate.View {
	return predicate.View(sql.FieldEQ(FieldFilepath, v))
}

// CreateTimeEQ applies the EQ predicate on the "create_time" field.
func CreateTimeEQ(v time.Time) predicate.View {
	return predicate.View(sql.FieldEQ(FieldCreateTime, v))
}

// CreateTimeNEQ applies the NEQ predicate on the "create_time" field.
func CreateTimeNEQ(v time.Time) predicate.View {
	return predicate.View(sql.FieldNEQ(FieldCreateTime, v))
}

// CreateTimeIn applies the In predicate on the "create_time" field.
func CreateTimeIn(vs ...time.Time) predicate.View {
	return predicate.View(sql.FieldIn(FieldCreateTime, vs...))
}

// CreateTimeNotIn applies the NotIn predicate on the "create_time" field.
func CreateTimeNotIn(vs ...time.Time) predicate.View {
	return predicate.View(sql.FieldNotIn(FieldCreateTime, vs...))
}

// CreateTimeGT applies the GT predicate on the "create_time" field.
func CreateTimeGT(v time.Time) predicate.View {
	return predicate.View(sql.FieldGT(FieldCreateTime, v))
}

// CreateTimeGTE applies the GTE predicate on the "create_time" field.
func CreateTimeGTE(v time.Time) predicate.View {
	return predicate.View(sql.FieldGTE(FieldCreateTime, v))
}

// CreateTimeLT applies the LT predicate on the "create_time" field.
func CreateTimeLT(v time.Time) predicate.View {
	return predicate.View(sql.FieldLT(FieldCreateTime, v))
}

// CreateTimeLTE applies the LTE predicate on the "create_time" field.
func CreateTimeLTE(v time.Time) predicate.View {
	return predicate.View(sql.FieldLTE(FieldCreateTime, v))
}

// UpdateTimeEQ applies the EQ predicate on the "update_time" field.
func UpdateTimeEQ(v time.Time) predicate.View {
	return predicate.View(sql.FieldEQ(FieldUpdateTime, v))
}

// UpdateTimeNEQ applies the NEQ predicate on the "update_time" field.
func UpdateTimeNEQ(v time.Time) predicate.View {
	return predicate.View(sql.FieldNEQ(FieldUpdateTime, v))
}

// UpdateTimeIn applies the In predicate on the "update_time" field.
func UpdateTimeIn(vs ...time.Time) predicate.View {
	return predicate.View(sql.FieldIn(FieldUpdateTime, vs...))
}

// UpdateTimeNotIn applies the NotIn predicate on the "update_time" field.
func UpdateTimeNotIn(vs ...time.Time) predicate.View {
	return predicate.View(sql.FieldNotIn(FieldUpdateTime, vs...))
}

// UpdateTimeGT applies the GT predicate on the "update_time" field.
func UpdateTimeGT(v time.Time) predicate.View {
	return predicate.View(sql.FieldGT(FieldUpdateTime, v))
}

// UpdateTimeGTE applies the GTE predicate on the "update_time" field.
func UpdateTimeGTE(v time.Time) predicate.View {
	return predicate.View(sql.FieldGTE(FieldUpdateTime, v))
}

// UpdateTimeLT applies the LT predicate on the "update_time" field.
func UpdateTimeLT(v time.Time) predicate.View {
	return predicate.View(sql.FieldLT(FieldUpdateTime, v))
}

// UpdateTimeLTE applies the LTE predicate on the "update_time" field.
func UpdateTimeLTE(v time.Time) predicate.View {
	return predicate.View(sql.FieldLTE(FieldUpdateTime, v))
}

// FilepathEQ applies the EQ predicate on the "filepath" field.
func FilepathEQ(v string) predicate.View {
	return predicate.View(sql.FieldEQ(FieldFilepath, v))
}

// FilepathNEQ applies the NEQ predicate on the "filepath" field.
func FilepathNEQ(v string) predicate.View {
	return predicate.View(sql.FieldNEQ(FieldFilepath, v))
}

// FilepathIn applies the In predicate on the "filepath" field.
func FilepathIn(vs ...string) predicate.View {
	return predicate.View(sql.FieldIn(FieldFilepath, vs...))
}

// FilepathNotIn applies the NotIn predicate on the "filepath" field.
func FilepathNotIn(vs ...string) predicate.View {
	return predicate.View(sql.FieldNotIn(FieldFilepath, vs...))
}

// FilepathGT applies the GT predicate on the "filepath" field.
func FilepathGT(v string) predicate.View {
	return predicate.View(sql.FieldGT(FieldFilepath, v))
}

// FilepathGTE applies the GTE predicate on the "filepath" field.
func FilepathGTE(v string) predicate.View {
	return predicate.View(sql.FieldGTE(FieldFilepath, v))
}

// FilepathLT applies the LT predicate on the "filepath" field.
func FilepathLT(v string) predicate.View {
	return predicate.View(sql.FieldLT(FieldFilepath, v))
}

// FilepathLTE applies the LTE predicate on the "filepath" field.
func FilepathLTE(v string) predicate.View {
	return predicate.View(sql.FieldLTE(FieldFilepath, v))
}

// FilepathContains applies the Contains predicate on the "filepath" field.
func FilepathContains(v string) predicate.View {
	return predicate.View(sql.FieldContains(FieldFilepath, v))
}

// FilepathHasPrefix applies the HasPrefix predicate on the "filepath" field.
func FilepathHasPrefix(v string) predicate.View {
	return predicate.View(sql.FieldHasPrefix(FieldFilepath, v))
}

// FilepathHasSuffix applies the HasSuffix predicate on the "filepath" field.
func FilepathHasSuffix(v string) predicate.View {
	return predicate.View(sql.FieldHasSuffix(FieldFilepath, v))
}

// FilepathEqualFold applies the EqualFold predicate on the "filepath" field.
func FilepathEqualFold(v string) predicate.View {
	return predicate.View(sql.FieldEqualFold(FieldFilepath, v))
}

// FilepathContainsFold applies the ContainsFold predicate on the "filepath" field.
func FilepathContainsFold(v string) predicate.View {
	return predicate.View(sql.FieldContainsFold(FieldFilepath, v))
}

// HasRoom applies the HasEdge predicate on the "room" edge.
func HasRoom() predicate.View {
	return predicate.View(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, RoomTable, RoomColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasRoomWith applies the HasEdge predicate on the "room" edge with a given conditions (other predicates).
func HasRoomWith(preds ...predicate.Room) predicate.View {
	return predicate.View(func(s *sql.Selector) {
		step := newRoomStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.View) predicate.View {
	return predicate.View(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.View) predicate.View {
	return predicate.View(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.View) predicate.View {
	return predicate.View(sql.NotPredicates(p))
}
