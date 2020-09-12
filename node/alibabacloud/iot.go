package alibabacloud

import "github.com/blushft/go-diagrams/node"

type iotContainer struct {
	path string
	opts []node.Option
}

var Iot = &iotContainer{
	opts: node.OptionSet{node.Provider("alibabacloud"), node.Shape("none")},
	path: "assets/alibabacloud/iot",
}

func (c *iotContainer) IotPlatform(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/alibabacloud/iot/iot-platform.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *iotContainer) IotInternetDeviceId(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/alibabacloud/iot/iot-internet-device-id.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *iotContainer) IotLinkWan(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/alibabacloud/iot/iot-link-wan.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *iotContainer) IotMobileConnectionPackage(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/alibabacloud/iot/iot-mobile-connection-package.png")}, c.opts, opts)
	return node.New(nopts...)
}
