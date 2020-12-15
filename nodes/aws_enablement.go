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
	opts = append(opts, attr.AssetImage("assets/aws/enablement/iq.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("iq", opts...)
}

func (c *enablementContainer) ManagedServices(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/enablement/managed-services.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("managed-services", opts...)
}

func (c *enablementContainer) ProfessionalServices(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/enablement/professional-services.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("professional-services", opts...)
}

func (c *enablementContainer) Support(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/enablement/support.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("support", opts...)
}

func init() {
  const namespace = "aws.enablement"
  Register(namespace, "Iq", Enablement.Iq)
  Register(namespace, "ManagedServices", Enablement.ManagedServices)
  Register(namespace, "ProfessionalServices", Enablement.ProfessionalServices)
  Register(namespace, "Support", Enablement.Support)
}
