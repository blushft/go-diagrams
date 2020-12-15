package nodes

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
	opts = append(opts, attr.AssetImage("assets/apps/mlops/polyaxon.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("polyaxon", opts...)
}

func init() {
	const namespace = "apps.mlops"
	Register(namespace, "Polyaxon", Mlops.Polyaxon)
}