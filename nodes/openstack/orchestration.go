package openstack

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type orchestrationContainer struct {
	path  string
	attrs []attr.Attribute
}

var Orchestration = &orchestrationContainer{path: "assets/openstack/orchestration"}

func (c *orchestrationContainer) Blazar(opts ...attr.Attribute) *node.Node {
	return node.New("blazar", attr.AssetImage("assets/openstack/orchestration/blazar.png"), attr.Shape(attr.None))
}

func (c *orchestrationContainer) Heat(opts ...attr.Attribute) *node.Node {
	return node.New("heat", attr.AssetImage("assets/openstack/orchestration/heat.png"), attr.Shape(attr.None))
}

func (c *orchestrationContainer) Mistral(opts ...attr.Attribute) *node.Node {
	return node.New("mistral", attr.AssetImage("assets/openstack/orchestration/mistral.png"), attr.Shape(attr.None))
}

func (c *orchestrationContainer) Senlin(opts ...attr.Attribute) *node.Node {
	return node.New("senlin", attr.AssetImage("assets/openstack/orchestration/senlin.png"), attr.Shape(attr.None))
}

func (c *orchestrationContainer) Zaqar(opts ...attr.Attribute) *node.Node {
	return node.New("zaqar", attr.AssetImage("assets/openstack/orchestration/zaqar.png"), attr.Shape(attr.None))
}
