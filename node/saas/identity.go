package saas

import "github.com/blushft/go-diagrams/node"

type identityContainer struct {
	path string
	opts []node.Option
}

var Identity = &identityContainer{
	opts: node.OptionSet{node.Provider("saas"), node.Shape("none")},
	path: "assets/saas/identity",
}

func (c *identityContainer) Auth0(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/saas/identity/auth0.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *identityContainer) Okta(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/saas/identity/okta.png")}, c.opts, opts)
	return node.New(nopts...)
}
