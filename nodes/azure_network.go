package nodes

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type azurenetworkContainer struct {
	path  string
	attrs []attr.Attribute
}

var azureNetwork = &azurenetworkContainer{path: "assets/azure/network"}

func (c *azurenetworkContainer) DdosProtectionPlans(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/azure/network/ddos-protection-plans.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("ddos-protection-plans", opts...)
}

func (c *azurenetworkContainer) DnsPrivateZones(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/azure/network/dns-private-zones.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("dns-private-zones", opts...)
}

func (c *azurenetworkContainer) ExpressrouteCircuits(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/azure/network/expressroute-circuits.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("expressroute-circuits", opts...)
}

func (c *azurenetworkContainer) LoadBalancers(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/azure/network/load-balancers.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("load-balancers", opts...)
}

func (c *azurenetworkContainer) NetworkInterfaces(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/azure/network/network-interfaces.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("network-interfaces", opts...)
}

func (c *azurenetworkContainer) TrafficManagerProfiles(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/azure/network/traffic-manager-profiles.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("traffic-manager-profiles", opts...)
}

func (c *azurenetworkContainer) VirtualNetworkGateways(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/azure/network/virtual-network-gateways.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("virtual-network-gateways", opts...)
}

func (c *azurenetworkContainer) VirtualNetworks(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/azure/network/virtual-networks.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("virtual-networks", opts...)
}

func (c *azurenetworkContainer) VirtualWans(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/azure/network/virtual-wans.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("virtual-wans", opts...)
}

func (c *azurenetworkContainer) ApplicationSecurityGroups(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/azure/network/application-security-groups.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("application-security-groups", opts...)
}

func (c *azurenetworkContainer) CdnProfiles(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/azure/network/cdn-profiles.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("cdn-profiles", opts...)
}

func (c *azurenetworkContainer) DnsZones(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/azure/network/dns-zones.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("dns-zones", opts...)
}

func (c *azurenetworkContainer) LocalNetworkGateways(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/azure/network/local-network-gateways.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("local-network-gateways", opts...)
}

func (c *azurenetworkContainer) NetworkSecurityGroupsClassic(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/azure/network/network-security-groups-classic.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("network-security-groups-classic", opts...)
}

func (c *azurenetworkContainer) NetworkWatcher(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/azure/network/network-watcher.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("network-watcher", opts...)
}

func (c *azurenetworkContainer) OnPremisesDataGateways(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/azure/network/on-premises-data-gateways.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("on-premises-data-gateways", opts...)
}

func (c *azurenetworkContainer) ServiceEndpointPolicies(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/azure/network/service-endpoint-policies.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("service-endpoint-policies", opts...)
}

func (c *azurenetworkContainer) ApplicationGateway(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/azure/network/application-gateway.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("application-gateway", opts...)
}

func (c *azurenetworkContainer) Connections(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/azure/network/connections.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("connections", opts...)
}

func (c *azurenetworkContainer) Firewall(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/azure/network/firewall.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("firewall", opts...)
}

func (c *azurenetworkContainer) FrontDoors(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/azure/network/front-doors.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("front-doors", opts...)
}

func (c *azurenetworkContainer) PublicIpAddresses(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/azure/network/public-ip-addresses.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("public-ip-addresses", opts...)
}

func (c *azurenetworkContainer) RouteFilters(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/azure/network/route-filters.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("route-filters", opts...)
}

func (c *azurenetworkContainer) ReservedIpAddressesClassic(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/azure/network/reserved-ip-addresses-classic.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("reserved-ip-addresses-classic", opts...)
}

func (c *azurenetworkContainer) RouteTables(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/azure/network/route-tables.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("route-tables", opts...)
}

func (c *azurenetworkContainer) VirtualNetworkClassic(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/azure/network/virtual-network-classic.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("virtual-network-classic", opts...)
}

func init() {
  const namespace = "azure.network"
  Register(namespace, "DdosProtectionPlans", azureNetwork.DdosProtectionPlans)
  Register(namespace, "DnsPrivateZones", azureNetwork.DnsPrivateZones)
  Register(namespace, "ExpressrouteCircuits", azureNetwork.ExpressrouteCircuits)
  Register(namespace, "LoadBalancers", azureNetwork.LoadBalancers)
  Register(namespace, "NetworkInterfaces", azureNetwork.NetworkInterfaces)
  Register(namespace, "TrafficManagerProfiles", azureNetwork.TrafficManagerProfiles)
  Register(namespace, "VirtualNetworkGateways", azureNetwork.VirtualNetworkGateways)
  Register(namespace, "VirtualNetworks", azureNetwork.VirtualNetworks)
  Register(namespace, "VirtualWans", azureNetwork.VirtualWans)
  Register(namespace, "ApplicationSecurityGroups", azureNetwork.ApplicationSecurityGroups)
  Register(namespace, "CdnProfiles", azureNetwork.CdnProfiles)
  Register(namespace, "DnsZones", azureNetwork.DnsZones)
  Register(namespace, "LocalNetworkGateways", azureNetwork.LocalNetworkGateways)
  Register(namespace, "NetworkSecurityGroupsClassic", azureNetwork.NetworkSecurityGroupsClassic)
  Register(namespace, "NetworkWatcher", azureNetwork.NetworkWatcher)
  Register(namespace, "OnPremisesDataGateways", azureNetwork.OnPremisesDataGateways)
  Register(namespace, "ServiceEndpointPolicies", azureNetwork.ServiceEndpointPolicies)
  Register(namespace, "ApplicationGateway", azureNetwork.ApplicationGateway)
  Register(namespace, "Connections", azureNetwork.Connections)
  Register(namespace, "Firewall", azureNetwork.Firewall)
  Register(namespace, "FrontDoors", azureNetwork.FrontDoors)
  Register(namespace, "PublicIpAddresses", azureNetwork.PublicIpAddresses)
  Register(namespace, "RouteFilters", azureNetwork.RouteFilters)
  Register(namespace, "ReservedIpAddressesClassic", azureNetwork.ReservedIpAddressesClassic)
  Register(namespace, "RouteTables", azureNetwork.RouteTables)
  Register(namespace, "VirtualNetworkClassic", azureNetwork.VirtualNetworkClassic)
}
