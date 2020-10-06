package aws

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type networkContainer struct {
	path  string
	attrs []attr.Attribute
}

var Network = &networkContainer{path: "assets/aws/network"}

func (c *networkContainer) ApiGateway(opts ...attr.Attribute) *node.Node {
	return node.New("api-gateway", attr.AssetImage("assets/aws/network/api-gateway.png"), attr.Shape(attr.None))
}

func (c *networkContainer) Nacl(opts ...attr.Attribute) *node.Node {
	return node.New("nacl", attr.AssetImage("assets/aws/network/nacl.png"), attr.Shape(attr.None))
}

func (c *networkContainer) PrivateSubnet(opts ...attr.Attribute) *node.Node {
	return node.New("private-subnet", attr.AssetImage("assets/aws/network/private-subnet.png"), attr.Shape(attr.None))
}

func (c *networkContainer) VpcPeering(opts ...attr.Attribute) *node.Node {
	return node.New("vpc-peering", attr.AssetImage("assets/aws/network/vpc-peering.png"), attr.Shape(attr.None))
}

func (c *networkContainer) Cloudfront(opts ...attr.Attribute) *node.Node {
	return node.New("cloudfront", attr.AssetImage("assets/aws/network/cloudfront.png"), attr.Shape(attr.None))
}

func (c *networkContainer) VpcRouter(opts ...attr.Attribute) *node.Node {
	return node.New("vpc-router", attr.AssetImage("assets/aws/network/vpc-router.png"), attr.Shape(attr.None))
}

func (c *networkContainer) Vpc(opts ...attr.Attribute) *node.Node {
	return node.New("vpc", attr.AssetImage("assets/aws/network/vpc.png"), attr.Shape(attr.None))
}

func (c *networkContainer) NatGateway(opts ...attr.Attribute) *node.Node {
	return node.New("nat-gateway", attr.AssetImage("assets/aws/network/nat-gateway.png"), attr.Shape(attr.None))
}

func (c *networkContainer) PublicSubnet(opts ...attr.Attribute) *node.Node {
	return node.New("public-subnet", attr.AssetImage("assets/aws/network/public-subnet.png"), attr.Shape(attr.None))
}

func (c *networkContainer) Route53(opts ...attr.Attribute) *node.Node {
	return node.New("route-53", attr.AssetImage("assets/aws/network/route-53.png"), attr.Shape(attr.None))
}

func (c *networkContainer) RouteTable(opts ...attr.Attribute) *node.Node {
	return node.New("route-table", attr.AssetImage("assets/aws/network/route-table.png"), attr.Shape(attr.None))
}

func (c *networkContainer) AppMesh(opts ...attr.Attribute) *node.Node {
	return node.New("app-mesh", attr.AssetImage("assets/aws/network/app-mesh.png"), attr.Shape(attr.None))
}

func (c *networkContainer) CloudMap(opts ...attr.Attribute) *node.Node {
	return node.New("cloud-map", attr.AssetImage("assets/aws/network/cloud-map.png"), attr.Shape(attr.None))
}

func (c *networkContainer) Endpoint(opts ...attr.Attribute) *node.Node {
	return node.New("endpoint", attr.AssetImage("assets/aws/network/endpoint.png"), attr.Shape(attr.None))
}

func (c *networkContainer) InternetGateway(opts ...attr.Attribute) *node.Node {
	return node.New("internet-gateway", attr.AssetImage("assets/aws/network/internet-gateway.png"), attr.Shape(attr.None))
}

func (c *networkContainer) SiteToSiteVpn(opts ...attr.Attribute) *node.Node {
	return node.New("site-to-site-vpn", attr.AssetImage("assets/aws/network/site-to-site-vpn.png"), attr.Shape(attr.None))
}

func (c *networkContainer) NetworkingAndContentDelivery(opts ...attr.Attribute) *node.Node {
	return node.New("networking-and-content-delivery", attr.AssetImage("assets/aws/network/networking-and-content-delivery.png"), attr.Shape(attr.None))
}

func (c *networkContainer) Privatelink(opts ...attr.Attribute) *node.Node {
	return node.New("privatelink", attr.AssetImage("assets/aws/network/privatelink.png"), attr.Shape(attr.None))
}

func (c *networkContainer) TransitGateway(opts ...attr.Attribute) *node.Node {
	return node.New("transit-gateway", attr.AssetImage("assets/aws/network/transit-gateway.png"), attr.Shape(attr.None))
}

func (c *networkContainer) ClientVpn(opts ...attr.Attribute) *node.Node {
	return node.New("client-vpn", attr.AssetImage("assets/aws/network/client-vpn.png"), attr.Shape(attr.None))
}

func (c *networkContainer) DirectConnect(opts ...attr.Attribute) *node.Node {
	return node.New("direct-connect", attr.AssetImage("assets/aws/network/direct-connect.png"), attr.Shape(attr.None))
}

func (c *networkContainer) ElasticLoadBalancing(opts ...attr.Attribute) *node.Node {
	return node.New("elastic-load-balancing", attr.AssetImage("assets/aws/network/elastic-load-balancing.png"), attr.Shape(attr.None))
}

func (c *networkContainer) GlobalAccelerator(opts ...attr.Attribute) *node.Node {
	return node.New("global-accelerator", attr.AssetImage("assets/aws/network/global-accelerator.png"), attr.Shape(attr.None))
}
