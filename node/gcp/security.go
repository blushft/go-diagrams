package gcp

import "github.com/blushft/go-diagrams/node"

type securityContainer struct {
	path string
	opts []node.Option
}

var Security = &securityContainer{
	opts: node.OptionSet{node.Provider("gcp"), node.Shape("none")},
	path: "assets/gcp/security",
}

func (c *securityContainer) SecurityCommandCenter(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/gcp/security/security-command-center.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *securityContainer) SecurityScanner(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/gcp/security/security-scanner.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *securityContainer) Iam(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/gcp/security/iam.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *securityContainer) Iap(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/gcp/security/iap.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *securityContainer) KeyManagementService(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/gcp/security/key-management-service.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *securityContainer) ResourceManager(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/gcp/security/resource-manager.png")}, c.opts, opts)
	return node.New(nopts...)
}
