package azure

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type iotContainer struct {
	path  string
	attrs []attr.Attribute
}

var Iot = &iotContainer{path: "assets/azure/iot"}

func (c *iotContainer) DeviceProvisioningServices(opts ...attr.Attribute) *node.Node {
	return node.New("device-provisioning-services", attr.AssetImage("assets/azure/iot/device-provisioning-services.png"), attr.Shape(attr.None))
}

func (c *iotContainer) DigitalTwins(opts ...attr.Attribute) *node.Node {
	return node.New("digital-twins", attr.AssetImage("assets/azure/iot/digital-twins.png"), attr.Shape(attr.None))
}

func (c *iotContainer) IotHub(opts ...attr.Attribute) *node.Node {
	return node.New("iot-hub", attr.AssetImage("assets/azure/iot/iot-hub.png"), attr.Shape(attr.None))
}

func (c *iotContainer) Maps(opts ...attr.Attribute) *node.Node {
	return node.New("maps", attr.AssetImage("assets/azure/iot/maps.png"), attr.Shape(attr.None))
}

func (c *iotContainer) TimeSeriesInsightsEnvironments(opts ...attr.Attribute) *node.Node {
	return node.New("time-series-insights-environments", attr.AssetImage("assets/azure/iot/time-series-insights-environments.png"), attr.Shape(attr.None))
}

func (c *iotContainer) TimeSeriesInsightsEventsSources(opts ...attr.Attribute) *node.Node {
	return node.New("time-series-insights-events-sources", attr.AssetImage("assets/azure/iot/time-series-insights-events-sources.png"), attr.Shape(attr.None))
}

func (c *iotContainer) Windows10IotCoreServices(opts ...attr.Attribute) *node.Node {
	return node.New("windows-10-iot-core-services", attr.AssetImage("assets/azure/iot/windows-10-iot-core-services.png"), attr.Shape(attr.None))
}

func (c *iotContainer) IotCentralApplications(opts ...attr.Attribute) *node.Node {
	return node.New("iot-central-applications", attr.AssetImage("assets/azure/iot/iot-central-applications.png"), attr.Shape(attr.None))
}

func (c *iotContainer) IotHubSecurity(opts ...attr.Attribute) *node.Node {
	return node.New("iot-hub-security", attr.AssetImage("assets/azure/iot/iot-hub-security.png"), attr.Shape(attr.None))
}

func (c *iotContainer) Sphere(opts ...attr.Attribute) *node.Node {
	return node.New("sphere", attr.AssetImage("assets/azure/iot/sphere.png"), attr.Shape(attr.None))
}
