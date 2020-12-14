package nodes

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type awsArContainer struct {
	path  string
	attrs []attr.Attribute
}

var AWSAr = &awsArContainer{path: "assets/aws/ar"}

func (c *awsArContainer) Sumerian(opts ...attr.Attribute) *node.Node {
	return node.New("sumerian", attr.AssetImage("assets/aws/ar/sumerian.png"), attr.Shape(attr.None))
}

func init() {
  const namespace = "aws.ar"
  Register(namespace, "Sumerian", AWSAr.Sumerian)
}
