package tree

import (
	"errors"

	"github.com/negrel/paon/pkg/pdk/events"
	"github.com/negrel/paon/pkg/pdk/id"
)

var _ Node = &leafNode{}

type leafNode struct {
	events.Target

	id    id.ID
	data  interface{}
	stage LifeCycleStage

	next     Node
	previous Node
	parent   Node
}

// NewLeafNode returns a Node object that can't have children.
// The given data is stored in the node an returned on Unwrap call.
func NewLeafNode(data interface{}) Node {
	return newLeafNode(data)
}

func newLeafNode(data interface{}) *leafNode {
	ln := &leafNode{
		Target: events.NewTarget(),
		id:     id.Make(),
		data:   data,
	}

	// Update internal stage value
	ln.Target.AddEventListener(LifeCycleEventListener(func(event LifeCycleEvent) {
		ln.stage = event.Stage
	}))

	return ln
}

func (ln *leafNode) Stage() LifeCycleStage {
	return ln.stage
}

func (ln *leafNode) Unwrap() interface{} {
	return ln.data
}

func (ln *leafNode) ID() id.ID {
	return ln.id
}

func (ln *leafNode) IsSame(other Node) bool {
	if other == nil {
		return false
	}

	return ln.ID() == other.ID()
}

func (ln *leafNode) Next() Node {
	return ln.next
}

func (ln *leafNode) SetNext(next Node) {
	ln.next = next
}

func (ln *leafNode) Previous() Node {
	return ln.previous
}

func (ln *leafNode) SetPrevious(previous Node) {
	ln.previous = previous
}

func (ln *leafNode) Parent() Node {
	return ln.parent
}

func (ln *leafNode) SetParent(parent Node) {
	ln.parent = parent
}

func (ln *leafNode) Root() Node {
	if ln.parent != nil {
		return ln.parent.Root()
	}

	return nil
}

func (ln *leafNode) FirstChild() Node {
	return nil
}

func (ln *leafNode) LastChild() Node {
	return nil
}

func (ln *leafNode) AppendChild(_ Node) error {
	return errors.New("leaf node don't have children")
}

func (ln *leafNode) InsertBefore(_, _ Node) error {
	return errors.New("leaf node don't have children")
}

func (ln *leafNode) RemoveChild(Node) error {
	return errors.New("leaf node don't have children")
}

func (ln *leafNode) IsAncestorOf(other Node) bool {
	if other == nil {
		return false
	}

	return other.IsDescendantOf(ln)
}

func (ln *leafNode) IsDescendantOf(parent Node) bool {
	if parent == nil {
		return false
	}

	var node Node = ln
	for node != nil {
		if node.IsSame(parent) {
			return true
		}

		node = node.Parent()
	}

	return false
}

// String implements the fmt.Stringer interface.
func (ln *leafNode) String() string {
	return string(ln.id)
}
