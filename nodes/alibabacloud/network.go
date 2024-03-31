package alibabacloud

import "github.com/blushft/go-diagrams/diagram"

type networkContainer struct {
	path string
	opts []diagram.NodeOption
}

var Network = &networkContainer{
	opts: diagram.OptionSet{diagram.Provider("alibabacloud"), diagram.NodeShape("none")},
	path: "assets/alibabacloud/network",
}

func (c *networkContainer) CloudEnterpriseNetwork(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/alibabacloud/network/cloud-enterprise-network.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *networkContainer) ElasticIpAddress(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/alibabacloud/network/elastic-ip-address.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *networkContainer) NatGateway(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/alibabacloud/network/nat-gateway.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *networkContainer) SmartAccessGateway(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/alibabacloud/network/smart-access-gateway.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *networkContainer) VirtualPrivateCloud(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/alibabacloud/network/virtual-private-cloud.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *networkContainer) Cdn(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/alibabacloud/network/cdn.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *networkContainer) ExpressConnect(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/alibabacloud/network/express-connect.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *networkContainer) ServerLoadBalancer(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/alibabacloud/network/server-load-balancer.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *networkContainer) VpnGateway(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/alibabacloud/network/vpn-gateway.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
