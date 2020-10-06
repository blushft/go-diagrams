package apps

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type etlContainer struct {
	path  string
	attrs []attr.Attribute
}

var Etl = &etlContainer{path: "assets/apps/etl"}

func (c *etlContainer) Embulk(opts ...attr.Attribute) *node.Node {
	return node.New("embulk", attr.AssetImage("assets/apps/etl/embulk.png"), attr.Shape(attr.None))
}
