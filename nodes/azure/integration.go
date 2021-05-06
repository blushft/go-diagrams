package azure

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type integrationContainer struct {
	path  string
	attrs []attr.Attribute
}

var Integration = &integrationContainer{path: "assets/azure/integration"}

func (c *integrationContainer) ApiForFhir(opts ...attr.Attribute) *node.Node {
	return node.New("api-for-fhir", attr.AssetImage("assets/azure/integration/api-for-fhir.png"), attr.Shape(attr.None))
}

func (c *integrationContainer) EventGridSubscriptions(opts ...attr.Attribute) *node.Node {
	return node.New("event-grid-subscriptions", attr.AssetImage("assets/azure/integration/event-grid-subscriptions.png"), attr.Shape(attr.None))
}

func (c *integrationContainer) EventGridTopics(opts ...attr.Attribute) *node.Node {
	return node.New("event-grid-topics", attr.AssetImage("assets/azure/integration/event-grid-topics.png"), attr.Shape(attr.None))
}

func (c *integrationContainer) LogicAppsCustomConnector(opts ...attr.Attribute) *node.Node {
	return node.New("logic-apps-custom-connector", attr.AssetImage("assets/azure/integration/logic-apps-custom-connector.png"), attr.Shape(attr.None))
}

func (c *integrationContainer) ServiceBus(opts ...attr.Attribute) *node.Node {
	return node.New("service-bus", attr.AssetImage("assets/azure/integration/service-bus.png"), attr.Shape(attr.None))
}

func (c *integrationContainer) ServiceCatalogManagedApplicationDefinitions(opts ...attr.Attribute) *node.Node {
	return node.New("service-catalog-managed-application-definitions", attr.AssetImage("assets/azure/integration/service-catalog-managed-application-definitions.png"), attr.Shape(attr.None))
}

func (c *integrationContainer) SoftwareAsAService(opts ...attr.Attribute) *node.Node {
	return node.New("software-as-a-service", attr.AssetImage("assets/azure/integration/software-as-a-service.png"), attr.Shape(attr.None))
}

func (c *integrationContainer) StorsimpleDeviceManagers(opts ...attr.Attribute) *node.Node {
	return node.New("storsimple-device-managers", attr.AssetImage("assets/azure/integration/storsimple-device-managers.png"), attr.Shape(attr.None))
}

func (c *integrationContainer) DataCatalog(opts ...attr.Attribute) *node.Node {
	return node.New("data-catalog", attr.AssetImage("assets/azure/integration/data-catalog.png"), attr.Shape(attr.None))
}

func (c *integrationContainer) ServiceBusRelays(opts ...attr.Attribute) *node.Node {
	return node.New("service-bus-relays", attr.AssetImage("assets/azure/integration/service-bus-relays.png"), attr.Shape(attr.None))
}

func (c *integrationContainer) ApiManagement(opts ...attr.Attribute) *node.Node {
	return node.New("api-management", attr.AssetImage("assets/azure/integration/api-management.png"), attr.Shape(attr.None))
}

func (c *integrationContainer) IntegrationServiceEnvironments(opts ...attr.Attribute) *node.Node {
	return node.New("integration-service-environments", attr.AssetImage("assets/azure/integration/integration-service-environments.png"), attr.Shape(attr.None))
}

func (c *integrationContainer) AppConfiguration(opts ...attr.Attribute) *node.Node {
	return node.New("app-configuration", attr.AssetImage("assets/azure/integration/app-configuration.png"), attr.Shape(attr.None))
}

func (c *integrationContainer) EventGridDomains(opts ...attr.Attribute) *node.Node {
	return node.New("event-grid-domains", attr.AssetImage("assets/azure/integration/event-grid-domains.png"), attr.Shape(attr.None))
}

func (c *integrationContainer) IntegrationAccounts(opts ...attr.Attribute) *node.Node {
	return node.New("integration-accounts", attr.AssetImage("assets/azure/integration/integration-accounts.png"), attr.Shape(attr.None))
}

func (c *integrationContainer) LogicApps(opts ...attr.Attribute) *node.Node {
	return node.New("logic-apps", attr.AssetImage("assets/azure/integration/logic-apps.png"), attr.Shape(attr.None))
}

func (c *integrationContainer) SendgridAccounts(opts ...attr.Attribute) *node.Node {
	return node.New("sendgrid-accounts", attr.AssetImage("assets/azure/integration/sendgrid-accounts.png"), attr.Shape(attr.None))
}
