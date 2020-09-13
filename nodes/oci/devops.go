package oci

import "github.com/blushft/go-diagrams/diagram"

type devopsContainer struct {
	path string
	opts []diagram.NodeOption
}

var Devops = &devopsContainer{
	opts: diagram.OptionSet{diagram.Provider("oci"), diagram.NodeShape("none")},
	path: "assets/oci/devops",
}

func (c *devopsContainer) ResourceMgmt(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/oci/devops/resource-mgmt.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *devopsContainer) ApiGatewayWhite(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/oci/devops/api-gateway-white.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *devopsContainer) ApiGateway(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/oci/devops/api-gateway.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *devopsContainer) ApiServiceWhite(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/oci/devops/api-service-white.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *devopsContainer) ApiService(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/oci/devops/api-service.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *devopsContainer) ResourceMgmtWhite(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/oci/devops/resource-mgmt-white.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
