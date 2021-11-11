package events

import (
	"fmt"

	"github.com/negrel/paon/pdk/id"
)

// Type is the type of an Event
type Type id.ID

func (t Type) name() string {
	return typeName.Get(id.ID(t))
}

// String implements the fmt.Stringer interface.
func (t Type) String() string {
	return fmt.Sprintf("%v", t.name())
}

var typeName = id.NewStringMap()

// NewType returns a new Type that can be used for custom events
// type.
func NewType(name string) Type {
	t := id.New()
	typeName.Set(t, name)

	return Type(t)
}
