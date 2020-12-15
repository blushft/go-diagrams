package nodes

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type computeContainer struct {
	path  string
	attrs []attr.Attribute
}

var OpenstackCompute =&computeContainer{path: "assets/openstack/compute"}

func (c *computeContainer) Nova(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/openstack/compute/nova.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("nova", opts...)
}

func (c *computeContainer) Qinling(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/openstack/compute/qinling.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("qinling", opts...)
}

func (c *computeContainer) Zun(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/openstack/compute/zun.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("zun", opts...)
}

func init() {
  const namespace = "openstack.compute"
  Register(namespace, "Nova", OpenstackCompute.Nova)
  Register(namespace, "Qinling", OpenstackCompute.Qinling)
  Register(namespace, "Zun", OpenstackCompute.Zun)
}
