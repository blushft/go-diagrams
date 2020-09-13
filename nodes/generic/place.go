package generic

import "github.com/blushft/go-diagrams/diagram"

type placeContainer struct {
	path string
	opts []diagram.NodeOption
}

var Place = &placeContainer{
	opts: diagram.OptionSet{diagram.Provider("generic"), diagram.NodeShape("none")},
	path: "assets/generic/place",
}

func (c *placeContainer) Datacenter(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/generic/place/datacenter.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
