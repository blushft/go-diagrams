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
	opts = append(opts, attr.AssetImage("assets/oci/network/security-lists-white.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("security-lists-white", opts...)
}

func (c *ociNetworkContainer) Vcn(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/oci/network/vcn.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("vcn", opts...)
}

func (c *ociNetworkContainer) LoadBalancerWhite(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/oci/network/load-balancer-white.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("load-balancer-white", opts...)
}

func (c *ociNetworkContainer) RouteTableWhite(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/oci/network/route-table-white.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("route-table-white", opts...)
}

func (c *ociNetworkContainer) RouteTable(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/oci/network/route-table.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("route-table", opts...)
}

func (c *ociNetworkContainer) SecurityLists(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/oci/network/security-lists.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("security-lists", opts...)
}

func (c *ociNetworkContainer) FirewallWhite(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/oci/network/firewall-white.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("firewall-white", opts...)
}

func (c *ociNetworkContainer) LoadBalancer(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/oci/network/load-balancer.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("load-balancer", opts...)
}

func (c *ociNetworkContainer) ServiceGatewayWhite(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/oci/network/service-gateway-white.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("service-gateway-white", opts...)
}

func (c *ociNetworkContainer) VcnWhite(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/oci/network/vcn-white.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("vcn-white", opts...)
}

func (c *ociNetworkContainer) ServiceGateway(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/oci/network/service-gateway.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("service-gateway", opts...)
}

func (c *ociNetworkContainer) DrgWhite(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/oci/network/drg-white.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("drg-white", opts...)
}

func (c *ociNetworkContainer) Drg(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/oci/network/drg.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("drg", opts...)
}

func (c *ociNetworkContainer) Firewall(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/oci/network/firewall.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("firewall", opts...)
}

func (c *ociNetworkContainer) InternetGatewayWhite(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/oci/network/internet-gateway-white.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("internet-gateway-white", opts...)
}

func (c *ociNetworkContainer) InternetGateway(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/oci/network/internet-gateway.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("internet-gateway", opts...)
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
