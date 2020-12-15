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
	opts = append(opts, attr.AssetImage("assets/gcp/network/cdn.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("cdn", opts...)
}

func (c *gcpNetworkContainer) DedicatedInterconnect(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/gcp/network/dedicated-interconnect.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("dedicated-interconnect", opts...)
}

func (c *gcpNetworkContainer) LoadBalancing(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/gcp/network/load-balancing.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("load-balancing", opts...)
}

func (c *gcpNetworkContainer) Router(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/gcp/network/router.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("router", opts...)
}

func (c *gcpNetworkContainer) StandardNetworkTier(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/gcp/network/standard-network-tier.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("standard-network-tier", opts...)
}

func (c *gcpNetworkContainer) Dns(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/gcp/network/dns.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("dns", opts...)
}

func (c *gcpNetworkContainer) Nat(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/gcp/network/nat.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("nat", opts...)
}

func (c *gcpNetworkContainer) Network(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/gcp/network/network.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("network", opts...)
}

func (c *gcpNetworkContainer) VirtualPrivateCloud(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/gcp/network/virtual-private-cloud.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("virtual-private-cloud", opts...)
}

func (c *gcpNetworkContainer) Armor(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/gcp/network/armor.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("armor", opts...)
}

func (c *gcpNetworkContainer) FirewallRules(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/gcp/network/firewall-rules.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("firewall-rules", opts...)
}

func (c *gcpNetworkContainer) PartnerInterconnect(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/gcp/network/partner-interconnect.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("partner-interconnect", opts...)
}

func (c *gcpNetworkContainer) PremiumNetworkTier(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/gcp/network/premium-network-tier.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("premium-network-tier", opts...)
}

func (c *gcpNetworkContainer) ExternalIpAddresses(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/gcp/network/external-ip-addresses.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("external-ip-addresses", opts...)
}

func (c *gcpNetworkContainer) Routes(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/gcp/network/routes.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("routes", opts...)
}

func (c *gcpNetworkContainer) TrafficDirector(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/gcp/network/traffic-director.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("traffic-director", opts...)
}

func (c *gcpNetworkContainer) Vpn(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/gcp/network/vpn.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("vpn", opts...)
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