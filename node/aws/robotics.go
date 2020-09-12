package aws

import "github.com/blushft/go-diagrams/node"

type roboticsContainer struct {
	path string
	opts []node.Option
}

var Robotics = &roboticsContainer{
	opts: node.OptionSet{node.Provider("aws"), node.Shape("none")},
	path: "assets/aws/robotics",
}

func (c *roboticsContainer) RobomakerSimulator(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/robotics/robomaker-simulator.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *roboticsContainer) Robomaker(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/robotics/robomaker.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *roboticsContainer) Robotics(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/robotics/robotics.png")}, c.opts, opts)
	return node.New(nopts...)
}
