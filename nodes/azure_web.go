package nodes

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type azurewebContainer struct {
	path  string
	attrs []attr.Attribute
}

var azureWeb = &webContainer{path: "assets/azure/web"}

func (c *azurewebContainer) MediaServices(opts ...attr.Attribute) *node.Node {
	return node.New("media-services", attr.AssetImage("assets/azure/web/media-services.png"), attr.Shape(attr.None))
}

func (c *azurewebContainer) NotificationHubNamespaces(opts ...attr.Attribute) *node.Node {
	return node.New("notification-hub-namespaces", attr.AssetImage("assets/azure/web/notification-hub-namespaces.png"), attr.Shape(attr.None))
}

func (c *azurewebContainer) Signalr(opts ...attr.Attribute) *node.Node {
	return node.New("signalr", attr.AssetImage("assets/azure/web/signalr.png"), attr.Shape(attr.None))
}

func (c *azurewebContainer) AppServiceDomains(opts ...attr.Attribute) *node.Node {
	return node.New("app-service-domains", attr.AssetImage("assets/azure/web/app-service-domains.png"), attr.Shape(attr.None))
}

func (c *azurewebContainer) AppServicePlans(opts ...attr.Attribute) *node.Node {
	return node.New("app-service-plans", attr.AssetImage("assets/azure/web/app-service-plans.png"), attr.Shape(attr.None))
}

func (c *azurewebContainer) AppServiceEnvironments(opts ...attr.Attribute) *node.Node {
	return node.New("app-service-environments", attr.AssetImage("assets/azure/web/app-service-environments.png"), attr.Shape(attr.None))
}

func (c *azurewebContainer) AppServices(opts ...attr.Attribute) *node.Node {
	return node.New("app-services", attr.AssetImage("assets/azure/web/app-services.png"), attr.Shape(attr.None))
}

func (c *azurewebContainer) Search(opts ...attr.Attribute) *node.Node {
	return node.New("search", attr.AssetImage("assets/azure/web/search.png"), attr.Shape(attr.None))
}

func (c *azurewebContainer) ApiConnections(opts ...attr.Attribute) *node.Node {
	return node.New("api-connections", attr.AssetImage("assets/azure/web/api-connections.png"), attr.Shape(attr.None))
}

func (c *azurewebContainer) AppServiceCertificates(opts ...attr.Attribute) *node.Node {
	return node.New("app-service-certificates", attr.AssetImage("assets/azure/web/app-service-certificates.png"), attr.Shape(attr.None))
}

func init() {
  const namespace = "azure.web"
  Register(namespace, "MediaServices", azureWeb.MediaServices)
  Register(namespace, "NotificationHubNamespaces", azureWeb.NotificationHubNamespaces)
  Register(namespace, "Signalr", azureWeb.Signalr)
  Register(namespace, "AppServiceDomains", azureWeb.AppServiceDomains)
  Register(namespace, "AppServicePlans", azureWeb.AppServicePlans)
  Register(namespace, "AppServiceEnvironments", azureWeb.AppServiceEnvironments)
  Register(namespace, "AppServices", azureWeb.AppServices)
  Register(namespace, "Search", azureWeb.Search)
  Register(namespace, "ApiConnections", azureWeb.ApiConnections)
  Register(namespace, "AppServiceCertificates", azureWeb.AppServiceCertificates)
  Register(namespace, "init", azureWeb.init)
}
