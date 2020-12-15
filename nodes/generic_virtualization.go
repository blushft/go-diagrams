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
	opts = append(opts, attr.AssetImage("assets/generic/virtualization/vmware.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("vmware", opts...)
}

func (c *genericVirtualizationContainer) Xen(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/generic/virtualization/xen.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("xen", opts...)
}

func (c *genericVirtualizationContainer) Virtualbox(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/generic/virtualization/virtualbox.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("virtualbox", opts...)
}

func init() {
	const namespace = "generic.virtualization"
	Register(namespace, "Vmware", GenericVirtualization.Vmware)
	Register(namespace, "Xen", GenericVirtualization.Xen)
	Register(namespace, "Virtualbox", GenericVirtualization.Virtualbox)
}