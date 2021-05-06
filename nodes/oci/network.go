package oci

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type networkContainer struct {
	path  string
	attrs []attr.Attribute
}

var Network = &networkContainer{path: "assets/oci/network"}

func (c *networkContainer) LoadBalancerWhite(opts ...attr.Attribute) *node.Node {
	return node.New("load-balancer-white", attr.AssetImage("assets/oci/network/load-balancer-white.png"), attr.Shape(attr.None))
}

func (c *networkContainer) ServiceGatewayWhite(opts ...attr.Attribute) *node.Node {
	return node.New("service-gateway-white", attr.AssetImage("assets/oci/network/service-gateway-white.png"), attr.Shape(attr.None))
}

func (c *networkContainer) ServiceGateway(opts ...attr.Attribute) *node.Node {
	return node.New("service-gateway", attr.AssetImage("assets/oci/network/service-gateway.png"), attr.Shape(attr.None))
}

func (c *networkContainer) SecurityLists(opts ...attr.Attribute) *node.Node {
	return node.New("security-lists", attr.AssetImage("assets/oci/network/security-lists.png"), attr.Shape(attr.None))
}

func (c *networkContainer) DrgWhite(opts ...attr.Attribute) *node.Node {
	return node.New("drg-white", attr.AssetImage("assets/oci/network/drg-white.png"), attr.Shape(attr.None))
}

func (c *networkContainer) FirewallWhite(opts ...attr.Attribute) *node.Node {
	return node.New("firewall-white", attr.AssetImage("assets/oci/network/firewall-white.png"), attr.Shape(attr.None))
}

func (c *networkContainer) Firewall(opts ...attr.Attribute) *node.Node {
	return node.New("firewall", attr.AssetImage("assets/oci/network/firewall.png"), attr.Shape(attr.None))
}

func (c *networkContainer) InternetGatewayWhite(opts ...attr.Attribute) *node.Node {
	return node.New("internet-gateway-white", attr.AssetImage("assets/oci/network/internet-gateway-white.png"), attr.Shape(attr.None))
}

func (c *networkContainer) LoadBalancer(opts ...attr.Attribute) *node.Node {
	return node.New("load-balancer", attr.AssetImage("assets/oci/network/load-balancer.png"), attr.Shape(attr.None))
}

func (c *networkContainer) RouteTableWhite(opts ...attr.Attribute) *node.Node {
	return node.New("route-table-white", attr.AssetImage("assets/oci/network/route-table-white.png"), attr.Shape(attr.None))
}

func (c *networkContainer) SecurityListsWhite(opts ...attr.Attribute) *node.Node {
	return node.New("security-lists-white", attr.AssetImage("assets/oci/network/security-lists-white.png"), attr.Shape(attr.None))
}

func (c *networkContainer) Drg(opts ...attr.Attribute) *node.Node {
	return node.New("drg", attr.AssetImage("assets/oci/network/drg.png"), attr.Shape(attr.None))
}

func (c *networkContainer) Vcn(opts ...attr.Attribute) *node.Node {
	return node.New("vcn", attr.AssetImage("assets/oci/network/vcn.png"), attr.Shape(attr.None))
}

func (c *networkContainer) InternetGateway(opts ...attr.Attribute) *node.Node {
	return node.New("internet-gateway", attr.AssetImage("assets/oci/network/internet-gateway.png"), attr.Shape(attr.None))
}

func (c *networkContainer) RouteTable(opts ...attr.Attribute) *node.Node {
	return node.New("route-table", attr.AssetImage("assets/oci/network/route-table.png"), attr.Shape(attr.None))
}

func (c *networkContainer) VcnWhite(opts ...attr.Attribute) *node.Node {
	return node.New("vcn-white", attr.AssetImage("assets/oci/network/vcn-white.png"), attr.Shape(attr.None))
}
