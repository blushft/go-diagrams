package generic

import "github.com/blushft/go-diagrams/node"

type placeContainer struct {
	path string
	opts []node.Option
}

var Place = &placeContainer{
	opts: node.OptionSet{node.Provider("generic"), node.Shape("none")},
	path: "assets/generic/place",
}

func (c *placeContainer) Datacenter(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/generic/place/datacenter.png")}, c.opts, opts)
	return node.New(nopts...)
}
