package nodes

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type azureiotContainer struct {
	path  string
	attrs []attr.Attribute
}

var azureIot = &azureiotContainer{path: "assets/azure/iot"}

func (c *azureiotContainer) Sphere(opts ...attr.Attribute) *node.Node {
	return node.New("sphere", attr.AssetImage("assets/azure/iot/sphere.png"), attr.Shape(attr.None))
}

func (c *azureiotContainer) TimeSeriesInsightsEnvironments(opts ...attr.Attribute) *node.Node {
	return node.New("time-series-insights-environments", attr.AssetImage("assets/azure/iot/time-series-insights-environments.png"), attr.Shape(attr.None))
}

func (c *azureiotContainer) TimeSeriesInsightsEventsSources(opts ...attr.Attribute) *node.Node {
	return node.New("time-series-insights-events-sources", attr.AssetImage("assets/azure/iot/time-series-insights-events-sources.png"), attr.Shape(attr.None))
}

func (c *azureiotContainer) DeviceProvisioningServices(opts ...attr.Attribute) *node.Node {
	return node.New("device-provisioning-services", attr.AssetImage("assets/azure/iot/device-provisioning-services.png"), attr.Shape(attr.None))
}

func (c *azureiotContainer) DigitalTwins(opts ...attr.Attribute) *node.Node {
	return node.New("digital-twins", attr.AssetImage("assets/azure/iot/digital-twins.png"), attr.Shape(attr.None))
}

func (c *azureiotContainer) IotCentralApplications(opts ...attr.Attribute) *node.Node {
	return node.New("iot-central-applications", attr.AssetImage("assets/azure/iot/iot-central-applications.png"), attr.Shape(attr.None))
}

func (c *azureiotContainer) IotHubSecurity(opts ...attr.Attribute) *node.Node {
	return node.New("iot-hub-security", attr.AssetImage("assets/azure/iot/iot-hub-security.png"), attr.Shape(attr.None))
}

func (c *azureiotContainer) IotHub(opts ...attr.Attribute) *node.Node {
	return node.New("iot-hub", attr.AssetImage("assets/azure/iot/iot-hub.png"), attr.Shape(attr.None))
}

func (c *azureiotContainer) Maps(opts ...attr.Attribute) *node.Node {
	return node.New("maps", attr.AssetImage("assets/azure/iot/maps.png"), attr.Shape(attr.None))
}

func (c *azureiotContainer) Windows10IotCoreServices(opts ...attr.Attribute) *node.Node {
	return node.New("windows-10-iot-core-services", attr.AssetImage("assets/azure/iot/windows-10-iot-core-services.png"), attr.Shape(attr.None))
}

func init() {
  const namespace = "azure.iot"
  Register(namespace, "Sphere", azureIot.Sphere)
  Register(namespace, "TimeSeriesInsightsEnvironments", azureIot.TimeSeriesInsightsEnvironments)
  Register(namespace, "TimeSeriesInsightsEventsSources", azureIot.TimeSeriesInsightsEventsSources)
  Register(namespace, "DeviceProvisioningServices", azureIot.DeviceProvisioningServices)
  Register(namespace, "DigitalTwins", azureIot.DigitalTwins)
  Register(namespace, "IotCentralApplications", azureIot.IotCentralApplications)
  Register(namespace, "IotHubSecurity", azureIot.IotHubSecurity)
  Register(namespace, "IotHub", azureIot.IotHub)
  Register(namespace, "Maps", azureIot.Maps)
  Register(namespace, "Windows10IotCoreServices", azureIot.Windows10IotCoreServices)
}
