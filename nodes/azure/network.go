package azure

import "github.com/blushft/go-diagrams/diagram"

type networkContainer struct {
	path string
	opts []diagram.NodeOption
}

var Network = &networkContainer{
	opts: diagram.OptionSet{diagram.Provider("azure"), diagram.NodeShape("none")},
	path: "assets/azure/network",
}

func (c *networkContainer) VirtualWans(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/network/virtual-wans.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *networkContainer) ApplicationGateway(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/network/application-gateway.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *networkContainer) FrontDoors(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/network/front-doors.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *networkContainer) LoadBalancers(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/network/load-balancers.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *networkContainer) NetworkInterfaces(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/network/network-interfaces.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *networkContainer) NetworkWatcher(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/network/network-watcher.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *networkContainer) RouteTables(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/network/route-tables.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *networkContainer) VirtualNetworkGateways(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/network/virtual-network-gateways.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *networkContainer) Connections(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/network/connections.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *networkContainer) DnsPrivateZones(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/network/dns-private-zones.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *networkContainer) DnsZones(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/network/dns-zones.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *networkContainer) NetworkSecurityGroupsClassic(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/network/network-security-groups-classic.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *networkContainer) OnPremisesDataGateways(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/network/on-premises-data-gateways.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *networkContainer) VirtualNetworks(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/network/virtual-networks.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *networkContainer) ApplicationSecurityGroups(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/network/application-security-groups.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *networkContainer) CdnProfiles(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/network/cdn-profiles.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *networkContainer) ExpressrouteCircuits(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/network/expressroute-circuits.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *networkContainer) ServiceEndpointPolicies(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/network/service-endpoint-policies.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *networkContainer) VirtualNetworkClassic(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/network/virtual-network-classic.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *networkContainer) RouteFilters(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/network/route-filters.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *networkContainer) TrafficManagerProfiles(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/network/traffic-manager-profiles.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *networkContainer) DdosProtectionPlans(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/network/ddos-protection-plans.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *networkContainer) Firewall(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/network/firewall.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *networkContainer) LocalNetworkGateways(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/network/local-network-gateways.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *networkContainer) PublicIpAddresses(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/network/public-ip-addresses.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *networkContainer) ReservedIpAddressesClassic(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/network/reserved-ip-addresses-classic.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
