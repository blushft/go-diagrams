package elastic

import "github.com/blushft/go-diagrams/diagram"

type saasContainer struct {
	path string
	opts []diagram.NodeOption
}

var Saas = &saasContainer{
	opts: diagram.OptionSet{diagram.Provider("elastic"), diagram.NodeShape("none")},
	path: "assets/elastic/saas",
}

func (c *saasContainer) Cloud(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/elastic/saas/cloud.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *saasContainer) Elastic(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/elastic/saas/elastic.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
