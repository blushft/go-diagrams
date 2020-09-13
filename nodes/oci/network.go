package oci

import "github.com/blushft/go-diagrams/diagram"

type networkContainer struct {
	path string
	opts []diagram.NodeOption
}

var Network = &networkContainer{
	opts: diagram.OptionSet{diagram.Provider("oci"), diagram.NodeShape("none")},
	path: "assets/oci/network",
}

func (c *networkContainer) InternetGateway(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/oci/network/internet-gateway.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *networkContainer) Vcn(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/oci/network/vcn.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *networkContainer) FirewallWhite(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/oci/network/firewall-white.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *networkContainer) Firewall(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/oci/network/firewall.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *networkContainer) InternetGatewayWhite(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/oci/network/internet-gateway-white.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *networkContainer) LoadBalancer(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/oci/network/load-balancer.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *networkContainer) SecurityLists(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/oci/network/security-lists.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *networkContainer) ServiceGatewayWhite(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/oci/network/service-gateway-white.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *networkContainer) VcnWhite(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/oci/network/vcn-white.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *networkContainer) DrgWhite(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/oci/network/drg-white.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *networkContainer) Drg(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/oci/network/drg.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *networkContainer) LoadBalancerWhite(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/oci/network/load-balancer-white.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *networkContainer) ServiceGateway(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/oci/network/service-gateway.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *networkContainer) RouteTableWhite(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/oci/network/route-table-white.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *networkContainer) RouteTable(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/oci/network/route-table.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *networkContainer) SecurityListsWhite(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/oci/network/security-lists-white.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
