package elastic

import "github.com/blushft/go-diagrams/diagram"

type securityContainer struct {
	path string
	opts []diagram.NodeOption
}

var Security = &securityContainer{
	opts: diagram.OptionSet{diagram.Provider("elastic"), diagram.NodeShape("none")},
	path: "assets/elastic/security",
}

func (c *securityContainer) Endpoint(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/elastic/security/endpoint.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *securityContainer) Security(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/elastic/security/security.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *securityContainer) Siem(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/elastic/security/siem.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
