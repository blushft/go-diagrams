package azure

import "github.com/blushft/go-diagrams/diagram"

type mobileContainer struct {
	path string
	opts []diagram.NodeOption
}

var Mobile = &mobileContainer{
	opts: diagram.OptionSet{diagram.Provider("azure"), diagram.NodeShape("none")},
	path: "assets/azure/mobile",
}

func (c *mobileContainer) AppServiceMobile(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/mobile/app-service---mobile.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *mobileContainer) MobileEngagement(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/mobile/mobile-engagement.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *mobileContainer) NotificationHubs(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/mobile/notification-hubs.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
