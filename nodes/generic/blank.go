package generic

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
	return node.New("blank", attr.AssetImage("assets/generic/blank/blank.png"), attr.Shape(attr.None))
}
