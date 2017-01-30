package models

import (
	"encoding/json"
	"fmt"
	"time"
)

// AppKey is a models.Key implementation specifically for Todo models
type TodoKey struct {
	str string
}

// NewAppKey creates and returns a new AppKey with the given todo name
func NewAppKey(name string) *TodoKey {
	return &TodoKey{str: name}
}

// String is the fmt.Stringer interface implementation
func (t *TodoKey) String() string {
	return t.str
}

// App is a Model implementation for an application in the PaaS
type Todo struct {
	Name      string    `json:"name"`
	Completed bool      `json:"completed"`
	Due       time.Time `json:"due"`
}

// MarshalBinary is the encoding.BinaryMarshaler interface implementation
func (t *Todo) MarshalBinary() ([]byte, error) {
	return json.Marshal(t)
}

// UnmarshalBinary is the encoding.BinaryUnmarshaler interface implementation
func (t *Todo) UnmarshalBinary(b []byte) error {
	return json.Unmarshal(b, t)
}

// Set is the Model interface implementation
func (t *Todo) Set(m Model) error {
	todo, ok := m.(*Todo)
	if !ok {
		return fmt.Errorf("given model %+v was not an *Todo", m)
	}
	*t = *todo
	return nil
}
