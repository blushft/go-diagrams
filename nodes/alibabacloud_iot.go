package nodes

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type alibabaCloudIotContainer struct {
	path  string
	attrs []attr.Attribute
}

var AlibabacloudAlibabaCloudIot =&alibabaCloudIotContainer{path: "assets/alibabacloud/iot"}

func (c *alibabaCloudIotContainer) IotLinkWan(opts ...attr.Attribute) *node.Node {
	return node.New("iot-link-wan", attr.AssetImage("assets/alibabacloud/iot/iot-link-wan.png"), attr.Shape(attr.None))
}

func (c *alibabaCloudIotContainer) IotMobileConnectionPackage(opts ...attr.Attribute) *node.Node {
	return node.New("iot-mobile-connection-package", attr.AssetImage("assets/alibabacloud/iot/iot-mobile-connection-package.png"), attr.Shape(attr.None))
}

func (c *alibabaCloudIotContainer) IotPlatform(opts ...attr.Attribute) *node.Node {
	return node.New("iot-platform", attr.AssetImage("assets/alibabacloud/iot/iot-platform.png"), attr.Shape(attr.None))
}

func (c *alibabaCloudIotContainer) IotInternetDeviceId(opts ...attr.Attribute) *node.Node {
	return node.New("iot-internet-device-id", attr.AssetImage("assets/alibabacloud/iot/iot-internet-device-id.png"), attr.Shape(attr.None))
}

func init() {
  const namespace = "alibabacloud.iot"
  Register(namespace, "IotLinkWan", AlibabacloudAlibabaCloudIot.IotLinkWan)
  Register(namespace, "IotMobileConnectionPackage", AlibabacloudAlibabaCloudIot.IotMobileConnectionPackage)
  Register(namespace, "IotPlatform", AlibabacloudAlibabaCloudIot.IotPlatform)
  Register(namespace, "IotInternetDeviceId", AlibabacloudAlibabaCloudIot.IotInternetDeviceId)
}
