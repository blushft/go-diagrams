package aws

import "github.com/blushft/go-diagrams/node"

type gameContainer struct {
	path string
	opts []node.Option
}

var Game = &gameContainer{
	opts: node.OptionSet{node.Provider("aws"), node.Shape("none")},
	path: "assets/aws/game",
}

func (c *gameContainer) Gamelift(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/game/gamelift.png")}, c.opts, opts)
	return node.New(nopts...)
}
