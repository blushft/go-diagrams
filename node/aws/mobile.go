package aws

import "github.com/blushft/go-diagrams/node"

type mobileContainer struct {
	path string
	opts []node.Option
}

var Mobile = &mobileContainer{
	opts: node.OptionSet{node.Provider("aws"), node.Shape("none")},
	path: "assets/aws/mobile",
}

func (c *mobileContainer) ApiGateway(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/mobile/api-gateway.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *mobileContainer) Appsync(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/mobile/appsync.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *mobileContainer) DeviceFarm(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/mobile/device-farm.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *mobileContainer) Pinpoint(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/mobile/pinpoint.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *mobileContainer) Amplify(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/mobile/amplify.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *mobileContainer) ApiGatewayEndpoint(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/mobile/api-gateway-endpoint.png")}, c.opts, opts)
	return node.New(nopts...)
}
