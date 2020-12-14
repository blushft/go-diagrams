package nodes

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type genericDeviceContainer struct {
	path  string
	attrs []attr.Attribute
}

var GenericDevice = &genericDeviceContainer{path: "assets/generic/device"}

func (c *genericDeviceContainer) Tablet(opts ...attr.Attribute) *node.Node {
	return node.New("tablet", attr.AssetImage("assets/generic/device/tablet.png"), attr.Shape(attr.None))
}

func (c *genericDeviceContainer) Mobile(opts ...attr.Attribute) *node.Node {
	return node.New("mobile", attr.AssetImage("assets/generic/device/mobile.png"), attr.Shape(attr.None))
}

func init() {
	const namespace = "generic.device"
	Register(namespace, "Tablet", GenericDevice.Tablet)
	Register(namespace, "Mobile", GenericDevice.Mobile)
}
