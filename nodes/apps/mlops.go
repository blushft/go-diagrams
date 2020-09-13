package apps

import "github.com/blushft/go-diagrams/diagram"

type mlopsContainer struct {
	path string
	opts []diagram.NodeOption
}

var Mlops = &mlopsContainer{
	opts: diagram.OptionSet{diagram.Provider("apps"), diagram.NodeShape("none")},
	path: "assets/apps/mlops",
}

func (c *mlopsContainer) Polyaxon(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/apps/mlops/polyaxon.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
