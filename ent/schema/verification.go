package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"github.com/google/uuid"
)

// Verification holds the schema definition for the Verification entity.
type Verification struct {
	ent.Schema
}

func (Verification) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

// Fields of the Verification.
func (Verification) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("uuid", uuid.UUID{}).Unique(),
		field.String("code").MinLen(6).NotEmpty(),
		field.Time("expired_at").Default(time.Now().Add(10 * time.Minute)),
	}
}

// Edges of the Verification.
func (Verification) Edges() []ent.Edge {
	return nil
}
