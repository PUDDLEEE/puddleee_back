package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// View holds the schema definition for the View entity.
type View struct {
	ent.Schema
}

func (View) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

// Fields of the View.
func (View) Fields() []ent.Field {
	return []ent.Field{
		field.String("filepath").NotEmpty(),
	}
}

// Edges of the View.
func (View) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("room", Room.Type).
			Ref("views").
			Unique().Required(),
	}
}
