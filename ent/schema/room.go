package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// Room holds the schema definition for the Room entity.
type Room struct {
	ent.Schema
}

func (Room) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

// Fields of the Room.
func (Room) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").NotEmpty(),
		field.Bool("is_completed").Default(false),
	}
}

// Edges of the Room.
func (Room) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("questioner", User.Type).
			Ref("own_rooms").
			Unique(),
		edge.From("respondent", User.Type).
			Ref("participant_rooms"),
		edge.From("category", Category.Type).
			Ref("rooms").Unique(),
		edge.To("messages", Message.Type),
		edge.To("views", View.Type),
	}
}
