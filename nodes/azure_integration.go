package nodes

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type azureintegrationContainer struct {
	path  string
	attrs []attr.Attribute
}

var azureIntegration = &azureintegrationContainer{path: "assets/azure/integration"}

func (c *azureintegrationContainer) ServiceCatalogManagedApplicationDefinitions(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/azure/integration/service-catalog-managed-application-definitions.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("service-catalog-managed-application-definitions", opts...)
}

func (c *azureintegrationContainer) SoftwareAsAService(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/azure/integration/software-as-a-service.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("software-as-a-service", opts...)
}

func (c *azureintegrationContainer) ApiManagement(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/azure/integration/api-management.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("api-management", opts...)
}

func (c *azureintegrationContainer) EventGridTopics(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/azure/integration/event-grid-topics.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("event-grid-topics", opts...)
}

func (c *azureintegrationContainer) IntegrationServiceEnvironments(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/azure/integration/integration-service-environments.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("integration-service-environments", opts...)
}

func (c *azureintegrationContainer) SendgridAccounts(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/azure/integration/sendgrid-accounts.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("sendgrid-accounts", opts...)
}

func (c *azureintegrationContainer) ApiForFhir(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/azure/integration/api-for-fhir.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("api-for-fhir", opts...)
}

func (c *azureintegrationContainer) LogicApps(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/azure/integration/logic-apps.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("logic-apps", opts...)
}

func (c *azureintegrationContainer) ServiceBusRelays(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/azure/integration/service-bus-relays.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("service-bus-relays", opts...)
}

func (c *azureintegrationContainer) IntegrationAccounts(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/azure/integration/integration-accounts.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("integration-accounts", opts...)
}

func (c *azureintegrationContainer) AppConfiguration(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/azure/integration/app-configuration.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("app-configuration", opts...)
}

func (c *azureintegrationContainer) DataCatalog(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/azure/integration/data-catalog.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("data-catalog", opts...)
}

func (c *azureintegrationContainer) EventGridDomains(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/azure/integration/event-grid-domains.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("event-grid-domains", opts...)
}

func (c *azureintegrationContainer) EventGridSubscriptions(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/azure/integration/event-grid-subscriptions.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("event-grid-subscriptions", opts...)
}

func (c *azureintegrationContainer) LogicAppsCustomConnector(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/azure/integration/logic-apps-custom-connector.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("logic-apps-custom-connector", opts...)
}

func (c *azureintegrationContainer) ServiceBus(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/azure/integration/service-bus.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("service-bus", opts...)
}

func (c *azureintegrationContainer) StorsimpleDeviceManagers(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/azure/integration/storsimple-device-managers.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("storsimple-device-managers", opts...)
}

func init() {
  const namespace = "azure.integration"
  Register(namespace, "ServiceCatalogManagedApplicationDefinitions", azureIntegration.ServiceCatalogManagedApplicationDefinitions)
  Register(namespace, "SoftwareAsAService", azureIntegration.SoftwareAsAService)
  Register(namespace, "ApiManagement", azureIntegration.ApiManagement)
  Register(namespace, "EventGridTopics", azureIntegration.EventGridTopics)
  Register(namespace, "IntegrationServiceEnvironments", azureIntegration.IntegrationServiceEnvironments)
  Register(namespace, "SendgridAccounts", azureIntegration.SendgridAccounts)
  Register(namespace, "ApiForFhir", azureIntegration.ApiForFhir)
  Register(namespace, "LogicApps", azureIntegration.LogicApps)
  Register(namespace, "ServiceBusRelays", azureIntegration.ServiceBusRelays)
  Register(namespace, "IntegrationAccounts", azureIntegration.IntegrationAccounts)
  Register(namespace, "AppConfiguration", azureIntegration.AppConfiguration)
  Register(namespace, "DataCatalog", azureIntegration.DataCatalog)
  Register(namespace, "EventGridDomains", azureIntegration.EventGridDomains)
  Register(namespace, "EventGridSubscriptions", azureIntegration.EventGridSubscriptions)
  Register(namespace, "LogicAppsCustomConnector", azureIntegration.LogicAppsCustomConnector)
  Register(namespace, "ServiceBus", azureIntegration.ServiceBus)
  Register(namespace, "StorsimpleDeviceManagers", azureIntegration.StorsimpleDeviceManagers)
}
