package nodes

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
	opts = append(opts, attr.AssetImage("assets/apps/etl/embulk.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("embulk", opts...)
}

func init() {
	const namespace = "apps.etl"
	Register(namespace, "Embulk", Etl.Embulk)
}
