package nodes

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type workloadprovisioningContainer struct {
	path  string
	attrs []attr.Attribute
}

var OpenstackWorkloadprovisioning =&workloadprovisioningContainer{path: "assets/openstack/workloadprovisioning"}

func (c *workloadprovisioningContainer) Sahara(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/openstack/workloadprovisioning/sahara.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("sahara", opts...)
}

func (c *workloadprovisioningContainer) Trove(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/openstack/workloadprovisioning/trove.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("trove", opts...)
}

func (c *workloadprovisioningContainer) Magnum(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/openstack/workloadprovisioning/magnum.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("magnum", opts...)
}

func init() {
  const namespace = "openstack.workloadprovisioning"
  Register(namespace, "Sahara", OpenstackWorkloadprovisioning.Sahara)
  Register(namespace, "Trove", OpenstackWorkloadprovisioning.Trove)
  Register(namespace, "Magnum", OpenstackWorkloadprovisioning.Magnum)
}
