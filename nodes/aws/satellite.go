package aws

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type satelliteContainer struct {
	path  string
	attrs []attr.Attribute
}

var Satellite = &satelliteContainer{path: "assets/aws/satellite"}

func (c *satelliteContainer) GroundStation(opts ...attr.Attribute) *node.Node {
	return node.New("ground-station", attr.AssetImage("assets/aws/satellite/ground-station.png"), attr.Shape(attr.None))
}
