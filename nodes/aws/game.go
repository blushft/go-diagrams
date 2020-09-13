package aws

import "github.com/blushft/go-diagrams/diagram"

type gameContainer struct {
	path string
	opts []diagram.NodeOption
}

var Game = &gameContainer{
	opts: diagram.OptionSet{diagram.Provider("aws"), diagram.NodeShape("none")},
	path: "assets/aws/game",
}

func (c *gameContainer) Gamelift(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/game/gamelift.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
