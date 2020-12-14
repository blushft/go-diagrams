package nodes

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type apiproxiesContainer struct {
	path  string
	attrs []attr.Attribute
}

var OpenstackApiproxies =&apiproxiesContainer{path: "assets/openstack/apiproxies"}

func (c *apiproxiesContainer) Ec2Api(opts ...attr.Attribute) *node.Node {
	return node.New("ec2api", attr.AssetImage("assets/openstack/apiproxies/ec2api.png"), attr.Shape(attr.None))
}

func init() {
  const namespace = "openstack.apiproxies"
  Register(namespace, "Ec2Api", OpenstackApiproxies.Ec2Api)
}
