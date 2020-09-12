package gcp

import "github.com/blushft/go-diagrams/node"

type networkContainer struct {
	path string
	opts []node.Option
}

var Network = &networkContainer{
	opts: node.OptionSet{node.Provider("gcp"), node.Shape("none")},
	path: "assets/gcp/network",
}

func (c *networkContainer) Dns(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/gcp/network/dns.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *networkContainer) FirewallRules(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/gcp/network/firewall-rules.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *networkContainer) PartnerInterconnect(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/gcp/network/partner-interconnect.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *networkContainer) Network(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/gcp/network/network.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *networkContainer) VirtualPrivateCloud(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/gcp/network/virtual-private-cloud.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *networkContainer) Vpn(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/gcp/network/vpn.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *networkContainer) Armor(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/gcp/network/armor.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *networkContainer) DedicatedInterconnect(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/gcp/network/dedicated-interconnect.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *networkContainer) ExternalIpAddresses(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/gcp/network/external-ip-addresses.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *networkContainer) PremiumNetworkTier(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/gcp/network/premium-network-tier.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *networkContainer) Router(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/gcp/network/router.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *networkContainer) Routes(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/gcp/network/routes.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *networkContainer) StandardNetworkTier(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/gcp/network/standard-network-tier.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *networkContainer) TrafficDirector(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/gcp/network/traffic-director.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *networkContainer) Cdn(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/gcp/network/cdn.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *networkContainer) LoadBalancing(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/gcp/network/load-balancing.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *networkContainer) Nat(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/gcp/network/nat.png")}, c.opts, opts)
	return node.New(nopts...)
}
