package outscale

import "github.com/blushft/go-diagrams/diagram"

type networkContainer struct {
	path string
	opts []diagram.NodeOption
}

var Network = &networkContainer{
	opts: diagram.OptionSet{diagram.Provider("outscale"), diagram.NodeShape("none")},
	path: "assets/outscale/network",
}

func (c *networkContainer) InternetService(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/outscale/network/internet-service.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *networkContainer) LoadBalancer(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/outscale/network/load-balancer.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *networkContainer) NatService(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/outscale/network/nat-service.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *networkContainer) Net(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/outscale/network/net.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *networkContainer) SiteToSiteVpng(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/outscale/network/site-to-site-vpng.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *networkContainer) ClientVpn(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/outscale/network/client-vpn.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
