package aws

import "github.com/blushft/go-diagrams/diagram"

type satelliteContainer struct {
	path string
	opts []diagram.NodeOption
}

var Satellite = &satelliteContainer{
	opts: diagram.OptionSet{diagram.Provider("aws"), diagram.NodeShape("none")},
	path: "assets/aws/satellite",
}

func (c *satelliteContainer) GroundStation(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/satellite/ground-station.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
