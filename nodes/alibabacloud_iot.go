package nodes

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type iotContainer struct {
	path  string
	attrs []attr.Attribute
}

var Iot = &iotContainer{path: "assets/alibabacloud/iot"}

func (c *iotContainer) IotLinkWan(opts ...attr.Attribute) *node.Node {
	return node.New("iot-link-wan", attr.AssetImage("assets/alibabacloud/iot/iot-link-wan.png"), attr.Shape(attr.None))
}

func (c *iotContainer) IotMobileConnectionPackage(opts ...attr.Attribute) *node.Node {
	return node.New("iot-mobile-connection-package", attr.AssetImage("assets/alibabacloud/iot/iot-mobile-connection-package.png"), attr.Shape(attr.None))
}

func (c *iotContainer) IotPlatform(opts ...attr.Attribute) *node.Node {
	return node.New("iot-platform", attr.AssetImage("assets/alibabacloud/iot/iot-platform.png"), attr.Shape(attr.None))
}

func (c *iotContainer) IotInternetDeviceId(opts ...attr.Attribute) *node.Node {
	return node.New("iot-internet-device-id", attr.AssetImage("assets/alibabacloud/iot/iot-internet-device-id.png"), attr.Shape(attr.None))
}
