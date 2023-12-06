// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/PUDDLEEE/puddleee_back/ent/message"
	"github.com/PUDDLEEE/puddleee_back/ent/user"
)

// Message is the model entity for the Message schema.
type Message struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
	// UpdateTime holds the value of the "update_time" field.
	UpdateTime time.Time `json:"update_time,omitempty"`
	// Payload holds the value of the "payload" field.
	Payload string `json:"payload,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the MessageQuery when eager-loading is set.
	Edges         MessageEdges `json:"edges"`
	user_messages *int
	selectValues  sql.SelectValues
}

// MessageEdges holds the relations/edges for other nodes in the graph.
type MessageEdges struct {
	// Room holds the value of the room edge.
	Room []*Room `json:"room,omitempty"`
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// RoomOrErr returns the Room value or an error if the edge
// was not loaded in eager-loading.
func (e MessageEdges) RoomOrErr() ([]*Room, error) {
	if e.loadedTypes[0] {
		return e.Room, nil
	}
	return nil, &NotLoadedError{edge: "room"}
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e MessageEdges) UserOrErr() (*User, error) {
	if e.loadedTypes[1] {
		if e.User == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.User, nil
	}
	return nil, &NotLoadedError{edge: "user"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Message) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case message.FieldID:
			values[i] = new(sql.NullInt64)
		case message.FieldPayload:
			values[i] = new(sql.NullString)
		case message.FieldCreateTime, message.FieldUpdateTime:
			values[i] = new(sql.NullTime)
		case message.ForeignKeys[0]: // user_messages
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Message fields.
func (m *Message) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case message.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			m.ID = int(value.Int64)
		case message.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				m.CreateTime = value.Time
			}
		case message.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				m.UpdateTime = value.Time
			}
		case message.FieldPayload:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field payload", values[i])
			} else if value.Valid {
				m.Payload = value.String
			}
		case message.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field user_messages", value)
			} else if value.Valid {
				m.user_messages = new(int)
				*m.user_messages = int(value.Int64)
			}
		default:
			m.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Message.
// This includes values selected through modifiers, order, etc.
func (m *Message) Value(name string) (ent.Value, error) {
	return m.selectValues.Get(name)
}

// QueryRoom queries the "room" edge of the Message entity.
func (m *Message) QueryRoom() *RoomQuery {
	return NewMessageClient(m.config).QueryRoom(m)
}

// QueryUser queries the "user" edge of the Message entity.
func (m *Message) QueryUser() *UserQuery {
	return NewMessageClient(m.config).QueryUser(m)
}

// Update returns a builder for updating this Message.
// Note that you need to call Message.Unwrap() before calling this method if this Message
// was returned from a transaction, and the transaction was committed or rolled back.
func (m *Message) Update() *MessageUpdateOne {
	return NewMessageClient(m.config).UpdateOne(m)
}

// Unwrap unwraps the Message entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (m *Message) Unwrap() *Message {
	_tx, ok := m.config.driver.(*txDriver)
	if !ok {
		panic("ent: Message is not a transactional entity")
	}
	m.config.driver = _tx.drv
	return m
}

// String implements the fmt.Stringer.
func (m *Message) String() string {
	var builder strings.Builder
	builder.WriteString("Message(")
	builder.WriteString(fmt.Sprintf("id=%v, ", m.ID))
	builder.WriteString("create_time=")
	builder.WriteString(m.CreateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("update_time=")
	builder.WriteString(m.UpdateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("payload=")
	builder.WriteString(m.Payload)
	builder.WriteByte(')')
	return builder.String()
}

// Messages is a parsable slice of Message.
type Messages []*Message
