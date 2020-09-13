package oci

import "github.com/blushft/go-diagrams/diagram"

type connectivityContainer struct {
	path string
	opts []diagram.NodeOption
}

var Connectivity = &connectivityContainer{
	opts: diagram.OptionSet{diagram.Provider("oci"), diagram.NodeShape("none")},
	path: "assets/oci/connectivity",
}

func (c *connectivityContainer) Dns(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/oci/connectivity/dns.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *connectivityContainer) DisconnectedRegions(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/oci/connectivity/disconnected-regions.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *connectivityContainer) CustomerPremise(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/oci/connectivity/customer-premise.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *connectivityContainer) CustomerDatacntrWhite(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/oci/connectivity/customer-datacntr-white.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *connectivityContainer) CustomerPremiseWhite(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/oci/connectivity/customer-premise-white.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *connectivityContainer) DnsWhite(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/oci/connectivity/dns-white.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *connectivityContainer) FastConnectWhite(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/oci/connectivity/fast-connect-white.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *connectivityContainer) NatGateway(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/oci/connectivity/nat-gateway.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *connectivityContainer) Vpn(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/oci/connectivity/vpn.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *connectivityContainer) CustomerDatacenter(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/oci/connectivity/customer-datacenter.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *connectivityContainer) Backbone(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/oci/connectivity/backbone.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *connectivityContainer) CdnWhite(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/oci/connectivity/cdn-white.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *connectivityContainer) Cdn(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/oci/connectivity/cdn.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *connectivityContainer) DisconnectedRegionsWhite(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/oci/connectivity/disconnected-regions-white.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *connectivityContainer) FastConnect(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/oci/connectivity/fast-connect.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *connectivityContainer) NatGatewayWhite(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/oci/connectivity/nat-gateway-white.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *connectivityContainer) VpnWhite(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/oci/connectivity/vpn-white.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *connectivityContainer) BackboneWhite(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/oci/connectivity/backbone-white.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
