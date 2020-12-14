package nodes

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type gcpNetworkContainer struct {
	path  string
	attrs []attr.Attribute
}

var GCPNetwork = &gcpNetworkContainer{path: "assets/gcp/network"}

func (c *gcpNetworkContainer) Cdn(opts ...attr.Attribute) *node.Node {
	return node.New("cdn", attr.AssetImage("assets/gcp/network/cdn.png"), attr.Shape(attr.None))
}

func (c *gcpNetworkContainer) DedicatedInterconnect(opts ...attr.Attribute) *node.Node {
	return node.New("dedicated-interconnect", attr.AssetImage("assets/gcp/network/dedicated-interconnect.png"), attr.Shape(attr.None))
}

func (c *gcpNetworkContainer) LoadBalancing(opts ...attr.Attribute) *node.Node {
	return node.New("load-balancing", attr.AssetImage("assets/gcp/network/load-balancing.png"), attr.Shape(attr.None))
}

func (c *gcpNetworkContainer) Router(opts ...attr.Attribute) *node.Node {
	return node.New("router", attr.AssetImage("assets/gcp/network/router.png"), attr.Shape(attr.None))
}

func (c *gcpNetworkContainer) StandardNetworkTier(opts ...attr.Attribute) *node.Node {
	return node.New("standard-network-tier", attr.AssetImage("assets/gcp/network/standard-network-tier.png"), attr.Shape(attr.None))
}

func (c *gcpNetworkContainer) Dns(opts ...attr.Attribute) *node.Node {
	return node.New("dns", attr.AssetImage("assets/gcp/network/dns.png"), attr.Shape(attr.None))
}

func (c *gcpNetworkContainer) Nat(opts ...attr.Attribute) *node.Node {
	return node.New("nat", attr.AssetImage("assets/gcp/network/nat.png"), attr.Shape(attr.None))
}

func (c *gcpNetworkContainer) Network(opts ...attr.Attribute) *node.Node {
	return node.New("network", attr.AssetImage("assets/gcp/network/network.png"), attr.Shape(attr.None))
}

func (c *gcpNetworkContainer) VirtualPrivateCloud(opts ...attr.Attribute) *node.Node {
	return node.New("virtual-private-cloud", attr.AssetImage("assets/gcp/network/virtual-private-cloud.png"), attr.Shape(attr.None))
}

func (c *gcpNetworkContainer) Armor(opts ...attr.Attribute) *node.Node {
	return node.New("armor", attr.AssetImage("assets/gcp/network/armor.png"), attr.Shape(attr.None))
}

func (c *gcpNetworkContainer) FirewallRules(opts ...attr.Attribute) *node.Node {
	return node.New("firewall-rules", attr.AssetImage("assets/gcp/network/firewall-rules.png"), attr.Shape(attr.None))
}

func (c *gcpNetworkContainer) PartnerInterconnect(opts ...attr.Attribute) *node.Node {
	return node.New("partner-interconnect", attr.AssetImage("assets/gcp/network/partner-interconnect.png"), attr.Shape(attr.None))
}

func (c *gcpNetworkContainer) PremiumNetworkTier(opts ...attr.Attribute) *node.Node {
	return node.New("premium-network-tier", attr.AssetImage("assets/gcp/network/premium-network-tier.png"), attr.Shape(attr.None))
}

func (c *gcpNetworkContainer) ExternalIpAddresses(opts ...attr.Attribute) *node.Node {
	return node.New("external-ip-addresses", attr.AssetImage("assets/gcp/network/external-ip-addresses.png"), attr.Shape(attr.None))
}

func (c *gcpNetworkContainer) Routes(opts ...attr.Attribute) *node.Node {
	return node.New("routes", attr.AssetImage("assets/gcp/network/routes.png"), attr.Shape(attr.None))
}

func (c *gcpNetworkContainer) TrafficDirector(opts ...attr.Attribute) *node.Node {
	return node.New("traffic-director", attr.AssetImage("assets/gcp/network/traffic-director.png"), attr.Shape(attr.None))
}

func (c *gcpNetworkContainer) Vpn(opts ...attr.Attribute) *node.Node {
	return node.New("vpn", attr.AssetImage("assets/gcp/network/vpn.png"), attr.Shape(attr.None))
}

func init() {
	const namespace = "gcp.network"
	Register(namespace, "Cdn", GCPNetwork.Cdn)
	Register(namespace, "DedicatedInterconnect", GCPNetwork.DedicatedInterconnect)
	Register(namespace, "LoadBalancing", GCPNetwork.LoadBalancing)
	Register(namespace, "Router", GCPNetwork.Router)
	Register(namespace, "StandardNetworkTier", GCPNetwork.StandardNetworkTier)
	Register(namespace, "Dns", GCPNetwork.Dns)
	Register(namespace, "Nat", GCPNetwork.Nat)
	Register(namespace, "Network", GCPNetwork.Network)
	Register(namespace, "VirtualPrivateCloud", GCPNetwork.VirtualPrivateCloud)
	Register(namespace, "Armor", GCPNetwork.Armor)
	Register(namespace, "FirewallRules", GCPNetwork.FirewallRules)
	Register(namespace, "PartnerInterconnect", GCPNetwork.PartnerInterconnect)
	Register(namespace, "PremiumNetworkTier", GCPNetwork.PremiumNetworkTier)
	Register(namespace, "ExternalIpAddresses", GCPNetwork.ExternalIpAddresses)
	Register(namespace, "Routes", GCPNetwork.Routes)
	Register(namespace, "TrafficDirector", GCPNetwork.TrafficDirector)
	Register(namespace, "Vpn", GCPNetwork.Vpn)
}