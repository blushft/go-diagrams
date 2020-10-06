package azure

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type webContainer struct {
	path  string
	attrs []attr.Attribute
}

var Web = &webContainer{path: "assets/azure/web"}

func (c *webContainer) MediaServices(opts ...attr.Attribute) *node.Node {
	return node.New("media-services", attr.AssetImage("assets/azure/web/media-services.png"), attr.Shape(attr.None))
}

func (c *webContainer) NotificationHubNamespaces(opts ...attr.Attribute) *node.Node {
	return node.New("notification-hub-namespaces", attr.AssetImage("assets/azure/web/notification-hub-namespaces.png"), attr.Shape(attr.None))
}

func (c *webContainer) Signalr(opts ...attr.Attribute) *node.Node {
	return node.New("signalr", attr.AssetImage("assets/azure/web/signalr.png"), attr.Shape(attr.None))
}

func (c *webContainer) AppServiceDomains(opts ...attr.Attribute) *node.Node {
	return node.New("app-service-domains", attr.AssetImage("assets/azure/web/app-service-domains.png"), attr.Shape(attr.None))
}

func (c *webContainer) AppServicePlans(opts ...attr.Attribute) *node.Node {
	return node.New("app-service-plans", attr.AssetImage("assets/azure/web/app-service-plans.png"), attr.Shape(attr.None))
}

func (c *webContainer) AppServiceEnvironments(opts ...attr.Attribute) *node.Node {
	return node.New("app-service-environments", attr.AssetImage("assets/azure/web/app-service-environments.png"), attr.Shape(attr.None))
}

func (c *webContainer) AppServices(opts ...attr.Attribute) *node.Node {
	return node.New("app-services", attr.AssetImage("assets/azure/web/app-services.png"), attr.Shape(attr.None))
}

func (c *webContainer) Search(opts ...attr.Attribute) *node.Node {
	return node.New("search", attr.AssetImage("assets/azure/web/search.png"), attr.Shape(attr.None))
}

func (c *webContainer) ApiConnections(opts ...attr.Attribute) *node.Node {
	return node.New("api-connections", attr.AssetImage("assets/azure/web/api-connections.png"), attr.Shape(attr.None))
}

func (c *webContainer) AppServiceCertificates(opts ...attr.Attribute) *node.Node {
	return node.New("app-service-certificates", attr.AssetImage("assets/azure/web/app-service-certificates.png"), attr.Shape(attr.None))
}
