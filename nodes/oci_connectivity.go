package nodes

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type ociConnectivityContainer struct {
	path  string
	attrs []attr.Attribute
}

var OciConnectivity = &ociConnectivityContainer{path: "assets/oci/connectivity"}

func (c *ociConnectivityContainer) DnsWhite(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/oci/connectivity/dns-white.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("dns-white", opts...)
}

func (c *ociConnectivityContainer) VpnWhite(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/oci/connectivity/vpn-white.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("vpn-white", opts...)
}

func (c *ociConnectivityContainer) BackboneWhite(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/oci/connectivity/backbone-white.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("backbone-white", opts...)
}

func (c *ociConnectivityContainer) CustomerDatacenter(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/oci/connectivity/customer-datacenter.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("customer-datacenter", opts...)
}

func (c *ociConnectivityContainer) CustomerPremise(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/oci/connectivity/customer-premise.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("customer-premise", opts...)
}

func (c *ociConnectivityContainer) FastConnectWhite(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/oci/connectivity/fast-connect-white.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("fast-connect-white", opts...)
}

func (c *ociConnectivityContainer) FastConnect(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/oci/connectivity/fast-connect.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("fast-connect", opts...)
}

func (c *ociConnectivityContainer) NatGatewayWhite(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/oci/connectivity/nat-gateway-white.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("nat-gateway-white", opts...)
}

func (c *ociConnectivityContainer) Backbone(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/oci/connectivity/backbone.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("backbone", opts...)
}

func (c *ociConnectivityContainer) CustomerPremiseWhite(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/oci/connectivity/customer-premise-white.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("customer-premise-white", opts...)
}

func (c *ociConnectivityContainer) DisconnectedRegionsWhite(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/oci/connectivity/disconnected-regions-white.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("disconnected-regions-white", opts...)
}

func (c *ociConnectivityContainer) CdnWhite(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/oci/connectivity/cdn-white.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("cdn-white", opts...)
}

func (c *ociConnectivityContainer) Cdn(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/oci/connectivity/cdn.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("cdn", opts...)
}

func (c *ociConnectivityContainer) Dns(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/oci/connectivity/dns.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("dns", opts...)
}

func (c *ociConnectivityContainer) NatGateway(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/oci/connectivity/nat-gateway.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("nat-gateway", opts...)
}

func (c *ociConnectivityContainer) Vpn(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/oci/connectivity/vpn.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("vpn", opts...)
}

func (c *ociConnectivityContainer) CustomerDatacntrWhite(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/oci/connectivity/customer-datacntr-white.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("customer-datacntr-white", opts...)
}

func (c *ociConnectivityContainer) DisconnectedRegions(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/oci/connectivity/disconnected-regions.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("disconnected-regions", opts...)
}

func init() {
  const namespace = "oci.connectivity"
  Register(namespace, "DnsWhite", OciConnectivity.DnsWhite)
  Register(namespace, "VpnWhite", OciConnectivity.VpnWhite)
  Register(namespace, "BackboneWhite", OciConnectivity.BackboneWhite)
  Register(namespace, "CustomerDatacenter", OciConnectivity.CustomerDatacenter)
  Register(namespace, "CustomerPremise", OciConnectivity.CustomerPremise)
  Register(namespace, "FastConnectWhite", OciConnectivity.FastConnectWhite)
  Register(namespace, "FastConnect", OciConnectivity.FastConnect)
  Register(namespace, "NatGatewayWhite", OciConnectivity.NatGatewayWhite)
  Register(namespace, "Backbone", OciConnectivity.Backbone)
  Register(namespace, "CustomerPremiseWhite", OciConnectivity.CustomerPremiseWhite)
  Register(namespace, "DisconnectedRegionsWhite", OciConnectivity.DisconnectedRegionsWhite)
  Register(namespace, "CdnWhite", OciConnectivity.CdnWhite)
  Register(namespace, "Cdn", OciConnectivity.Cdn)
  Register(namespace, "Dns", OciConnectivity.Dns)
  Register(namespace, "NatGateway", OciConnectivity.NatGateway)
  Register(namespace, "Vpn", OciConnectivity.Vpn)
  Register(namespace, "CustomerDatacntrWhite", OciConnectivity.CustomerDatacntrWhite)
  Register(namespace, "DisconnectedRegions", OciConnectivity.DisconnectedRegions)
}
