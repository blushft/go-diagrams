package oci

import "github.com/blushft/go-diagrams/node"

type devopsContainer struct {
	path string
	opts []node.Option
}

var Devops = &devopsContainer{
	opts: node.OptionSet{node.Provider("oci"), node.Shape("none")},
	path: "assets/oci/devops",
}

func (c *devopsContainer) ApiGatewayWhite(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/oci/devops/api-gateway-white.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *devopsContainer) ApiGateway(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/oci/devops/api-gateway.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *devopsContainer) ApiServiceWhite(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/oci/devops/api-service-white.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *devopsContainer) ApiService(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/oci/devops/api-service.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *devopsContainer) ResourceMgmtWhite(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/oci/devops/resource-mgmt-white.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *devopsContainer) ResourceMgmt(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/oci/devops/resource-mgmt.png")}, c.opts, opts)
	return node.New(nopts...)
}
