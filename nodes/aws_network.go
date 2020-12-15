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
	opts = append(opts, attr.AssetImage("assets/aws/network/api-gateway.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("api-gateway", opts...)
}

func (c *awsNetworkContainer) Nacl(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/network/nacl.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("nacl", opts...)
}

func (c *awsNetworkContainer) PrivateSubnet(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/network/private-subnet.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("private-subnet", opts...)
}

func (c *awsNetworkContainer) VpcPeering(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/network/vpc-peering.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("vpc-peering", opts...)
}

func (c *awsNetworkContainer) Cloudfront(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/network/cloudfront.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("cloudfront", opts...)
}

func (c *awsNetworkContainer) VpcRouter(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/network/vpc-router.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("vpc-router", opts...)
}

func (c *awsNetworkContainer) Vpc(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/network/vpc.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("vpc", opts...)
}

func (c *awsNetworkContainer) NatGateway(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/network/nat-gateway.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("nat-gateway", opts...)
}

func (c *awsNetworkContainer) PublicSubnet(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/network/public-subnet.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("public-subnet", opts...)
}

func (c *awsNetworkContainer) Route53(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/network/route-53.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("route-53", opts...)
}

func (c *awsNetworkContainer) RouteTable(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/network/route-table.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("route-table", opts...)
}

func (c *awsNetworkContainer) AppMesh(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/network/app-mesh.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("app-mesh", opts...)
}

func (c *awsNetworkContainer) CloudMap(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/network/cloud-map.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("cloud-map", opts...)
}

func (c *awsNetworkContainer) Endpoint(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/network/endpoint.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("endpoint", opts...)
}

func (c *awsNetworkContainer) InternetGateway(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/network/internet-gateway.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("internet-gateway", opts...)
}

func (c *awsNetworkContainer) SiteToSiteVpn(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/network/site-to-site-vpn.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("site-to-site-vpn", opts...)
}

func (c *awsNetworkContainer) NetworkingAndContentDelivery(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/network/networking-and-content-delivery.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("networking-and-content-delivery", opts...)
}

func (c *awsNetworkContainer) Privatelink(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/network/privatelink.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("privatelink", opts...)
}

func (c *awsNetworkContainer) TransitGateway(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/network/transit-gateway.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("transit-gateway", opts...)
}

func (c *awsNetworkContainer) ClientVpn(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/network/client-vpn.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("client-vpn", opts...)
}

func (c *awsNetworkContainer) DirectConnect(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/network/direct-connect.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("direct-connect", opts...)
}

func (c *awsNetworkContainer) ElasticLoadBalancing(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/network/elastic-load-balancing.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("elastic-load-balancing", opts...)
}

func (c *awsNetworkContainer) GlobalAccelerator(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/network/global-accelerator.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("global-accelerator", opts...)
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
