package nodes

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type azurewebContainer struct {
	path  string
	attrs []attr.Attribute
}

var AzureWeb = &azurewebContainer{path: "assets/azure/web"}

func (c *azurewebContainer) MediaServices(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/azure/web/media-services.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("media-services", opts...)
}

func (c *azurewebContainer) NotificationHubNamespaces(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/azure/web/notification-hub-namespaces.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("notification-hub-namespaces", opts...)
}

func (c *azurewebContainer) Signalr(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/azure/web/signalr.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("signalr", opts...)
}

func (c *azurewebContainer) AppServiceDomains(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/azure/web/app-service-domains.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("app-service-domains", opts...)
}

func (c *azurewebContainer) AppServicePlans(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/azure/web/app-service-plans.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("app-service-plans", opts...)
}

func (c *azurewebContainer) AppServiceEnvironments(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/azure/web/app-service-environments.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("app-service-environments", opts...)
}

func (c *azurewebContainer) AppServices(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/azure/web/app-services.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("app-services", opts...)
}

func (c *azurewebContainer) Search(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/azure/web/search.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("search", opts...)
}

func (c *azurewebContainer) ApiConnections(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/azure/web/api-connections.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("api-connections", opts...)
}

func (c *azurewebContainer) AppServiceCertificates(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/azure/web/app-service-certificates.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("app-service-certificates", opts...)
}

func init() {
  const namespace = "azure.web"
  Register(namespace, "MediaServices", AzureWeb.MediaServices)
  Register(namespace, "NotificationHubNamespaces", AzureWeb.NotificationHubNamespaces)
  Register(namespace, "Signalr", AzureWeb.Signalr)
  Register(namespace, "AppServiceDomains", AzureWeb.AppServiceDomains)
  Register(namespace, "AppServicePlans", AzureWeb.AppServicePlans)
  Register(namespace, "AppServiceEnvironments", AzureWeb.AppServiceEnvironments)
  Register(namespace, "AppServices", AzureWeb.AppServices)
  Register(namespace, "Search", AzureWeb.Search)
  Register(namespace, "ApiConnections", AzureWeb.ApiConnections)
  Register(namespace, "AppServiceCertificates", AzureWeb.AppServiceCertificates)
}
