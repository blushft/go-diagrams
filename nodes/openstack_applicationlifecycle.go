package nodes

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type applicationlifecycleContainer struct {
	path  string
	attrs []attr.Attribute
}

var OpenstackApplicationlifecycle =&applicationlifecycleContainer{path: "assets/openstack/applicationlifecycle"}

func (c *applicationlifecycleContainer) Murano(opts ...attr.Attribute) *node.Node {
	return node.New("murano", attr.AssetImage("assets/openstack/applicationlifecycle/murano.png"), attr.Shape(attr.None))
}

func (c *applicationlifecycleContainer) Solum(opts ...attr.Attribute) *node.Node {
	return node.New("solum", attr.AssetImage("assets/openstack/applicationlifecycle/solum.png"), attr.Shape(attr.None))
}

func (c *applicationlifecycleContainer) Freezer(opts ...attr.Attribute) *node.Node {
	return node.New("freezer", attr.AssetImage("assets/openstack/applicationlifecycle/freezer.png"), attr.Shape(attr.None))
}

func (c *applicationlifecycleContainer) Masakari(opts ...attr.Attribute) *node.Node {
	return node.New("masakari", attr.AssetImage("assets/openstack/applicationlifecycle/masakari.png"), attr.Shape(attr.None))
}

func init() {
  const namespace = "openstack.applicationlifecycle"
  Register(namespace, "Murano", OpenstackApplicationlifecycle.Murano)
  Register(namespace, "Solum", OpenstackApplicationlifecycle.Solum)
  Register(namespace, "Freezer", OpenstackApplicationlifecycle.Freezer)
  Register(namespace, "Masakari", OpenstackApplicationlifecycle.Masakari)
}
