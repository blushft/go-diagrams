package alibabacloud

import "github.com/blushft/go-diagrams/diagram"

type iotContainer struct {
	path string
	opts []diagram.NodeOption
}

var Iot = &iotContainer{
	opts: diagram.OptionSet{diagram.Provider("alibabacloud"), diagram.NodeShape("none")},
	path: "assets/alibabacloud/iot",
}

func (c *iotContainer) IotInternetDeviceId(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/alibabacloud/iot/iot-internet-device-id.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *iotContainer) IotLinkWan(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/alibabacloud/iot/iot-link-wan.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *iotContainer) IotMobileConnectionPackage(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/alibabacloud/iot/iot-mobile-connection-package.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *iotContainer) IotPlatform(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/alibabacloud/iot/iot-platform.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
