package nodes

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type networkContainer struct {
	path  string
	attrs []attr.Attribute
}

var Network = &networkContainer{path: "assets/outscale/network"}

func (c *networkContainer) LoadBalancer(opts ...attr.Attribute) *node.Node {
	return node.New("load-balancer", attr.AssetImage("assets/outscale/network/load-balancer.png"), attr.Shape(attr.None))
}

func (c *networkContainer) NatService(opts ...attr.Attribute) *node.Node {
	return node.New("nat-service", attr.AssetImage("assets/outscale/network/nat-service.png"), attr.Shape(attr.None))
}

func (c *networkContainer) Net(opts ...attr.Attribute) *node.Node {
	return node.New("net", attr.AssetImage("assets/outscale/network/net.png"), attr.Shape(attr.None))
}

func (c *networkContainer) SiteToSiteVpng(opts ...attr.Attribute) *node.Node {
	return node.New("site-to-site-vpng", attr.AssetImage("assets/outscale/network/site-to-site-vpng.png"), attr.Shape(attr.None))
}

func (c *networkContainer) ClientVpn(opts ...attr.Attribute) *node.Node {
	return node.New("client-vpn", attr.AssetImage("assets/outscale/network/client-vpn.png"), attr.Shape(attr.None))
}

func (c *networkContainer) InternetService(opts ...attr.Attribute) *node.Node {
	return node.New("internet-service", attr.AssetImage("assets/outscale/network/internet-service.png"), attr.Shape(attr.None))
}
