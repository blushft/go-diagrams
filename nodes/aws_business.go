package nodes

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type awsBusinessContainer struct {
	path  string
	attrs []attr.Attribute
}

var AWSBusiness = &awsBusinessContainer{path: "assets/aws/business"}

func (c *awsBusinessContainer) AlexaForBusiness(opts ...attr.Attribute) *node.Node {
	return node.New("alexa-for-business", attr.AssetImage("assets/aws/business/alexa-for-business.png"), attr.Shape(attr.None))
}

func (c *awsBusinessContainer) Chime(opts ...attr.Attribute) *node.Node {
	return node.New("chime", attr.AssetImage("assets/aws/business/chime.png"), attr.Shape(attr.None))
}

func (c *awsBusinessContainer) Workmail(opts ...attr.Attribute) *node.Node {
	return node.New("workmail", attr.AssetImage("assets/aws/business/workmail.png"), attr.Shape(attr.None))
}

func init() {
  const namespace = "aws.business"
  Register(namespace, "AlexaForBusiness", AWSBusiness.AlexaForBusiness)
  Register(namespace, "Chime", AWSBusiness.Chime)
  Register(namespace, "Workmail", AWSBusiness.Workmail)
}
