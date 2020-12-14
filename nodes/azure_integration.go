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
	return node.New("service-catalog-managed-application-definitions", attr.AssetImage("assets/azure/integration/service-catalog-managed-application-definitions.png"), attr.Shape(attr.None))
}

func (c *azureintegrationContainer) SoftwareAsAService(opts ...attr.Attribute) *node.Node {
	return node.New("software-as-a-service", attr.AssetImage("assets/azure/integration/software-as-a-service.png"), attr.Shape(attr.None))
}

func (c *azureintegrationContainer) ApiManagement(opts ...attr.Attribute) *node.Node {
	return node.New("api-management", attr.AssetImage("assets/azure/integration/api-management.png"), attr.Shape(attr.None))
}

func (c *azureintegrationContainer) EventGridTopics(opts ...attr.Attribute) *node.Node {
	return node.New("event-grid-topics", attr.AssetImage("assets/azure/integration/event-grid-topics.png"), attr.Shape(attr.None))
}

func (c *azureintegrationContainer) IntegrationServiceEnvironments(opts ...attr.Attribute) *node.Node {
	return node.New("integration-service-environments", attr.AssetImage("assets/azure/integration/integration-service-environments.png"), attr.Shape(attr.None))
}

func (c *azureintegrationContainer) SendgridAccounts(opts ...attr.Attribute) *node.Node {
	return node.New("sendgrid-accounts", attr.AssetImage("assets/azure/integration/sendgrid-accounts.png"), attr.Shape(attr.None))
}

func (c *azureintegrationContainer) ApiForFhir(opts ...attr.Attribute) *node.Node {
	return node.New("api-for-fhir", attr.AssetImage("assets/azure/integration/api-for-fhir.png"), attr.Shape(attr.None))
}

func (c *azureintegrationContainer) LogicApps(opts ...attr.Attribute) *node.Node {
	return node.New("logic-apps", attr.AssetImage("assets/azure/integration/logic-apps.png"), attr.Shape(attr.None))
}

func (c *azureintegrationContainer) ServiceBusRelays(opts ...attr.Attribute) *node.Node {
	return node.New("service-bus-relays", attr.AssetImage("assets/azure/integration/service-bus-relays.png"), attr.Shape(attr.None))
}

func (c *azureintegrationContainer) IntegrationAccounts(opts ...attr.Attribute) *node.Node {
	return node.New("integration-accounts", attr.AssetImage("assets/azure/integration/integration-accounts.png"), attr.Shape(attr.None))
}

func (c *azureintegrationContainer) AppConfiguration(opts ...attr.Attribute) *node.Node {
	return node.New("app-configuration", attr.AssetImage("assets/azure/integration/app-configuration.png"), attr.Shape(attr.None))
}

func (c *azureintegrationContainer) DataCatalog(opts ...attr.Attribute) *node.Node {
	return node.New("data-catalog", attr.AssetImage("assets/azure/integration/data-catalog.png"), attr.Shape(attr.None))
}

func (c *azureintegrationContainer) EventGridDomains(opts ...attr.Attribute) *node.Node {
	return node.New("event-grid-domains", attr.AssetImage("assets/azure/integration/event-grid-domains.png"), attr.Shape(attr.None))
}

func (c *azureintegrationContainer) EventGridSubscriptions(opts ...attr.Attribute) *node.Node {
	return node.New("event-grid-subscriptions", attr.AssetImage("assets/azure/integration/event-grid-subscriptions.png"), attr.Shape(attr.None))
}

func (c *azureintegrationContainer) LogicAppsCustomConnector(opts ...attr.Attribute) *node.Node {
	return node.New("logic-apps-custom-connector", attr.AssetImage("assets/azure/integration/logic-apps-custom-connector.png"), attr.Shape(attr.None))
}

func (c *azureintegrationContainer) ServiceBus(opts ...attr.Attribute) *node.Node {
	return node.New("service-bus", attr.AssetImage("assets/azure/integration/service-bus.png"), attr.Shape(attr.None))
}

func (c *azureintegrationContainer) StorsimpleDeviceManagers(opts ...attr.Attribute) *node.Node {
	return node.New("storsimple-device-managers", attr.AssetImage("assets/azure/integration/storsimple-device-managers.png"), attr.Shape(attr.None))
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
