package nodes

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
	opts = append(opts, attr.AssetImage("assets/elastic/orchestration/ece.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("ece", opts...)
}

func (c *orchestrationContainer) Eck(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/elastic/orchestration/eck.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("eck", opts...)
}
