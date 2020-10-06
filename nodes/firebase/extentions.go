package firebase

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
	return node.New("extensions", attr.AssetImage("assets/firebase/extentions/extensions.png"), attr.Shape(attr.None))
}
