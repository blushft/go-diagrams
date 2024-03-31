package azure

import "github.com/blushft/go-diagrams/diagram"

type integrationContainer struct {
	path string
	opts []diagram.NodeOption
}

var Integration = &integrationContainer{
	opts: diagram.OptionSet{diagram.Provider("azure"), diagram.NodeShape("none")},
	path: "assets/azure/integration",
}

func (c *integrationContainer) ApiForFhir(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/integration/api-for-fhir.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *integrationContainer) ApiManagement(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/integration/api-management.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *integrationContainer) AppConfiguration(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/integration/app-configuration.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *integrationContainer) ServiceBus(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/integration/service-bus.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *integrationContainer) EventGridDomains(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/integration/event-grid-domains.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *integrationContainer) EventGridSubscriptions(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/integration/event-grid-subscriptions.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *integrationContainer) LogicAppsCustomConnector(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/integration/logic-apps-custom-connector.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *integrationContainer) SendgridAccounts(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/integration/sendgrid-accounts.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *integrationContainer) SoftwareAsAService(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/integration/software-as-a-service.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *integrationContainer) DataCatalog(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/integration/data-catalog.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *integrationContainer) EventGridTopics(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/integration/event-grid-topics.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *integrationContainer) IntegrationAccounts(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/integration/integration-accounts.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *integrationContainer) LogicApps(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/integration/logic-apps.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *integrationContainer) IntegrationServiceEnvironments(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/integration/integration-service-environments.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *integrationContainer) ServiceBusRelays(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/integration/service-bus-relays.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *integrationContainer) ServiceCatalogManagedApplicationDefinitions(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/integration/service-catalog-managed-application-definitions.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *integrationContainer) StorsimpleDeviceManagers(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/integration/storsimple-device-managers.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
