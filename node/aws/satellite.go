package aws

import "github.com/blushft/go-diagrams/node"

type satelliteContainer struct {
	path string
	opts []node.Option
}

var Satellite = &satelliteContainer{
	opts: node.OptionSet{node.Provider("aws"), node.Shape("none")},
	path: "assets/aws/satellite",
}

func (c *satelliteContainer) GroundStation(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/satellite/ground-station.png")}, c.opts, opts)
	return node.New(nopts...)
}
