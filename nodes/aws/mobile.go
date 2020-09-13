package aws

import "github.com/blushft/go-diagrams/diagram"

type mobileContainer struct {
	path string
	opts []diagram.NodeOption
}

var Mobile = &mobileContainer{
	opts: diagram.OptionSet{diagram.Provider("aws"), diagram.NodeShape("none")},
	path: "assets/aws/mobile",
}

func (c *mobileContainer) Amplify(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/mobile/amplify.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *mobileContainer) ApiGatewayEndpoint(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/mobile/api-gateway-endpoint.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *mobileContainer) ApiGateway(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/mobile/api-gateway.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *mobileContainer) Appsync(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/mobile/appsync.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *mobileContainer) DeviceFarm(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/mobile/device-farm.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *mobileContainer) Pinpoint(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/mobile/pinpoint.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
