package onprem

import "github.com/blushft/go-diagrams/node"

type securityContainer struct {
	path string
	opts []node.Option
}

var Security = &securityContainer{
	opts: node.OptionSet{node.Provider("onprem"), node.Shape("none")},
	path: "assets/onprem/security",
}

func (c *securityContainer) Trivy(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/onprem/security/trivy.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *securityContainer) Vault(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/onprem/security/vault.png")}, c.opts, opts)
	return node.New(nopts...)
}
