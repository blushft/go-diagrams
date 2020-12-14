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
	return node.New("gamelift", attr.AssetImage("assets/aws/game/gamelift.png"), attr.Shape(attr.None))
}

func init() {
  const namespace = "aws.game"
  Register(namespace, "Gamelift", Game.Gamelift)
}
