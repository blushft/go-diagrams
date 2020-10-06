package oci

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type connectivityContainer struct {
	path  string
	attrs []attr.Attribute
}

var Connectivity = &connectivityContainer{path: "assets/oci/connectivity"}

func (c *connectivityContainer) DnsWhite(opts ...attr.Attribute) *node.Node {
	return node.New("dns-white", attr.AssetImage("assets/oci/connectivity/dns-white.png"), attr.Shape(attr.None))
}

func (c *connectivityContainer) VpnWhite(opts ...attr.Attribute) *node.Node {
	return node.New("vpn-white", attr.AssetImage("assets/oci/connectivity/vpn-white.png"), attr.Shape(attr.None))
}

func (c *connectivityContainer) BackboneWhite(opts ...attr.Attribute) *node.Node {
	return node.New("backbone-white", attr.AssetImage("assets/oci/connectivity/backbone-white.png"), attr.Shape(attr.None))
}

func (c *connectivityContainer) CustomerDatacenter(opts ...attr.Attribute) *node.Node {
	return node.New("customer-datacenter", attr.AssetImage("assets/oci/connectivity/customer-datacenter.png"), attr.Shape(attr.None))
}

func (c *connectivityContainer) CustomerPremise(opts ...attr.Attribute) *node.Node {
	return node.New("customer-premise", attr.AssetImage("assets/oci/connectivity/customer-premise.png"), attr.Shape(attr.None))
}

func (c *connectivityContainer) FastConnectWhite(opts ...attr.Attribute) *node.Node {
	return node.New("fast-connect-white", attr.AssetImage("assets/oci/connectivity/fast-connect-white.png"), attr.Shape(attr.None))
}

func (c *connectivityContainer) FastConnect(opts ...attr.Attribute) *node.Node {
	return node.New("fast-connect", attr.AssetImage("assets/oci/connectivity/fast-connect.png"), attr.Shape(attr.None))
}

func (c *connectivityContainer) NatGatewayWhite(opts ...attr.Attribute) *node.Node {
	return node.New("nat-gateway-white", attr.AssetImage("assets/oci/connectivity/nat-gateway-white.png"), attr.Shape(attr.None))
}

func (c *connectivityContainer) Backbone(opts ...attr.Attribute) *node.Node {
	return node.New("backbone", attr.AssetImage("assets/oci/connectivity/backbone.png"), attr.Shape(attr.None))
}

func (c *connectivityContainer) CustomerPremiseWhite(opts ...attr.Attribute) *node.Node {
	return node.New("customer-premise-white", attr.AssetImage("assets/oci/connectivity/customer-premise-white.png"), attr.Shape(attr.None))
}

func (c *connectivityContainer) DisconnectedRegionsWhite(opts ...attr.Attribute) *node.Node {
	return node.New("disconnected-regions-white", attr.AssetImage("assets/oci/connectivity/disconnected-regions-white.png"), attr.Shape(attr.None))
}

func (c *connectivityContainer) CdnWhite(opts ...attr.Attribute) *node.Node {
	return node.New("cdn-white", attr.AssetImage("assets/oci/connectivity/cdn-white.png"), attr.Shape(attr.None))
}

func (c *connectivityContainer) Cdn(opts ...attr.Attribute) *node.Node {
	return node.New("cdn", attr.AssetImage("assets/oci/connectivity/cdn.png"), attr.Shape(attr.None))
}

func (c *connectivityContainer) Dns(opts ...attr.Attribute) *node.Node {
	return node.New("dns", attr.AssetImage("assets/oci/connectivity/dns.png"), attr.Shape(attr.None))
}

func (c *connectivityContainer) NatGateway(opts ...attr.Attribute) *node.Node {
	return node.New("nat-gateway", attr.AssetImage("assets/oci/connectivity/nat-gateway.png"), attr.Shape(attr.None))
}

func (c *connectivityContainer) Vpn(opts ...attr.Attribute) *node.Node {
	return node.New("vpn", attr.AssetImage("assets/oci/connectivity/vpn.png"), attr.Shape(attr.None))
}

func (c *connectivityContainer) CustomerDatacntrWhite(opts ...attr.Attribute) *node.Node {
	return node.New("customer-datacntr-white", attr.AssetImage("assets/oci/connectivity/customer-datacntr-white.png"), attr.Shape(attr.None))
}

func (c *connectivityContainer) DisconnectedRegions(opts ...attr.Attribute) *node.Node {
	return node.New("disconnected-regions", attr.AssetImage("assets/oci/connectivity/disconnected-regions.png"), attr.Shape(attr.None))
}
