package outscale

import "github.com/blushft/go-diagrams/node"

type securityContainer struct {
	path string
	opts []node.Option
}

var Security = &securityContainer{
	opts: node.OptionSet{node.Provider("outscale"), node.Shape("none")},
	path: "assets/outscale/security",
}

func (c *securityContainer) IdentityAndAccessManagement(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/outscale/security/identity-and-access-management.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *securityContainer) Firewall(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/outscale/security/firewall.png")}, c.opts, opts)
	return node.New(nopts...)
}
