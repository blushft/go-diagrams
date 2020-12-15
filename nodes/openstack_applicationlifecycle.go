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
	opts = append(opts, attr.AssetImage("assets/openstack/applicationlifecycle/murano.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("murano", opts...)
}

func (c *applicationlifecycleContainer) Solum(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/openstack/applicationlifecycle/solum.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("solum", opts...)
}

func (c *applicationlifecycleContainer) Freezer(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/openstack/applicationlifecycle/freezer.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("freezer", opts...)
}

func (c *applicationlifecycleContainer) Masakari(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/openstack/applicationlifecycle/masakari.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("masakari", opts...)
}

func init() {
  const namespace = "openstack.applicationlifecycle"
  Register(namespace, "Murano", OpenstackApplicationlifecycle.Murano)
  Register(namespace, "Solum", OpenstackApplicationlifecycle.Solum)
  Register(namespace, "Freezer", OpenstackApplicationlifecycle.Freezer)
  Register(namespace, "Masakari", OpenstackApplicationlifecycle.Masakari)
}
