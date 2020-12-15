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
	opts = append(opts, attr.AssetImage("assets/aws/business/alexa-for-business.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("alexa-for-business", opts...)
}

func (c *awsBusinessContainer) Chime(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/business/chime.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("chime", opts...)
}

func (c *awsBusinessContainer) Workmail(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/business/workmail.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("workmail", opts...)
}

func init() {
  const namespace = "aws.business"
  Register(namespace, "AlexaForBusiness", AWSBusiness.AlexaForBusiness)
  Register(namespace, "Chime", AWSBusiness.Chime)
  Register(namespace, "Workmail", AWSBusiness.Workmail)
}
