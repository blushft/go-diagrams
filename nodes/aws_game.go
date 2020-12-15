package nodes

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type gameContainer struct {
	path  string
	attrs []attr.Attribute
}

var Game = &gameContainer{path: "assets/aws/game"}

func (c *gameContainer) Gamelift(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/game/gamelift.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("gamelift", opts...)
}

func init() {
  const namespace = "aws.game"
  Register(namespace, "Gamelift", Game.Gamelift)
}
