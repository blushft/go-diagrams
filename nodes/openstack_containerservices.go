package nodes

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type containerservicesContainer struct {
	path  string
	attrs []attr.Attribute
}

var OpenstackContainerservices =&containerservicesContainer{path: "assets/openstack/containerservices"}

func (c *containerservicesContainer) Kuryr(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/openstack/containerservices/kuryr.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("kuryr", opts...)
}

func init() {
  const namespace = "openstack.containerservices"
  Register(namespace, "Kuryr", OpenstackContainerservices.Kuryr)
}
