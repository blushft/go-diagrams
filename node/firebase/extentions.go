package firebase

import "github.com/blushft/go-diagrams/node"

type extentionsContainer struct {
	path string
	opts []node.Option
}

var Extentions = &extentionsContainer{
	opts: node.OptionSet{node.Provider("firebase"), node.Shape("none")},
	path: "assets/firebase/extentions",
}

func (c *extentionsContainer) Extensions(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/firebase/extentions/extensions.png")}, c.opts, opts)
	return node.New(nopts...)
}
