package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// Message holds the schema definition for the Message entity.
type Message struct {
	ent.Schema
}

func (Message) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

// Fields of the Message.
func (Message) Fields() []ent.Field {
	return []ent.Field{
		field.String("payload").NotEmpty().Immutable(),
	}
}

// Edges of the Message.
func (Message) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("room", Room.Type).
			Ref("messages").Required(),
		edge.From("user", User.Type).
			Ref("messages").Unique().Required(),
	}
}
