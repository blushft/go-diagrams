package aws

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type businessContainer struct {
	path  string
	attrs []attr.Attribute
}

var Business = &businessContainer{path: "assets/aws/business"}

func (c *businessContainer) AlexaForBusiness(opts ...attr.Attribute) *node.Node {
	return node.New("alexa-for-business", attr.AssetImage("assets/aws/business/alexa-for-business.png"), attr.Shape(attr.None))
}

func (c *businessContainer) Chime(opts ...attr.Attribute) *node.Node {
	return node.New("chime", attr.AssetImage("assets/aws/business/chime.png"), attr.Shape(attr.None))
}

func (c *businessContainer) Workmail(opts ...attr.Attribute) *node.Node {
	return node.New("workmail", attr.AssetImage("assets/aws/business/workmail.png"), attr.Shape(attr.None))
}
