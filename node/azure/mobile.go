package azure

import "github.com/blushft/go-diagrams/node"

type mobileContainer struct {
	path string
	opts []node.Option
}

var Mobile = &mobileContainer{
	opts: node.OptionSet{node.Provider("azure"), node.Shape("none")},
	path: "assets/azure/mobile",
}

func (c *mobileContainer) AppServiceMobile(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/azure/mobile/app-service---mobile.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *mobileContainer) MobileEngagement(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/azure/mobile/mobile-engagement.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *mobileContainer) NotificationHubs(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/azure/mobile/notification-hubs.png")}, c.opts, opts)
	return node.New(nopts...)
}
