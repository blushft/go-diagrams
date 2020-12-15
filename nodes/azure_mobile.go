package nodes

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type azuremobileContainer struct {
	path  string
	attrs []attr.Attribute
}

var azureMobile = &azuremobileContainer{path: "assets/azure/mobile"}

func (c *azuremobileContainer) NotificationHubs(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/azure/mobile/notification-hubs.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("notification-hubs", opts...)
}

func (c *azuremobileContainer) AppServiceMobile(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/azure/mobile/app-service---mobile.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("app-service---mobile", opts...)
}

func (c *azuremobileContainer) MobileEngagement(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/azure/mobile/mobile-engagement.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("mobile-engagement", opts...)
}

func init() {
  const namespace = "azure.mobile"
  Register(namespace, "NotificationHubs", azureMobile.NotificationHubs)
  Register(namespace, "AppServiceMobile", azureMobile.AppServiceMobile)
  Register(namespace, "MobileEngagement", azureMobile.MobileEngagement)
}
