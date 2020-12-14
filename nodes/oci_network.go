package nodes

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type ociNetworkContainer struct {
	path  string
	attrs []attr.Attribute
}

var OciNetwork = &ociNetworkContainer{path: "assets/oci/network"}

func (c *ociNetworkContainer) SecurityListsWhite(opts ...attr.Attribute) *node.Node {
	return node.New("security-lists-white", attr.AssetImage("assets/oci/network/security-lists-white.png"), attr.Shape(attr.None))
}

func (c *ociNetworkContainer) Vcn(opts ...attr.Attribute) *node.Node {
	return node.New("vcn", attr.AssetImage("assets/oci/network/vcn.png"), attr.Shape(attr.None))
}

func (c *ociNetworkContainer) LoadBalancerWhite(opts ...attr.Attribute) *node.Node {
	return node.New("load-balancer-white", attr.AssetImage("assets/oci/network/load-balancer-white.png"), attr.Shape(attr.None))
}

func (c *ociNetworkContainer) RouteTableWhite(opts ...attr.Attribute) *node.Node {
	return node.New("route-table-white", attr.AssetImage("assets/oci/network/route-table-white.png"), attr.Shape(attr.None))
}

func (c *ociNetworkContainer) RouteTable(opts ...attr.Attribute) *node.Node {
	return node.New("route-table", attr.AssetImage("assets/oci/network/route-table.png"), attr.Shape(attr.None))
}

func (c *ociNetworkContainer) SecurityLists(opts ...attr.Attribute) *node.Node {
	return node.New("security-lists", attr.AssetImage("assets/oci/network/security-lists.png"), attr.Shape(attr.None))
}

func (c *ociNetworkContainer) FirewallWhite(opts ...attr.Attribute) *node.Node {
	return node.New("firewall-white", attr.AssetImage("assets/oci/network/firewall-white.png"), attr.Shape(attr.None))
}

func (c *ociNetworkContainer) LoadBalancer(opts ...attr.Attribute) *node.Node {
	return node.New("load-balancer", attr.AssetImage("assets/oci/network/load-balancer.png"), attr.Shape(attr.None))
}

func (c *ociNetworkContainer) ServiceGatewayWhite(opts ...attr.Attribute) *node.Node {
	return node.New("service-gateway-white", attr.AssetImage("assets/oci/network/service-gateway-white.png"), attr.Shape(attr.None))
}

func (c *ociNetworkContainer) VcnWhite(opts ...attr.Attribute) *node.Node {
	return node.New("vcn-white", attr.AssetImage("assets/oci/network/vcn-white.png"), attr.Shape(attr.None))
}

func (c *ociNetworkContainer) ServiceGateway(opts ...attr.Attribute) *node.Node {
	return node.New("service-gateway", attr.AssetImage("assets/oci/network/service-gateway.png"), attr.Shape(attr.None))
}

func (c *ociNetworkContainer) DrgWhite(opts ...attr.Attribute) *node.Node {
	return node.New("drg-white", attr.AssetImage("assets/oci/network/drg-white.png"), attr.Shape(attr.None))
}

func (c *ociNetworkContainer) Drg(opts ...attr.Attribute) *node.Node {
	return node.New("drg", attr.AssetImage("assets/oci/network/drg.png"), attr.Shape(attr.None))
}

func (c *ociNetworkContainer) Firewall(opts ...attr.Attribute) *node.Node {
	return node.New("firewall", attr.AssetImage("assets/oci/network/firewall.png"), attr.Shape(attr.None))
}

func (c *ociNetworkContainer) InternetGatewayWhite(opts ...attr.Attribute) *node.Node {
	return node.New("internet-gateway-white", attr.AssetImage("assets/oci/network/internet-gateway-white.png"), attr.Shape(attr.None))
}

func (c *ociNetworkContainer) InternetGateway(opts ...attr.Attribute) *node.Node {
	return node.New("internet-gateway", attr.AssetImage("assets/oci/network/internet-gateway.png"), attr.Shape(attr.None))
}

func init() {
  const namespace = "oci.network"
  Register(namespace, "SecurityListsWhite", OciNetwork.SecurityListsWhite)
  Register(namespace, "Vcn", OciNetwork.Vcn)
  Register(namespace, "LoadBalancerWhite", OciNetwork.LoadBalancerWhite)
  Register(namespace, "RouteTableWhite", OciNetwork.RouteTableWhite)
  Register(namespace, "RouteTable", OciNetwork.RouteTable)
  Register(namespace, "SecurityLists", OciNetwork.SecurityLists)
  Register(namespace, "FirewallWhite", OciNetwork.FirewallWhite)
  Register(namespace, "LoadBalancer", OciNetwork.LoadBalancer)
  Register(namespace, "ServiceGatewayWhite", OciNetwork.ServiceGatewayWhite)
  Register(namespace, "VcnWhite", OciNetwork.VcnWhite)
  Register(namespace, "ServiceGateway", OciNetwork.ServiceGateway)
  Register(namespace, "DrgWhite", OciNetwork.DrgWhite)
  Register(namespace, "Drg", OciNetwork.Drg)
  Register(namespace, "Firewall", OciNetwork.Firewall)
  Register(namespace, "InternetGatewayWhite", OciNetwork.InternetGatewayWhite)
  Register(namespace, "InternetGateway", OciNetwork.InternetGateway)
}
