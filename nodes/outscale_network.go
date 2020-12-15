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
	opts = append(opts, attr.AssetImage("assets/outscale/network/load-balancer.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("load-balancer", opts...)
}

func (c *outscaleNetworkContainer) NatService(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/outscale/network/nat-service.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("nat-service", opts...)
}

func (c *outscaleNetworkContainer) Net(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/outscale/network/net.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("net", opts...)
}

func (c *outscaleNetworkContainer) SiteToSiteVpng(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/outscale/network/site-to-site-vpng.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("site-to-site-vpng", opts...)
}

func (c *outscaleNetworkContainer) ClientVpn(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/outscale/network/client-vpn.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("client-vpn", opts...)
}

func (c *outscaleNetworkContainer) InternetService(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/outscale/network/internet-service.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("internet-service", opts...)
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
