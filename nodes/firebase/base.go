package firebase

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
	return node.New("firebase", attr.AssetImage("assets/firebase/base/firebase.png"), attr.Shape(attr.None))
}
