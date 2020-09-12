package alibabacloud

import "github.com/blushft/go-diagrams/node"

type networkContainer struct {
	path string
	opts []node.Option
}

var Network = &networkContainer{
	opts: node.OptionSet{node.Provider("alibabacloud"), node.Shape("none")},
	path: "assets/alibabacloud/network",
}

func (c *networkContainer) Cdn(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/alibabacloud/network/cdn.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *networkContainer) ElasticIpAddress(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/alibabacloud/network/elastic-ip-address.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *networkContainer) ExpressConnect(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/alibabacloud/network/express-connect.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *networkContainer) VpnGateway(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/alibabacloud/network/vpn-gateway.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *networkContainer) CloudEnterpriseNetwork(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/alibabacloud/network/cloud-enterprise-network.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *networkContainer) NatGateway(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/alibabacloud/network/nat-gateway.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *networkContainer) ServerLoadBalancer(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/alibabacloud/network/server-load-balancer.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *networkContainer) SmartAccessGateway(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/alibabacloud/network/smart-access-gateway.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *networkContainer) VirtualPrivateCloud(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/alibabacloud/network/virtual-private-cloud.png")}, c.opts, opts)
	return node.New(nopts...)
}
