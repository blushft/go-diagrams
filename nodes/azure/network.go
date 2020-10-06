package azure

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type networkContainer struct {
	path  string
	attrs []attr.Attribute
}

var Network = &networkContainer{path: "assets/azure/network"}

func (c *networkContainer) DdosProtectionPlans(opts ...attr.Attribute) *node.Node {
	return node.New("ddos-protection-plans", attr.AssetImage("assets/azure/network/ddos-protection-plans.png"), attr.Shape(attr.None))
}

func (c *networkContainer) DnsPrivateZones(opts ...attr.Attribute) *node.Node {
	return node.New("dns-private-zones", attr.AssetImage("assets/azure/network/dns-private-zones.png"), attr.Shape(attr.None))
}

func (c *networkContainer) ExpressrouteCircuits(opts ...attr.Attribute) *node.Node {
	return node.New("expressroute-circuits", attr.AssetImage("assets/azure/network/expressroute-circuits.png"), attr.Shape(attr.None))
}

func (c *networkContainer) LoadBalancers(opts ...attr.Attribute) *node.Node {
	return node.New("load-balancers", attr.AssetImage("assets/azure/network/load-balancers.png"), attr.Shape(attr.None))
}

func (c *networkContainer) NetworkInterfaces(opts ...attr.Attribute) *node.Node {
	return node.New("network-interfaces", attr.AssetImage("assets/azure/network/network-interfaces.png"), attr.Shape(attr.None))
}

func (c *networkContainer) TrafficManagerProfiles(opts ...attr.Attribute) *node.Node {
	return node.New("traffic-manager-profiles", attr.AssetImage("assets/azure/network/traffic-manager-profiles.png"), attr.Shape(attr.None))
}

func (c *networkContainer) VirtualNetworkGateways(opts ...attr.Attribute) *node.Node {
	return node.New("virtual-network-gateways", attr.AssetImage("assets/azure/network/virtual-network-gateways.png"), attr.Shape(attr.None))
}

func (c *networkContainer) VirtualNetworks(opts ...attr.Attribute) *node.Node {
	return node.New("virtual-networks", attr.AssetImage("assets/azure/network/virtual-networks.png"), attr.Shape(attr.None))
}

func (c *networkContainer) VirtualWans(opts ...attr.Attribute) *node.Node {
	return node.New("virtual-wans", attr.AssetImage("assets/azure/network/virtual-wans.png"), attr.Shape(attr.None))
}

func (c *networkContainer) ApplicationSecurityGroups(opts ...attr.Attribute) *node.Node {
	return node.New("application-security-groups", attr.AssetImage("assets/azure/network/application-security-groups.png"), attr.Shape(attr.None))
}

func (c *networkContainer) CdnProfiles(opts ...attr.Attribute) *node.Node {
	return node.New("cdn-profiles", attr.AssetImage("assets/azure/network/cdn-profiles.png"), attr.Shape(attr.None))
}

func (c *networkContainer) DnsZones(opts ...attr.Attribute) *node.Node {
	return node.New("dns-zones", attr.AssetImage("assets/azure/network/dns-zones.png"), attr.Shape(attr.None))
}

func (c *networkContainer) LocalNetworkGateways(opts ...attr.Attribute) *node.Node {
	return node.New("local-network-gateways", attr.AssetImage("assets/azure/network/local-network-gateways.png"), attr.Shape(attr.None))
}

func (c *networkContainer) NetworkSecurityGroupsClassic(opts ...attr.Attribute) *node.Node {
	return node.New("network-security-groups-classic", attr.AssetImage("assets/azure/network/network-security-groups-classic.png"), attr.Shape(attr.None))
}

func (c *networkContainer) NetworkWatcher(opts ...attr.Attribute) *node.Node {
	return node.New("network-watcher", attr.AssetImage("assets/azure/network/network-watcher.png"), attr.Shape(attr.None))
}

func (c *networkContainer) OnPremisesDataGateways(opts ...attr.Attribute) *node.Node {
	return node.New("on-premises-data-gateways", attr.AssetImage("assets/azure/network/on-premises-data-gateways.png"), attr.Shape(attr.None))
}

func (c *networkContainer) ServiceEndpointPolicies(opts ...attr.Attribute) *node.Node {
	return node.New("service-endpoint-policies", attr.AssetImage("assets/azure/network/service-endpoint-policies.png"), attr.Shape(attr.None))
}

func (c *networkContainer) ApplicationGateway(opts ...attr.Attribute) *node.Node {
	return node.New("application-gateway", attr.AssetImage("assets/azure/network/application-gateway.png"), attr.Shape(attr.None))
}

func (c *networkContainer) Connections(opts ...attr.Attribute) *node.Node {
	return node.New("connections", attr.AssetImage("assets/azure/network/connections.png"), attr.Shape(attr.None))
}

func (c *networkContainer) Firewall(opts ...attr.Attribute) *node.Node {
	return node.New("firewall", attr.AssetImage("assets/azure/network/firewall.png"), attr.Shape(attr.None))
}

func (c *networkContainer) FrontDoors(opts ...attr.Attribute) *node.Node {
	return node.New("front-doors", attr.AssetImage("assets/azure/network/front-doors.png"), attr.Shape(attr.None))
}

func (c *networkContainer) PublicIpAddresses(opts ...attr.Attribute) *node.Node {
	return node.New("public-ip-addresses", attr.AssetImage("assets/azure/network/public-ip-addresses.png"), attr.Shape(attr.None))
}

func (c *networkContainer) RouteFilters(opts ...attr.Attribute) *node.Node {
	return node.New("route-filters", attr.AssetImage("assets/azure/network/route-filters.png"), attr.Shape(attr.None))
}

func (c *networkContainer) ReservedIpAddressesClassic(opts ...attr.Attribute) *node.Node {
	return node.New("reserved-ip-addresses-classic", attr.AssetImage("assets/azure/network/reserved-ip-addresses-classic.png"), attr.Shape(attr.None))
}

func (c *networkContainer) RouteTables(opts ...attr.Attribute) *node.Node {
	return node.New("route-tables", attr.AssetImage("assets/azure/network/route-tables.png"), attr.Shape(attr.None))
}

func (c *networkContainer) VirtualNetworkClassic(opts ...attr.Attribute) *node.Node {
	return node.New("virtual-network-classic", attr.AssetImage("assets/azure/network/virtual-network-classic.png"), attr.Shape(attr.None))
}
