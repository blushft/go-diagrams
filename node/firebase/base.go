package firebase

import "github.com/blushft/go-diagrams/node"

type baseContainer struct {
	path string
	opts []node.Option
}

var Base = &baseContainer{
	opts: node.OptionSet{node.Provider("firebase"), node.Shape("none")},
	path: "assets/firebase/base",
}

func (c *baseContainer) Firebase(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/firebase/base/firebase.png")}, c.opts, opts)
	return node.New(nopts...)
}
