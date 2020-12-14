package nodes

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type alibabaCloudNetworkContainer struct {
	path  string
	attrs []attr.Attribute
}

var AlibabacloudNetwork =&alibabaCloudNetworkContainer{path: "assets/alibabacloud/network"}

func (c *alibabaCloudNetworkContainer) ExpressConnect(opts ...attr.Attribute) *node.Node {
	return node.New("express-connect", attr.AssetImage("assets/alibabacloud/network/express-connect.png"), attr.Shape(attr.None))
}

func (c *alibabaCloudNetworkContainer) NatGateway(opts ...attr.Attribute) *node.Node {
	return node.New("nat-gateway", attr.AssetImage("assets/alibabacloud/network/nat-gateway.png"), attr.Shape(attr.None))
}

func (c *alibabaCloudNetworkContainer) ServerLoadBalancer(opts ...attr.Attribute) *node.Node {
	return node.New("server-load-balancer", attr.AssetImage("assets/alibabacloud/network/server-load-balancer.png"), attr.Shape(attr.None))
}

func (c *alibabaCloudNetworkContainer) SmartAccessGateway(opts ...attr.Attribute) *node.Node {
	return node.New("smart-access-gateway", attr.AssetImage("assets/alibabacloud/network/smart-access-gateway.png"), attr.Shape(attr.None))
}

func (c *alibabaCloudNetworkContainer) VirtualPrivateCloud(opts ...attr.Attribute) *node.Node {
	return node.New("virtual-private-cloud", attr.AssetImage("assets/alibabacloud/network/virtual-private-cloud.png"), attr.Shape(attr.None))
}

func (c *alibabaCloudNetworkContainer) VpnGateway(opts ...attr.Attribute) *node.Node {
	return node.New("vpn-gateway", attr.AssetImage("assets/alibabacloud/network/vpn-gateway.png"), attr.Shape(attr.None))
}

func (c *alibabaCloudNetworkContainer) Cdn(opts ...attr.Attribute) *node.Node {
	return node.New("cdn", attr.AssetImage("assets/alibabacloud/network/cdn.png"), attr.Shape(attr.None))
}

func (c *alibabaCloudNetworkContainer) ElasticIpAddress(opts ...attr.Attribute) *node.Node {
	return node.New("elastic-ip-address", attr.AssetImage("assets/alibabacloud/network/elastic-ip-address.png"), attr.Shape(attr.None))
}

func (c *alibabaCloudNetworkContainer) CloudEnterpriseNetwork(opts ...attr.Attribute) *node.Node {
	return node.New("cloud-enterprise-network", attr.AssetImage("assets/alibabacloud/network/cloud-enterprise-network.png"), attr.Shape(attr.None))
}

func init() {
  const namespace = "alibabacloud.network"
  Register(namespace, "ExpressConnect", AlibabacloudNetwork.ExpressConnect)
  Register(namespace, "NatGateway", AlibabacloudNetwork.NatGateway)
  Register(namespace, "ServerLoadBalancer", AlibabacloudNetwork.ServerLoadBalancer)
  Register(namespace, "SmartAccessGateway", AlibabacloudNetwork.SmartAccessGateway)
  Register(namespace, "VirtualPrivateCloud", AlibabacloudNetwork.VirtualPrivateCloud)
  Register(namespace, "VpnGateway", AlibabacloudNetwork.VpnGateway)
  Register(namespace, "Cdn", AlibabacloudNetwork.Cdn)
  Register(namespace, "ElasticIpAddress", AlibabacloudNetwork.ElasticIpAddress)
  Register(namespace, "CloudEnterpriseNetwork", AlibabacloudNetwork.CloudEnterpriseNetwork)
}
