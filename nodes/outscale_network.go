package nodes

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type outscaleNetworkContainer struct {
	path  string
	attrs []attr.Attribute
}

var OutscaleNetwork =&outscaleNetworkContainer{path: "assets/outscale/network"}

func (c *outscaleNetworkContainer) LoadBalancer(opts ...attr.Attribute) *node.Node {
	return node.New("load-balancer", attr.AssetImage("assets/outscale/network/load-balancer.png"), attr.Shape(attr.None))
}

func (c *outscaleNetworkContainer) NatService(opts ...attr.Attribute) *node.Node {
	return node.New("nat-service", attr.AssetImage("assets/outscale/network/nat-service.png"), attr.Shape(attr.None))
}

func (c *outscaleNetworkContainer) Net(opts ...attr.Attribute) *node.Node {
	return node.New("net", attr.AssetImage("assets/outscale/network/net.png"), attr.Shape(attr.None))
}

func (c *outscaleNetworkContainer) SiteToSiteVpng(opts ...attr.Attribute) *node.Node {
	return node.New("site-to-site-vpng", attr.AssetImage("assets/outscale/network/site-to-site-vpng.png"), attr.Shape(attr.None))
}

func (c *outscaleNetworkContainer) ClientVpn(opts ...attr.Attribute) *node.Node {
	return node.New("client-vpn", attr.AssetImage("assets/outscale/network/client-vpn.png"), attr.Shape(attr.None))
}

func (c *outscaleNetworkContainer) InternetService(opts ...attr.Attribute) *node.Node {
	return node.New("internet-service", attr.AssetImage("assets/outscale/network/internet-service.png"), attr.Shape(attr.None))
}

func init() {
  const namespace = "outscale.network"
  Register(namespace, "LoadBalancer", OutscaleNetwork.LoadBalancer)
  Register(namespace, "NatService", OutscaleNetwork.NatService)
  Register(namespace, "Net", OutscaleNetwork.Net)
  Register(namespace, "SiteToSiteVpng", OutscaleNetwork.SiteToSiteVpng)
  Register(namespace, "ClientVpn", OutscaleNetwork.ClientVpn)
  Register(namespace, "InternetService", OutscaleNetwork.InternetService)
}
