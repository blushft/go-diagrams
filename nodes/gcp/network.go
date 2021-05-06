package gcp

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type networkContainer struct {
	path  string
	attrs []attr.Attribute
}

var Network = &networkContainer{path: "assets/gcp/network"}

func (c *networkContainer) Cdn(opts ...attr.Attribute) *node.Node {
	return node.New("cdn", attr.AssetImage("assets/gcp/network/cdn.png"), attr.Shape(attr.None))
}

func (c *networkContainer) DedicatedInterconnect(opts ...attr.Attribute) *node.Node {
	return node.New("dedicated-interconnect", attr.AssetImage("assets/gcp/network/dedicated-interconnect.png"), attr.Shape(attr.None))
}

func (c *networkContainer) LoadBalancing(opts ...attr.Attribute) *node.Node {
	return node.New("load-balancing", attr.AssetImage("assets/gcp/network/load-balancing.png"), attr.Shape(attr.None))
}

func (c *networkContainer) Network(opts ...attr.Attribute) *node.Node {
	return node.New("network", attr.AssetImage("assets/gcp/network/network.png"), attr.Shape(attr.None))
}

func (c *networkContainer) StandardNetworkTier(opts ...attr.Attribute) *node.Node {
	return node.New("standard-network-tier", attr.AssetImage("assets/gcp/network/standard-network-tier.png"), attr.Shape(attr.None))
}

func (c *networkContainer) VirtualPrivateCloud(opts ...attr.Attribute) *node.Node {
	return node.New("virtual-private-cloud", attr.AssetImage("assets/gcp/network/virtual-private-cloud.png"), attr.Shape(attr.None))
}

func (c *networkContainer) ExternalIpAddresses(opts ...attr.Attribute) *node.Node {
	return node.New("external-ip-addresses", attr.AssetImage("assets/gcp/network/external-ip-addresses.png"), attr.Shape(attr.None))
}

func (c *networkContainer) Vpn(opts ...attr.Attribute) *node.Node {
	return node.New("vpn", attr.AssetImage("assets/gcp/network/vpn.png"), attr.Shape(attr.None))
}

func (c *networkContainer) Armor(opts ...attr.Attribute) *node.Node {
	return node.New("armor", attr.AssetImage("assets/gcp/network/armor.png"), attr.Shape(attr.None))
}

func (c *networkContainer) Dns(opts ...attr.Attribute) *node.Node {
	return node.New("dns", attr.AssetImage("assets/gcp/network/dns.png"), attr.Shape(attr.None))
}

func (c *networkContainer) FirewallRules(opts ...attr.Attribute) *node.Node {
	return node.New("firewall-rules", attr.AssetImage("assets/gcp/network/firewall-rules.png"), attr.Shape(attr.None))
}

func (c *networkContainer) PartnerInterconnect(opts ...attr.Attribute) *node.Node {
	return node.New("partner-interconnect", attr.AssetImage("assets/gcp/network/partner-interconnect.png"), attr.Shape(attr.None))
}

func (c *networkContainer) PremiumNetworkTier(opts ...attr.Attribute) *node.Node {
	return node.New("premium-network-tier", attr.AssetImage("assets/gcp/network/premium-network-tier.png"), attr.Shape(attr.None))
}

func (c *networkContainer) TrafficDirector(opts ...attr.Attribute) *node.Node {
	return node.New("traffic-director", attr.AssetImage("assets/gcp/network/traffic-director.png"), attr.Shape(attr.None))
}

func (c *networkContainer) Nat(opts ...attr.Attribute) *node.Node {
	return node.New("nat", attr.AssetImage("assets/gcp/network/nat.png"), attr.Shape(attr.None))
}

func (c *networkContainer) Router(opts ...attr.Attribute) *node.Node {
	return node.New("router", attr.AssetImage("assets/gcp/network/router.png"), attr.Shape(attr.None))
}

func (c *networkContainer) Routes(opts ...attr.Attribute) *node.Node {
	return node.New("routes", attr.AssetImage("assets/gcp/network/routes.png"), attr.Shape(attr.None))
}
