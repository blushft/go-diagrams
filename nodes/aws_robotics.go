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
	opts = append(opts, attr.AssetImage("assets/aws/robotics/robomaker-simulator.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("robomaker-simulator", opts...)
}

func (c *roboticsContainer) Robomaker(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/robotics/robomaker.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("robomaker", opts...)
}

func (c *roboticsContainer) Robotics(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/robotics/robotics.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("robotics", opts...)
}

func init() {
  const namespace= "aws.robotics"
  Register(namespace, "RobomakerSimulator", Robotics.RobomakerSimulator)
  Register(namespace, "Robomaker", Robotics.Robomaker)
  Register(namespace, "Robotics", Robotics.Robotics)
}
