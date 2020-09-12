package oci

import "github.com/blushft/go-diagrams/node"

type networkContainer struct {
	path string
	opts []node.Option
}

var Network = &networkContainer{
	opts: node.OptionSet{node.Provider("oci"), node.Shape("none")},
	path: "assets/oci/network",
}

func (c *networkContainer) ServiceGateway(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/oci/network/service-gateway.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *networkContainer) VcnWhite(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/oci/network/vcn-white.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *networkContainer) Vcn(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/oci/network/vcn.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *networkContainer) LoadBalancer(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/oci/network/load-balancer.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *networkContainer) RouteTableWhite(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/oci/network/route-table-white.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *networkContainer) LoadBalancerWhite(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/oci/network/load-balancer-white.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *networkContainer) SecurityLists(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/oci/network/security-lists.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *networkContainer) DrgWhite(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/oci/network/drg-white.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *networkContainer) Drg(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/oci/network/drg.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *networkContainer) InternetGatewayWhite(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/oci/network/internet-gateway-white.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *networkContainer) InternetGateway(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/oci/network/internet-gateway.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *networkContainer) RouteTable(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/oci/network/route-table.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *networkContainer) FirewallWhite(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/oci/network/firewall-white.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *networkContainer) Firewall(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/oci/network/firewall.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *networkContainer) SecurityListsWhite(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/oci/network/security-lists-white.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *networkContainer) ServiceGatewayWhite(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/oci/network/service-gateway-white.png")}, c.opts, opts)
	return node.New(nopts...)
}
