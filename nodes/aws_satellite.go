package nodes

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
	opts = append(opts, attr.AssetImage("assets/aws/satellite/ground-station.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("ground-station", opts...)
}

func init() {
  const namespace= "aws.satellite"
  Register(namespace, "GroundStation", Satellite.GroundStation)
}
