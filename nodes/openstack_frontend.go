package nodes

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type frontendContainer struct {
	path  string
	attrs []attr.Attribute
}

var Frontend = &frontendContainer{path: "assets/openstack/frontend"}

func (c *frontendContainer) Horizon(opts ...attr.Attribute) *node.Node {
	return node.New("horizon", attr.AssetImage("assets/openstack/frontend/horizon.png"), attr.Shape(attr.None))
}
