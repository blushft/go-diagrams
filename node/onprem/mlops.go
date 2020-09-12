package onprem

import "github.com/blushft/go-diagrams/node"

type mlopsContainer struct {
	path string
	opts []node.Option
}

var Mlops = &mlopsContainer{
	opts: node.OptionSet{node.Provider("onprem"), node.Shape("none")},
	path: "assets/onprem/mlops",
}

func (c *mlopsContainer) Polyaxon(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/onprem/mlops/polyaxon.png")}, c.opts, opts)
	return node.New(nopts...)
}
