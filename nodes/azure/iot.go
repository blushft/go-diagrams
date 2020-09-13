package azure

import "github.com/blushft/go-diagrams/diagram"

type iotContainer struct {
	path string
	opts []diagram.NodeOption
}

var Iot = &iotContainer{
	opts: diagram.OptionSet{diagram.Provider("azure"), diagram.NodeShape("none")},
	path: "assets/azure/iot",
}

func (c *iotContainer) Windows10IotCoreServices(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/iot/windows-10-iot-core-services.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *iotContainer) IotHub(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/iot/iot-hub.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *iotContainer) Maps(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/iot/maps.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *iotContainer) TimeSeriesInsightsEnvironments(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/iot/time-series-insights-environments.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *iotContainer) TimeSeriesInsightsEventsSources(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/iot/time-series-insights-events-sources.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *iotContainer) Sphere(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/iot/sphere.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *iotContainer) DeviceProvisioningServices(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/iot/device-provisioning-services.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *iotContainer) DigitalTwins(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/iot/digital-twins.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *iotContainer) IotCentralApplications(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/iot/iot-central-applications.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *iotContainer) IotHubSecurity(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/iot/iot-hub-security.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
