package nodes

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type enablementContainer struct {
	path  string
	attrs []attr.Attribute
}

var Enablement = &enablementContainer{path: "assets/aws/enablement"}

func (c *enablementContainer) Iq(opts ...attr.Attribute) *node.Node {
	return node.New("iq", attr.AssetImage("assets/aws/enablement/iq.png"), attr.Shape(attr.None))
}

func (c *enablementContainer) ManagedServices(opts ...attr.Attribute) *node.Node {
	return node.New("managed-services", attr.AssetImage("assets/aws/enablement/managed-services.png"), attr.Shape(attr.None))
}

func (c *enablementContainer) ProfessionalServices(opts ...attr.Attribute) *node.Node {
	return node.New("professional-services", attr.AssetImage("assets/aws/enablement/professional-services.png"), attr.Shape(attr.None))
}

func (c *enablementContainer) Support(opts ...attr.Attribute) *node.Node {
	return node.New("support", attr.AssetImage("assets/aws/enablement/support.png"), attr.Shape(attr.None))
}

func init() {
  const namespace = "aws.enablement"
  Register(namespace, "Iq", Enablement.Iq)
  Register(namespace, "ManagedServices", Enablement.ManagedServices)
  Register(namespace, "ProfessionalServices", Enablement.ProfessionalServices)
  Register(namespace, "Support", Enablement.Support)
}
