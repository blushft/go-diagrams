package apps

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type mlopsContainer struct {
	path  string
	attrs []attr.Attribute
}

var Mlops = &mlopsContainer{path: "assets/apps/mlops"}

func (c *mlopsContainer) Polyaxon(opts ...attr.Attribute) *node.Node {
	return node.New("polyaxon", attr.AssetImage("assets/apps/mlops/polyaxon.png"), attr.Shape(attr.None))
}
