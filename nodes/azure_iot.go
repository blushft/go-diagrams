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
	opts = append(opts, attr.AssetImage("assets/azure/iot/sphere.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("sphere", opts...)
}

func (c *azureiotContainer) TimeSeriesInsightsEnvironments(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/azure/iot/time-series-insights-environments.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("time-series-insights-environments", opts...)
}

func (c *azureiotContainer) TimeSeriesInsightsEventsSources(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/azure/iot/time-series-insights-events-sources.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("time-series-insights-events-sources", opts...)
}

func (c *azureiotContainer) DeviceProvisioningServices(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/azure/iot/device-provisioning-services.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("device-provisioning-services", opts...)
}

func (c *azureiotContainer) DigitalTwins(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/azure/iot/digital-twins.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("digital-twins", opts...)
}

func (c *azureiotContainer) IotCentralApplications(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/azure/iot/iot-central-applications.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("iot-central-applications", opts...)
}

func (c *azureiotContainer) IotHubSecurity(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/azure/iot/iot-hub-security.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("iot-hub-security", opts...)
}

func (c *azureiotContainer) IotHub(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/azure/iot/iot-hub.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("iot-hub", opts...)
}

func (c *azureiotContainer) Maps(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/azure/iot/maps.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("maps", opts...)
}

func (c *azureiotContainer) Windows10IotCoreServices(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/azure/iot/windows-10-iot-core-services.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("windows-10-iot-core-services", opts...)
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
