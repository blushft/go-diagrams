package openstack

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type workloadprovisioningContainer struct {
	path  string
	attrs []attr.Attribute
}

var Workloadprovisioning = &workloadprovisioningContainer{path: "assets/openstack/workloadprovisioning"}

func (c *workloadprovisioningContainer) Magnum(opts ...attr.Attribute) *node.Node {
	return node.New("magnum", attr.AssetImage("assets/openstack/workloadprovisioning/magnum.png"), attr.Shape(attr.None))
}

func (c *workloadprovisioningContainer) Sahara(opts ...attr.Attribute) *node.Node {
	return node.New("sahara", attr.AssetImage("assets/openstack/workloadprovisioning/sahara.png"), attr.Shape(attr.None))
}

func (c *workloadprovisioningContainer) Trove(opts ...attr.Attribute) *node.Node {
	return node.New("trove", attr.AssetImage("assets/openstack/workloadprovisioning/trove.png"), attr.Shape(attr.None))
}
