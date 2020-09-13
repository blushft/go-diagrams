package gcp

import "github.com/blushft/go-diagrams/diagram"

type networkContainer struct {
	path string
	opts []diagram.NodeOption
}

var Network = &networkContainer{
	opts: diagram.OptionSet{diagram.Provider("gcp"), diagram.NodeShape("none")},
	path: "assets/gcp/network",
}

func (c *networkContainer) LoadBalancing(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/gcp/network/load-balancing.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *networkContainer) TrafficDirector(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/gcp/network/traffic-director.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *networkContainer) Vpn(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/gcp/network/vpn.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *networkContainer) Cdn(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/gcp/network/cdn.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *networkContainer) DedicatedInterconnect(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/gcp/network/dedicated-interconnect.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *networkContainer) Dns(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/gcp/network/dns.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *networkContainer) Nat(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/gcp/network/nat.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *networkContainer) Network(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/gcp/network/network.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *networkContainer) PartnerInterconnect(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/gcp/network/partner-interconnect.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *networkContainer) Router(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/gcp/network/router.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *networkContainer) StandardNetworkTier(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/gcp/network/standard-network-tier.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *networkContainer) Armor(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/gcp/network/armor.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *networkContainer) ExternalIpAddresses(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/gcp/network/external-ip-addresses.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *networkContainer) FirewallRules(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/gcp/network/firewall-rules.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *networkContainer) PremiumNetworkTier(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/gcp/network/premium-network-tier.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *networkContainer) Routes(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/gcp/network/routes.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *networkContainer) VirtualPrivateCloud(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/gcp/network/virtual-private-cloud.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
