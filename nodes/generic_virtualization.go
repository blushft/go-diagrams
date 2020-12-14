package nodes

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type genericVirtualizationContainer struct {
	path  string
	attrs []attr.Attribute
}

var GenericVirtualization = &genericVirtualizationContainer{path: "assets/generic/virtualization"}

func (c *genericVirtualizationContainer) Vmware(opts ...attr.Attribute) *node.Node {
	return node.New("vmware", attr.AssetImage("assets/generic/virtualization/vmware.png"), attr.Shape(attr.None))
}

func (c *genericVirtualizationContainer) Xen(opts ...attr.Attribute) *node.Node {
	return node.New("xen", attr.AssetImage("assets/generic/virtualization/xen.png"), attr.Shape(attr.None))
}

func (c *genericVirtualizationContainer) Virtualbox(opts ...attr.Attribute) *node.Node {
	return node.New("virtualbox", attr.AssetImage("assets/generic/virtualization/virtualbox.png"), attr.Shape(attr.None))
}

func init() {
	const namespace = "generic.virtualization"
	Register(namespace, "Vmware", GenericVirtualization.Vmware)
	Register(namespace, "Xen", GenericVirtualization.Xen)
	Register(namespace, "Virtualbox", GenericVirtualization.Virtualbox)
}