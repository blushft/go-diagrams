package nodes

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type baremetalContainer struct {
	path  string
	attrs []attr.Attribute
}

var Baremetal = &baremetalContainer{path: "assets/openstack/baremetal"}

func (c *baremetalContainer) Cyborg(opts ...attr.Attribute) *node.Node {
	return node.New("cyborg", attr.AssetImage("assets/openstack/baremetal/cyborg.png"), attr.Shape(attr.None))
}

func (c *baremetalContainer) Ironic(opts ...attr.Attribute) *node.Node {
	return node.New("ironic", attr.AssetImage("assets/openstack/baremetal/ironic.png"), attr.Shape(attr.None))
}
