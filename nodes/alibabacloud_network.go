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
	opts = append(opts, attr.AssetImage("assets/alibabacloud/network/express-connect.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("express-connect", opts...)
}

func (c *alibabaCloudNetworkContainer) NatGateway(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/alibabacloud/network/nat-gateway.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("nat-gateway", opts...)
}

func (c *alibabaCloudNetworkContainer) ServerLoadBalancer(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/alibabacloud/network/server-load-balancer.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("server-load-balancer", opts...)
}

func (c *alibabaCloudNetworkContainer) SmartAccessGateway(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/alibabacloud/network/smart-access-gateway.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("smart-access-gateway", opts...)
}

func (c *alibabaCloudNetworkContainer) VirtualPrivateCloud(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/alibabacloud/network/virtual-private-cloud.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("virtual-private-cloud", opts...)
}

func (c *alibabaCloudNetworkContainer) VpnGateway(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/alibabacloud/network/vpn-gateway.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("vpn-gateway", opts...)
}

func (c *alibabaCloudNetworkContainer) Cdn(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/alibabacloud/network/cdn.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("cdn", opts...)
}

func (c *alibabaCloudNetworkContainer) ElasticIpAddress(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/alibabacloud/network/elastic-ip-address.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("elastic-ip-address", opts...)
}

func (c *alibabaCloudNetworkContainer) CloudEnterpriseNetwork(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/alibabacloud/network/cloud-enterprise-network.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("cloud-enterprise-network", opts...)
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
