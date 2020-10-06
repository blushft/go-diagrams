package elastic

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type orchestrationContainer struct {
	path  string
	attrs []attr.Attribute
}

var Orchestration = &orchestrationContainer{path: "assets/elastic/orchestration"}

func (c *orchestrationContainer) Ece(opts ...attr.Attribute) *node.Node {
	return node.New("ece", attr.AssetImage("assets/elastic/orchestration/ece.png"), attr.Shape(attr.None))
}

func (c *orchestrationContainer) Eck(opts ...attr.Attribute) *node.Node {
	return node.New("eck", attr.AssetImage("assets/elastic/orchestration/eck.png"), attr.Shape(attr.None))
}
