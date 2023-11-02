package events

import (
	"github.com/negrel/paon/events"
	"github.com/negrel/paon/tree"
)

type baseNodeOption struct {
	*BaseNode
	nodeConstructor func(data interface{}) tree.Node
	data            interface{}
}

// NodeOption define an option for BaseNode.
type NodeOption func(*baseNodeOption)

// EventTarget returns a NodeOption that sets the events.Target that will be used by the BaseNode.
func EventTarget(t events.Target) NodeOption {
	return func(bno *baseNodeOption) {
		bno.target = target{
			Target: t,
		}
	}
}

// NodeConstructor returns a NodeOption that sets the internal tree.Node constructor used by the BaseNode.
func NodeConstructor(constructor func(data interface{}) tree.Node) NodeOption {
	return func(bno *baseNodeOption) {
		bno.nodeConstructor = constructor
	}
}

// Wrap returns a NodeOption that sets the internal data used by the BaseNode.
// This data is accessible through the tree.Node.Unwrap method.
func Wrap(data interface{}) NodeOption {
	return func(bno *baseNodeOption) {
		bno.data = data
	}
}
