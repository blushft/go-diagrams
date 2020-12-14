package nodes

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type containerservicesContainer struct {
	path  string
	attrs []attr.Attribute
}

var Containerservices = &containerservicesContainer{path: "assets/openstack/containerservices"}

func (c *containerservicesContainer) Kuryr(opts ...attr.Attribute) *node.Node {
	return node.New("kuryr", attr.AssetImage("assets/openstack/containerservices/kuryr.png"), attr.Shape(attr.None))
}
