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
	opts = append(opts, attr.AssetImage("assets/alibabacloud/iot/iot-link-wan.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("iot-link-wan", opts...)
}

func (c *alibabaCloudIotContainer) IotMobileConnectionPackage(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/alibabacloud/iot/iot-mobile-connection-package.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("iot-mobile-connection-package", opts...)
}

func (c *alibabaCloudIotContainer) IotPlatform(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/alibabacloud/iot/iot-platform.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("iot-platform", opts...)
}

func (c *alibabaCloudIotContainer) IotInternetDeviceId(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/alibabacloud/iot/iot-internet-device-id.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("iot-internet-device-id", opts...)
}

func init() {
  const namespace = "alibabacloud.iot"
  Register(namespace, "IotLinkWan", AlibabacloudAlibabaCloudIot.IotLinkWan)
  Register(namespace, "IotMobileConnectionPackage", AlibabacloudAlibabaCloudIot.IotMobileConnectionPackage)
  Register(namespace, "IotPlatform", AlibabacloudAlibabaCloudIot.IotPlatform)
  Register(namespace, "IotInternetDeviceId", AlibabacloudAlibabaCloudIot.IotInternetDeviceId)
}
