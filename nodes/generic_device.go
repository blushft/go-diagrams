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
	opts = append(opts, attr.AssetImage("assets/generic/device/tablet.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("tablet", opts...)
}

func (c *genericDeviceContainer) Mobile(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/generic/device/mobile.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("mobile", opts...)
}

func init() {
	const namespace = "generic.device"
	Register(namespace, "Tablet", GenericDevice.Tablet)
	Register(namespace, "Mobile", GenericDevice.Mobile)
}
