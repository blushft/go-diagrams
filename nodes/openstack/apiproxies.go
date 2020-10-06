package openstack

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type apiproxiesContainer struct {
	path  string
	attrs []attr.Attribute
}

var Apiproxies = &apiproxiesContainer{path: "assets/openstack/apiproxies"}

func (c *apiproxiesContainer) Ec2Api(opts ...attr.Attribute) *node.Node {
	return node.New("ec2api", attr.AssetImage("assets/openstack/apiproxies/ec2api.png"), attr.Shape(attr.None))
}
