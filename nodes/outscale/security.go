package outscale

import "github.com/blushft/go-diagrams/diagram"

type securityContainer struct {
	path string
	opts []diagram.NodeOption
}

var Security = &securityContainer{
	opts: diagram.OptionSet{diagram.Provider("outscale"), diagram.NodeShape("none")},
	path: "assets/outscale/security",
}

func (c *securityContainer) Firewall(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/outscale/security/firewall.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *securityContainer) IdentityAndAccessManagement(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/outscale/security/identity-and-access-management.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
