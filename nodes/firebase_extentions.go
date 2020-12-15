package nodes

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type extentionsContainer struct {
	path  string
	attrs []attr.Attribute
}

var Extentions = &extentionsContainer{path: "assets/firebase/extentions"}

func (c *extentionsContainer) Extensions(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/firebase/extentions/extensions.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("extensions", opts...)
}
