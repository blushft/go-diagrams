package azure

import "github.com/blushft/go-diagrams/node"

type networkContainer struct {
	path string
	opts []node.Option
}

var Network = &networkContainer{
	opts: node.OptionSet{node.Provider("azure"), node.Shape("none")},
	path: "assets/azure/network",
}

func (c *networkContainer) DdosProtectionPlans(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/azure/network/ddos-protection-plans.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *networkContainer) Firewall(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/azure/network/firewall.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *networkContainer) NetworkWatcher(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/azure/network/network-watcher.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *networkContainer) OnPremisesDataGateways(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/azure/network/on-premises-data-gateways.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *networkContainer) RouteFilters(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/azure/network/route-filters.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *networkContainer) VirtualNetworkClassic(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/azure/network/virtual-network-classic.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *networkContainer) VirtualNetworks(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/azure/network/virtual-networks.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *networkContainer) CdnProfiles(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/azure/network/cdn-profiles.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *networkContainer) DnsZones(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/azure/network/dns-zones.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *networkContainer) LoadBalancers(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/azure/network/load-balancers.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *networkContainer) NetworkSecurityGroupsClassic(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/azure/network/network-security-groups-classic.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *networkContainer) ServiceEndpointPolicies(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/azure/network/service-endpoint-policies.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *networkContainer) ApplicationGateway(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/azure/network/application-gateway.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *networkContainer) Connections(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/azure/network/connections.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *networkContainer) DnsPrivateZones(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/azure/network/dns-private-zones.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *networkContainer) ExpressrouteCircuits(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/azure/network/expressroute-circuits.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *networkContainer) LocalNetworkGateways(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/azure/network/local-network-gateways.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *networkContainer) PublicIpAddresses(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/azure/network/public-ip-addresses.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *networkContainer) RouteTables(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/azure/network/route-tables.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *networkContainer) TrafficManagerProfiles(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/azure/network/traffic-manager-profiles.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *networkContainer) VirtualNetworkGateways(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/azure/network/virtual-network-gateways.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *networkContainer) VirtualWans(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/azure/network/virtual-wans.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *networkContainer) ApplicationSecurityGroups(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/azure/network/application-security-groups.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *networkContainer) FrontDoors(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/azure/network/front-doors.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *networkContainer) NetworkInterfaces(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/azure/network/network-interfaces.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *networkContainer) ReservedIpAddressesClassic(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/azure/network/reserved-ip-addresses-classic.png")}, c.opts, opts)
	return node.New(nopts...)
}
