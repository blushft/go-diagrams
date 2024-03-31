package apps

import "github.com/blushft/go-diagrams/diagram"

type cdContainer struct {
	path string
	opts []diagram.NodeOption
}

var Cd = &cdContainer{
	opts: diagram.OptionSet{diagram.Provider("apps"), diagram.NodeShape("none")},
	path: "assets/apps/cd",
}

func (c *cdContainer) TektonCli(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/apps/cd/tekton-cli.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *cdContainer) Tekton(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/apps/cd/tekton.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *cdContainer) Spinnaker(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/apps/cd/spinnaker.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
