package tree

import (
	"github.com/negrel/debuggo/pkg/assert"
	"github.com/negrel/debuggo/pkg/log"
)

type Root interface {
	ParentNode

	register(Node)
	unregister(Node)
	GetNode(NodeID) Node
}

var _ Root = &root{}

type root struct {
	*parentNode

	widgets map[NodeID]Node
}

func (r *root) GetNode(nID NodeID) Node {
	return r.widgets[nID]
}

func NewRoot(children Node) Root {
	assert.NotNil(children, "node must be non-nil to be the root")

	r := newRoot()
	r.parentNode.node.root = r
	r.appendChildNode(children)

	return r
}

func newRoot() *root {
	return &root{
		parentNode: newParent("root"),
		widgets:    make(map[NodeID]Node),
	}
}

func (r *root) register(child Node) {
	assert.NotNil(child, "can't register a non-nil child")

	id := child.ID()
	_, isAlreadyRegistered := r.widgets[id]
	if isAlreadyRegistered {
		log.Infoln(child, "is already registered")
	}

	log.Debugln("registering", child)
	r.widgets[id] = child
}

func (r *root) unregister(child Node) {
	assert.NotNil(child, "can't unregister a non-nil child")

	id := child.ID()
	_, isRegistered := r.widgets[id]
	if !isRegistered {
		log.Debugln(child, "is not registered")
	}

	delete(r.widgets, id)
}
