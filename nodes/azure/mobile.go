package azure

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type mobileContainer struct {
	path  string
	attrs []attr.Attribute
}

var Mobile = &mobileContainer{path: "assets/azure/mobile"}

func (c *mobileContainer) MobileEngagement(opts ...attr.Attribute) *node.Node {
	return node.New("mobile-engagement", attr.AssetImage("assets/azure/mobile/mobile-engagement.png"), attr.Shape(attr.None))
}

func (c *mobileContainer) NotificationHubs(opts ...attr.Attribute) *node.Node {
	return node.New("notification-hubs", attr.AssetImage("assets/azure/mobile/notification-hubs.png"), attr.Shape(attr.None))
}

func (c *mobileContainer) AppServiceMobile(opts ...attr.Attribute) *node.Node {
	return node.New("app-service---mobile", attr.AssetImage("assets/azure/mobile/app-service---mobile.png"), attr.Shape(attr.None))
}
