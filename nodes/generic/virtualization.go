package generic

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type virtualizationContainer struct {
	path  string
	attrs []attr.Attribute
}

var Virtualization = &virtualizationContainer{path: "assets/generic/virtualization"}

func (c *virtualizationContainer) Virtualbox(opts ...attr.Attribute) *node.Node {
	return node.New("virtualbox", attr.AssetImage("assets/generic/virtualization/virtualbox.png"), attr.Shape(attr.None))
}

func (c *virtualizationContainer) Vmware(opts ...attr.Attribute) *node.Node {
	return node.New("vmware", attr.AssetImage("assets/generic/virtualization/vmware.png"), attr.Shape(attr.None))
}

func (c *virtualizationContainer) Xen(opts ...attr.Attribute) *node.Node {
	return node.New("xen", attr.AssetImage("assets/generic/virtualization/xen.png"), attr.Shape(attr.None))
}
