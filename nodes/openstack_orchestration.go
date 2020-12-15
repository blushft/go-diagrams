package nodes

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type openstackOrchestrationContainer struct {
	path  string
	attrs []attr.Attribute
}

var OpenstackOrchestration = &openstackOrchestrationContainer{path: "assets/openstack/orchestration"}

func (c *openstackOrchestrationContainer) Blazar(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/openstack/orchestration/blazar.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("blazar", opts...)
}

func (c *openstackOrchestrationContainer) Heat(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/openstack/orchestration/heat.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("heat", opts...)
}

func (c *openstackOrchestrationContainer) Mistral(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/openstack/orchestration/mistral.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("mistral", opts...)
}

func (c *openstackOrchestrationContainer) Senlin(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/openstack/orchestration/senlin.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("senlin", opts...)
}

func (c *openstackOrchestrationContainer) Zaqar(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/openstack/orchestration/zaqar.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("zaqar", opts...)
}

func init() {
  const namespace = "openstack.orchestration"
  Register(namespace, "Blazar", OpenstackOrchestration.Blazar)
  Register(namespace, "Heat", OpenstackOrchestration.Heat)
  Register(namespace, "Mistral", OpenstackOrchestration.Mistral)
  Register(namespace, "Senlin", OpenstackOrchestration.Senlin)
  Register(namespace, "Zaqar", OpenstackOrchestration.Zaqar)
}
