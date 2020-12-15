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
	opts = append(opts, attr.AssetImage("assets/openstack/apiproxies/ec2api.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("ec2api", opts...)
}

func init() {
  const namespace = "openstack.apiproxies"
  Register(namespace, "Ec2Api", OpenstackApiproxies.Ec2Api)
}
