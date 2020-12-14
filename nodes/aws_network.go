package nodes

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type awsNetworkContainer struct {
	path  string
	attrs []attr.Attribute
}

var AWSNetwork = &awsNetworkContainer{path: "assets/aws/network"}

func (c *awsNetworkContainer) ApiGateway(opts ...attr.Attribute) *node.Node {
	return node.New("api-gateway", attr.AssetImage("assets/aws/network/api-gateway.png"), attr.Shape(attr.None))
}

func (c *awsNetworkContainer) Nacl(opts ...attr.Attribute) *node.Node {
	return node.New("nacl", attr.AssetImage("assets/aws/network/nacl.png"), attr.Shape(attr.None))
}

func (c *awsNetworkContainer) PrivateSubnet(opts ...attr.Attribute) *node.Node {
	return node.New("private-subnet", attr.AssetImage("assets/aws/network/private-subnet.png"), attr.Shape(attr.None))
}

func (c *awsNetworkContainer) VpcPeering(opts ...attr.Attribute) *node.Node {
	return node.New("vpc-peering", attr.AssetImage("assets/aws/network/vpc-peering.png"), attr.Shape(attr.None))
}

func (c *awsNetworkContainer) Cloudfront(opts ...attr.Attribute) *node.Node {
	return node.New("cloudfront", attr.AssetImage("assets/aws/network/cloudfront.png"), attr.Shape(attr.None))
}

func (c *awsNetworkContainer) VpcRouter(opts ...attr.Attribute) *node.Node {
	return node.New("vpc-router", attr.AssetImage("assets/aws/network/vpc-router.png"), attr.Shape(attr.None))
}

func (c *awsNetworkContainer) Vpc(opts ...attr.Attribute) *node.Node {
	return node.New("vpc", attr.AssetImage("assets/aws/network/vpc.png"), attr.Shape(attr.None))
}

func (c *awsNetworkContainer) NatGateway(opts ...attr.Attribute) *node.Node {
	return node.New("nat-gateway", attr.AssetImage("assets/aws/network/nat-gateway.png"), attr.Shape(attr.None))
}

func (c *awsNetworkContainer) PublicSubnet(opts ...attr.Attribute) *node.Node {
	return node.New("public-subnet", attr.AssetImage("assets/aws/network/public-subnet.png"), attr.Shape(attr.None))
}

func (c *awsNetworkContainer) Route53(opts ...attr.Attribute) *node.Node {
	return node.New("route-53", attr.AssetImage("assets/aws/network/route-53.png"), attr.Shape(attr.None))
}

func (c *awsNetworkContainer) RouteTable(opts ...attr.Attribute) *node.Node {
	return node.New("route-table", attr.AssetImage("assets/aws/network/route-table.png"), attr.Shape(attr.None))
}

func (c *awsNetworkContainer) AppMesh(opts ...attr.Attribute) *node.Node {
	return node.New("app-mesh", attr.AssetImage("assets/aws/network/app-mesh.png"), attr.Shape(attr.None))
}

func (c *awsNetworkContainer) CloudMap(opts ...attr.Attribute) *node.Node {
	return node.New("cloud-map", attr.AssetImage("assets/aws/network/cloud-map.png"), attr.Shape(attr.None))
}

func (c *awsNetworkContainer) Endpoint(opts ...attr.Attribute) *node.Node {
	return node.New("endpoint", attr.AssetImage("assets/aws/network/endpoint.png"), attr.Shape(attr.None))
}

func (c *awsNetworkContainer) InternetGateway(opts ...attr.Attribute) *node.Node {
	return node.New("internet-gateway", attr.AssetImage("assets/aws/network/internet-gateway.png"), attr.Shape(attr.None))
}

func (c *awsNetworkContainer) SiteToSiteVpn(opts ...attr.Attribute) *node.Node {
	return node.New("site-to-site-vpn", attr.AssetImage("assets/aws/network/site-to-site-vpn.png"), attr.Shape(attr.None))
}

func (c *awsNetworkContainer) NetworkingAndContentDelivery(opts ...attr.Attribute) *node.Node {
	return node.New("networking-and-content-delivery", attr.AssetImage("assets/aws/network/networking-and-content-delivery.png"), attr.Shape(attr.None))
}

func (c *awsNetworkContainer) Privatelink(opts ...attr.Attribute) *node.Node {
	return node.New("privatelink", attr.AssetImage("assets/aws/network/privatelink.png"), attr.Shape(attr.None))
}

func (c *awsNetworkContainer) TransitGateway(opts ...attr.Attribute) *node.Node {
	return node.New("transit-gateway", attr.AssetImage("assets/aws/network/transit-gateway.png"), attr.Shape(attr.None))
}

func (c *awsNetworkContainer) ClientVpn(opts ...attr.Attribute) *node.Node {
	return node.New("client-vpn", attr.AssetImage("assets/aws/network/client-vpn.png"), attr.Shape(attr.None))
}

func (c *awsNetworkContainer) DirectConnect(opts ...attr.Attribute) *node.Node {
	return node.New("direct-connect", attr.AssetImage("assets/aws/network/direct-connect.png"), attr.Shape(attr.None))
}

func (c *awsNetworkContainer) ElasticLoadBalancing(opts ...attr.Attribute) *node.Node {
	return node.New("elastic-load-balancing", attr.AssetImage("assets/aws/network/elastic-load-balancing.png"), attr.Shape(attr.None))
}

func (c *awsNetworkContainer) GlobalAccelerator(opts ...attr.Attribute) *node.Node {
	return node.New("global-accelerator", attr.AssetImage("assets/aws/network/global-accelerator.png"), attr.Shape(attr.None))
}

func init() {
  const namespace= "aws.network"
  Register(namespace, "ApiGateway", AWSNetwork.ApiGateway)
  Register(namespace, "Nacl", AWSNetwork.Nacl)
  Register(namespace, "PrivateSubnet", AWSNetwork.PrivateSubnet)
  Register(namespace, "VpcPeering", AWSNetwork.VpcPeering)
  Register(namespace, "Cloudfront", AWSNetwork.Cloudfront)
  Register(namespace, "VpcRouter", AWSNetwork.VpcRouter)
  Register(namespace, "Vpc", AWSNetwork.Vpc)
  Register(namespace, "NatGateway", AWSNetwork.NatGateway)
  Register(namespace, "PublicSubnet", AWSNetwork.PublicSubnet)
  Register(namespace, "Route53", AWSNetwork.Route53)
  Register(namespace, "RouteTable", AWSNetwork.RouteTable)
  Register(namespace, "AppMesh", AWSNetwork.AppMesh)
  Register(namespace, "CloudMap", AWSNetwork.CloudMap)
  Register(namespace, "Endpoint", AWSNetwork.Endpoint)
  Register(namespace, "InternetGateway", AWSNetwork.InternetGateway)
  Register(namespace, "SiteToSiteVpn", AWSNetwork.SiteToSiteVpn)
  Register(namespace, "NetworkingAndContentDelivery", AWSNetwork.NetworkingAndContentDelivery)
  Register(namespace, "Privatelink", AWSNetwork.Privatelink)
  Register(namespace, "TransitGateway", AWSNetwork.TransitGateway)
  Register(namespace, "ClientVpn", AWSNetwork.ClientVpn)
  Register(namespace, "DirectConnect", AWSNetwork.DirectConnect)
  Register(namespace, "ElasticLoadBalancing", AWSNetwork.ElasticLoadBalancing)
  Register(namespace, "GlobalAccelerator", AWSNetwork.GlobalAccelerator)
}
