package azure

import "github.com/blushft/go-diagrams/diagram"

type webContainer struct {
	path string
	opts []diagram.NodeOption
}

var Web = &webContainer{
	opts: diagram.OptionSet{diagram.Provider("azure"), diagram.NodeShape("none")},
	path: "assets/azure/web",
}

func (c *webContainer) NotificationHubNamespaces(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/web/notification-hub-namespaces.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *webContainer) Search(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/web/search.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *webContainer) Signalr(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/web/signalr.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *webContainer) ApiConnections(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/web/api-connections.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *webContainer) AppServiceCertificates(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/web/app-service-certificates.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *webContainer) AppServiceDomains(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/web/app-service-domains.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *webContainer) AppServices(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/web/app-services.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *webContainer) AppServiceEnvironments(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/web/app-service-environments.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *webContainer) AppServicePlans(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/web/app-service-plans.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *webContainer) MediaServices(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/web/media-services.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
