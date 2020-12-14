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
	return node.New("blazar", attr.AssetImage("assets/openstack/orchestration/blazar.png"), attr.Shape(attr.None))
}

func (c *openstackOrchestrationContainer) Heat(opts ...attr.Attribute) *node.Node {
	return node.New("heat", attr.AssetImage("assets/openstack/orchestration/heat.png"), attr.Shape(attr.None))
}

func (c *openstackOrchestrationContainer) Mistral(opts ...attr.Attribute) *node.Node {
	return node.New("mistral", attr.AssetImage("assets/openstack/orchestration/mistral.png"), attr.Shape(attr.None))
}

func (c *openstackOrchestrationContainer) Senlin(opts ...attr.Attribute) *node.Node {
	return node.New("senlin", attr.AssetImage("assets/openstack/orchestration/senlin.png"), attr.Shape(attr.None))
}

func (c *openstackOrchestrationContainer) Zaqar(opts ...attr.Attribute) *node.Node {
	return node.New("zaqar", attr.AssetImage("assets/openstack/orchestration/zaqar.png"), attr.Shape(attr.None))
}

func init() {
  const namespace = "openstack.orchestration"
  Register(namespace, "Blazar", OpenstackOrchestration.Blazar)
  Register(namespace, "Heat", OpenstackOrchestration.Heat)
  Register(namespace, "Mistral", OpenstackOrchestration.Mistral)
  Register(namespace, "Senlin", OpenstackOrchestration.Senlin)
  Register(namespace, "Zaqar", OpenstackOrchestration.Zaqar)
}
