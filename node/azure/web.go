package azure

import "github.com/blushft/go-diagrams/node"

type webContainer struct {
	path string
	opts []node.Option
}

var Web = &webContainer{
	opts: node.OptionSet{node.Provider("azure"), node.Shape("none")},
	path: "assets/azure/web",
}

func (c *webContainer) AppServiceCertificates(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/azure/web/app-service-certificates.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *webContainer) AppServiceDomains(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/azure/web/app-service-domains.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *webContainer) MediaServices(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/azure/web/media-services.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *webContainer) Search(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/azure/web/search.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *webContainer) Signalr(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/azure/web/signalr.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *webContainer) ApiConnections(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/azure/web/api-connections.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *webContainer) AppServiceEnvironments(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/azure/web/app-service-environments.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *webContainer) AppServicePlans(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/azure/web/app-service-plans.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *webContainer) AppServices(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/azure/web/app-services.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *webContainer) NotificationHubNamespaces(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/azure/web/notification-hub-namespaces.png")}, c.opts, opts)
	return node.New(nopts...)
}
