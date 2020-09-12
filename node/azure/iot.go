package azure

import "github.com/blushft/go-diagrams/node"

type iotContainer struct {
	path string
	opts []node.Option
}

var Iot = &iotContainer{
	opts: node.OptionSet{node.Provider("azure"), node.Shape("none")},
	path: "assets/azure/iot",
}

func (c *iotContainer) IotHub(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/azure/iot/iot-hub.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *iotContainer) Maps(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/azure/iot/maps.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *iotContainer) Sphere(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/azure/iot/sphere.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *iotContainer) TimeSeriesInsightsEnvironments(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/azure/iot/time-series-insights-environments.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *iotContainer) TimeSeriesInsightsEventsSources(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/azure/iot/time-series-insights-events-sources.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *iotContainer) DeviceProvisioningServices(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/azure/iot/device-provisioning-services.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *iotContainer) IotCentralApplications(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/azure/iot/iot-central-applications.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *iotContainer) Windows10IotCoreServices(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/azure/iot/windows-10-iot-core-services.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *iotContainer) DigitalTwins(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/azure/iot/digital-twins.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *iotContainer) IotHubSecurity(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/azure/iot/iot-hub-security.png")}, c.opts, opts)
	return node.New(nopts...)
}
