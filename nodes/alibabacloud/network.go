package alibabacloud

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type networkContainer struct {
	path  string
	attrs []attr.Attribute
}

var Network = &networkContainer{path: "assets/alibabacloud/network"}

func (c *networkContainer) ExpressConnect(opts ...attr.Attribute) *node.Node {
	return node.New("express-connect", attr.AssetImage("assets/alibabacloud/network/express-connect.png"), attr.Shape(attr.None))
}

func (c *networkContainer) NatGateway(opts ...attr.Attribute) *node.Node {
	return node.New("nat-gateway", attr.AssetImage("assets/alibabacloud/network/nat-gateway.png"), attr.Shape(attr.None))
}

func (c *networkContainer) ServerLoadBalancer(opts ...attr.Attribute) *node.Node {
	return node.New("server-load-balancer", attr.AssetImage("assets/alibabacloud/network/server-load-balancer.png"), attr.Shape(attr.None))
}

func (c *networkContainer) SmartAccessGateway(opts ...attr.Attribute) *node.Node {
	return node.New("smart-access-gateway", attr.AssetImage("assets/alibabacloud/network/smart-access-gateway.png"), attr.Shape(attr.None))
}

func (c *networkContainer) VirtualPrivateCloud(opts ...attr.Attribute) *node.Node {
	return node.New("virtual-private-cloud", attr.AssetImage("assets/alibabacloud/network/virtual-private-cloud.png"), attr.Shape(attr.None))
}

func (c *networkContainer) VpnGateway(opts ...attr.Attribute) *node.Node {
	return node.New("vpn-gateway", attr.AssetImage("assets/alibabacloud/network/vpn-gateway.png"), attr.Shape(attr.None))
}

func (c *networkContainer) Cdn(opts ...attr.Attribute) *node.Node {
	return node.New("cdn", attr.AssetImage("assets/alibabacloud/network/cdn.png"), attr.Shape(attr.None))
}

func (c *networkContainer) ElasticIpAddress(opts ...attr.Attribute) *node.Node {
	return node.New("elastic-ip-address", attr.AssetImage("assets/alibabacloud/network/elastic-ip-address.png"), attr.Shape(attr.None))
}

func (c *networkContainer) CloudEnterpriseNetwork(opts ...attr.Attribute) *node.Node {
	return node.New("cloud-enterprise-network", attr.AssetImage("assets/alibabacloud/network/cloud-enterprise-network.png"), attr.Shape(attr.None))
}
