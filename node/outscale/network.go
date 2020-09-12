package outscale

import "github.com/blushft/go-diagrams/node"

type networkContainer struct {
	path string
	opts []node.Option
}

var Network = &networkContainer{
	opts: node.OptionSet{node.Provider("outscale"), node.Shape("none")},
	path: "assets/outscale/network",
}

func (c *networkContainer) LoadBalancer(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/outscale/network/load-balancer.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *networkContainer) NatService(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/outscale/network/nat-service.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *networkContainer) Net(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/outscale/network/net.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *networkContainer) SiteToSiteVpng(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/outscale/network/site-to-site-vpng.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *networkContainer) ClientVpn(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/outscale/network/client-vpn.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *networkContainer) InternetService(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/outscale/network/internet-service.png")}, c.opts, opts)
	return node.New(nopts...)
}
