package nodes

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type baremetalContainer struct {
	path  string
	attrs []attr.Attribute
}

var OpenstackBaremetal =&baremetalContainer{path: "assets/openstack/baremetal"}

func (c *baremetalContainer) Cyborg(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/openstack/baremetal/cyborg.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("cyborg", opts...)
}

func (c *baremetalContainer) Ironic(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/openstack/baremetal/ironic.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("ironic", opts...)
}

func init() {
  const namespace = "openstack.baremetal"
  Register(namespace, "Cyborg", OpenstackBaremetal.Cyborg)
  Register(namespace, "Ironic", OpenstackBaremetal.Ironic)
}
