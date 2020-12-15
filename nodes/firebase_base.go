package nodes

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type baseContainer struct {
	path  string
	attrs []attr.Attribute
}

var Base = &baseContainer{path: "assets/firebase/base"}

func (c *baseContainer) Firebase(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/firebase/base/firebase.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("firebase", opts...)
}
