package saas

import "github.com/blushft/go-diagrams/diagram"

type identityContainer struct {
	path string
	opts []diagram.NodeOption
}

var Identity = &identityContainer{
	opts: diagram.OptionSet{diagram.Provider("saas"), diagram.NodeShape("none")},
	path: "assets/saas/identity",
}

func (c *identityContainer) Auth0(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/saas/identity/auth0.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *identityContainer) Okta(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/saas/identity/okta.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
