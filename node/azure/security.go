package azure

import "github.com/blushft/go-diagrams/node"

type securityContainer struct {
	path string
	opts []node.Option
}

var Security = &securityContainer{
	opts: node.OptionSet{node.Provider("azure"), node.Shape("none")},
	path: "assets/azure/security",
}

func (c *securityContainer) KeyVaults(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/azure/security/key-vaults.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *securityContainer) SecurityCenter(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/azure/security/security-center.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *securityContainer) Sentinel(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/azure/security/sentinel.png")}, c.opts, opts)
	return node.New(nopts...)
}
