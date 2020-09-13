package apps

import "github.com/blushft/go-diagrams/diagram"

type securityContainer struct {
	path string
	opts []diagram.NodeOption
}

var Security = &securityContainer{
	opts: diagram.OptionSet{diagram.Provider("apps"), diagram.NodeShape("none")},
	path: "assets/apps/security",
}

func (c *securityContainer) Trivy(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/apps/security/trivy.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *securityContainer) Vault(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/apps/security/vault.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
