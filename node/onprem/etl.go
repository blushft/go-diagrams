package onprem

import "github.com/blushft/go-diagrams/node"

type etlContainer struct {
	path string
	opts []node.Option
}

var Etl = &etlContainer{
	opts: node.OptionSet{node.Provider("onprem"), node.Shape("none")},
	path: "assets/onprem/etl",
}

func (c *etlContainer) Embulk(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/onprem/etl/embulk.png")}, c.opts, opts)
	return node.New(nopts...)
}
