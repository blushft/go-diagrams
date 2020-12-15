package nodes

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type blankContainer struct {
	path  string
	attrs []attr.Attribute
}

var Blank = &blankContainer{path: "assets/generic/blank"}

func (c *blankContainer) Blank(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/generic/blank/blank.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("blank", opts...)
}

func init() {
	const namespace = "generic.blank"
	Register(namespace, "Blank", Blank.Blank)
}