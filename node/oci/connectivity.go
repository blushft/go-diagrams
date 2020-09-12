package oci

import "github.com/blushft/go-diagrams/node"

type connectivityContainer struct {
	path string
	opts []node.Option
}

var Connectivity = &connectivityContainer{
	opts: node.OptionSet{node.Provider("oci"), node.Shape("none")},
	path: "assets/oci/connectivity",
}

func (c *connectivityContainer) FastConnectWhite(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/oci/connectivity/fast-connect-white.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *connectivityContainer) VpnWhite(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/oci/connectivity/vpn-white.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *connectivityContainer) Backbone(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/oci/connectivity/backbone.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *connectivityContainer) CustomerPremiseWhite(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/oci/connectivity/customer-premise-white.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *connectivityContainer) DisconnectedRegions(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/oci/connectivity/disconnected-regions.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *connectivityContainer) DnsWhite(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/oci/connectivity/dns-white.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *connectivityContainer) Dns(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/oci/connectivity/dns.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *connectivityContainer) FastConnect(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/oci/connectivity/fast-connect.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *connectivityContainer) Cdn(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/oci/connectivity/cdn.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *connectivityContainer) CustomerDatacenter(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/oci/connectivity/customer-datacenter.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *connectivityContainer) CustomerDatacntrWhite(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/oci/connectivity/customer-datacntr-white.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *connectivityContainer) DisconnectedRegionsWhite(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/oci/connectivity/disconnected-regions-white.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *connectivityContainer) NatGateway(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/oci/connectivity/nat-gateway.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *connectivityContainer) Vpn(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/oci/connectivity/vpn.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *connectivityContainer) BackboneWhite(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/oci/connectivity/backbone-white.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *connectivityContainer) CdnWhite(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/oci/connectivity/cdn-white.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *connectivityContainer) CustomerPremise(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/oci/connectivity/customer-premise.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *connectivityContainer) NatGatewayWhite(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/oci/connectivity/nat-gateway-white.png")}, c.opts, opts)
	return node.New(nopts...)
}
