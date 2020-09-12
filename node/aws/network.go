package aws

import "github.com/blushft/go-diagrams/node"

type networkContainer struct {
	path string
	opts []node.Option
}

var Network = &networkContainer{
	opts: node.OptionSet{node.Provider("aws"), node.Shape("none")},
	path: "assets/aws/network",
}

func (c *networkContainer) InternetGateway(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/network/internet-gateway.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *networkContainer) PrivateSubnet(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/network/private-subnet.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *networkContainer) VpcPeering(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/network/vpc-peering.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *networkContainer) AppMesh(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/network/app-mesh.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *networkContainer) Cloudfront(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/network/cloudfront.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *networkContainer) ElasticLoadBalancing(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/network/elastic-load-balancing.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *networkContainer) Endpoint(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/network/endpoint.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *networkContainer) NetworkingAndContentDelivery(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/network/networking-and-content-delivery.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *networkContainer) SiteToSiteVpn(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/network/site-to-site-vpn.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *networkContainer) TransitGateway(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/network/transit-gateway.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *networkContainer) ApiGateway(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/network/api-gateway.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *networkContainer) CloudMap(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/network/cloud-map.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *networkContainer) DirectConnect(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/network/direct-connect.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *networkContainer) NatGateway(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/network/nat-gateway.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *networkContainer) GlobalAccelerator(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/network/global-accelerator.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *networkContainer) RouteTable(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/network/route-table.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *networkContainer) Route53(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/network/route-53.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *networkContainer) VpcRouter(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/network/vpc-router.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *networkContainer) Vpc(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/network/vpc.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *networkContainer) ClientVpn(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/network/client-vpn.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *networkContainer) Nacl(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/network/nacl.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *networkContainer) Privatelink(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/network/privatelink.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *networkContainer) PublicSubnet(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/network/public-subnet.png")}, c.opts, opts)
	return node.New(nopts...)
}
