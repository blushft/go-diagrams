package azure

import "github.com/blushft/go-diagrams/node"

type integrationContainer struct {
	path string
	opts []node.Option
}

var Integration = &integrationContainer{
	opts: node.OptionSet{node.Provider("azure"), node.Shape("none")},
	path: "assets/azure/integration",
}

func (c *integrationContainer) ApiForFhir(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/azure/integration/api-for-fhir.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *integrationContainer) EventGridSubscriptions(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/azure/integration/event-grid-subscriptions.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *integrationContainer) SendgridAccounts(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/azure/integration/sendgrid-accounts.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *integrationContainer) ServiceBus(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/azure/integration/service-bus.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *integrationContainer) StorsimpleDeviceManagers(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/azure/integration/storsimple-device-managers.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *integrationContainer) AppConfiguration(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/azure/integration/app-configuration.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *integrationContainer) IntegrationServiceEnvironments(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/azure/integration/integration-service-environments.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *integrationContainer) LogicApps(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/azure/integration/logic-apps.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *integrationContainer) ServiceBusRelays(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/azure/integration/service-bus-relays.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *integrationContainer) ApiManagement(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/azure/integration/api-management.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *integrationContainer) DataCatalog(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/azure/integration/data-catalog.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *integrationContainer) EventGridDomains(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/azure/integration/event-grid-domains.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *integrationContainer) EventGridTopics(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/azure/integration/event-grid-topics.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *integrationContainer) IntegrationAccounts(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/azure/integration/integration-accounts.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *integrationContainer) LogicAppsCustomConnector(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/azure/integration/logic-apps-custom-connector.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *integrationContainer) ServiceCatalogManagedApplicationDefinitions(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/azure/integration/service-catalog-managed-application-definitions.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *integrationContainer) SoftwareAsAService(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/azure/integration/software-as-a-service.png")}, c.opts, opts)
	return node.New(nopts...)
}
