package nodes

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type azuremobileContainer struct {
	path  string
	attrs []attr.Attribute
}

var azureMobile = &mobileContainer{path: "assets/azure/mobile"}

func (c *azuremobileContainer) NotificationHubs(opts ...attr.Attribute) *node.Node {
	return node.New("notification-hubs", attr.AssetImage("assets/azure/mobile/notification-hubs.png"), attr.Shape(attr.None))
}

func (c *azuremobileContainer) AppServiceMobile(opts ...attr.Attribute) *node.Node {
	return node.New("app-service---mobile", attr.AssetImage("assets/azure/mobile/app-service---mobile.png"), attr.Shape(attr.None))
}

func (c *azuremobileContainer) MobileEngagement(opts ...attr.Attribute) *node.Node {
	return node.New("mobile-engagement", attr.AssetImage("assets/azure/mobile/mobile-engagement.png"), attr.Shape(attr.None))
}

func init() {
  const namespace = "azure.mobile"
  Register(namespace, "NotificationHubs", azureMobile.NotificationHubs)
  Register(namespace, "AppServiceMobile", azureMobile.AppServiceMobile)
  Register(namespace, "MobileEngagement", azureMobile.MobileEngagement)
  Register(namespace, "init", azureMobile.init)
}
