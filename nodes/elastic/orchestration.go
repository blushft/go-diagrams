package elastic

import "github.com/blushft/go-diagrams/diagram"

type orchestrationContainer struct {
	path string
	opts []diagram.NodeOption
}

var Orchestration = &orchestrationContainer{
	opts: diagram.OptionSet{diagram.Provider("elastic"), diagram.NodeShape("none")},
	path: "assets/elastic/orchestration",
}

func (c *orchestrationContainer) Ece(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/elastic/orchestration/ece.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *orchestrationContainer) Eck(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/elastic/orchestration/eck.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
