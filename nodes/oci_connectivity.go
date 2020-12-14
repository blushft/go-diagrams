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
	return node.New("dns-white", attr.AssetImage("assets/oci/connectivity/dns-white.png"), attr.Shape(attr.None))
}

func (c *ociConnectivityContainer) VpnWhite(opts ...attr.Attribute) *node.Node {
	return node.New("vpn-white", attr.AssetImage("assets/oci/connectivity/vpn-white.png"), attr.Shape(attr.None))
}

func (c *ociConnectivityContainer) BackboneWhite(opts ...attr.Attribute) *node.Node {
	return node.New("backbone-white", attr.AssetImage("assets/oci/connectivity/backbone-white.png"), attr.Shape(attr.None))
}

func (c *ociConnectivityContainer) CustomerDatacenter(opts ...attr.Attribute) *node.Node {
	return node.New("customer-datacenter", attr.AssetImage("assets/oci/connectivity/customer-datacenter.png"), attr.Shape(attr.None))
}

func (c *ociConnectivityContainer) CustomerPremise(opts ...attr.Attribute) *node.Node {
	return node.New("customer-premise", attr.AssetImage("assets/oci/connectivity/customer-premise.png"), attr.Shape(attr.None))
}

func (c *ociConnectivityContainer) FastConnectWhite(opts ...attr.Attribute) *node.Node {
	return node.New("fast-connect-white", attr.AssetImage("assets/oci/connectivity/fast-connect-white.png"), attr.Shape(attr.None))
}

func (c *ociConnectivityContainer) FastConnect(opts ...attr.Attribute) *node.Node {
	return node.New("fast-connect", attr.AssetImage("assets/oci/connectivity/fast-connect.png"), attr.Shape(attr.None))
}

func (c *ociConnectivityContainer) NatGatewayWhite(opts ...attr.Attribute) *node.Node {
	return node.New("nat-gateway-white", attr.AssetImage("assets/oci/connectivity/nat-gateway-white.png"), attr.Shape(attr.None))
}

func (c *ociConnectivityContainer) Backbone(opts ...attr.Attribute) *node.Node {
	return node.New("backbone", attr.AssetImage("assets/oci/connectivity/backbone.png"), attr.Shape(attr.None))
}

func (c *ociConnectivityContainer) CustomerPremiseWhite(opts ...attr.Attribute) *node.Node {
	return node.New("customer-premise-white", attr.AssetImage("assets/oci/connectivity/customer-premise-white.png"), attr.Shape(attr.None))
}

func (c *ociConnectivityContainer) DisconnectedRegionsWhite(opts ...attr.Attribute) *node.Node {
	return node.New("disconnected-regions-white", attr.AssetImage("assets/oci/connectivity/disconnected-regions-white.png"), attr.Shape(attr.None))
}

func (c *ociConnectivityContainer) CdnWhite(opts ...attr.Attribute) *node.Node {
	return node.New("cdn-white", attr.AssetImage("assets/oci/connectivity/cdn-white.png"), attr.Shape(attr.None))
}

func (c *ociConnectivityContainer) Cdn(opts ...attr.Attribute) *node.Node {
	return node.New("cdn", attr.AssetImage("assets/oci/connectivity/cdn.png"), attr.Shape(attr.None))
}

func (c *ociConnectivityContainer) Dns(opts ...attr.Attribute) *node.Node {
	return node.New("dns", attr.AssetImage("assets/oci/connectivity/dns.png"), attr.Shape(attr.None))
}

func (c *ociConnectivityContainer) NatGateway(opts ...attr.Attribute) *node.Node {
	return node.New("nat-gateway", attr.AssetImage("assets/oci/connectivity/nat-gateway.png"), attr.Shape(attr.None))
}

func (c *ociConnectivityContainer) Vpn(opts ...attr.Attribute) *node.Node {
	return node.New("vpn", attr.AssetImage("assets/oci/connectivity/vpn.png"), attr.Shape(attr.None))
}

func (c *ociConnectivityContainer) CustomerDatacntrWhite(opts ...attr.Attribute) *node.Node {
	return node.New("customer-datacntr-white", attr.AssetImage("assets/oci/connectivity/customer-datacntr-white.png"), attr.Shape(attr.None))
}

func (c *ociConnectivityContainer) DisconnectedRegions(opts ...attr.Attribute) *node.Node {
	return node.New("disconnected-regions", attr.AssetImage("assets/oci/connectivity/disconnected-regions.png"), attr.Shape(attr.None))
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
