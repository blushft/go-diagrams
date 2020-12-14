package nodes

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type roboticsContainer struct {
	path  string
	attrs []attr.Attribute
}

var Robotics = &roboticsContainer{path: "assets/aws/robotics"}

func (c *roboticsContainer) RobomakerSimulator(opts ...attr.Attribute) *node.Node {
	return node.New("robomaker-simulator", attr.AssetImage("assets/aws/robotics/robomaker-simulator.png"), attr.Shape(attr.None))
}

func (c *roboticsContainer) Robomaker(opts ...attr.Attribute) *node.Node {
	return node.New("robomaker", attr.AssetImage("assets/aws/robotics/robomaker.png"), attr.Shape(attr.None))
}

func (c *roboticsContainer) Robotics(opts ...attr.Attribute) *node.Node {
	return node.New("robotics", attr.AssetImage("assets/aws/robotics/robotics.png"), attr.Shape(attr.None))
}

func init() {
  const namespace= "aws.robotics"
  Register(namespace, "RobomakerSimulator", Robotics.RobomakerSimulator)
  Register(namespace, "Robomaker", Robotics.Robomaker)
  Register(namespace, "Robotics", Robotics.Robotics)
}
