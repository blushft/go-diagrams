package k8s

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type groupContainer struct {
	path  string
	attrs []attr.Attribute
}

var Group = &groupContainer{path: "assets/k8s/group"}

func (c *groupContainer) Ns(opts ...attr.Attribute) *node.Node {
	return node.New("ns", attr.AssetImage("assets/k8s/group/ns.png"), attr.Shape(attr.None))
}
