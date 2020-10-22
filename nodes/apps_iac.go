package nodes

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type iacContainer struct {
	path  string
	attrs []attr.Attribute
}

var Iac = &iacContainer{path: "assets/apps/iac"}

func (c *iacContainer) Ansible(opts ...attr.Attribute) *node.Node {
	return node.New("ansible", attr.AssetImage("assets/apps/iac/ansible.png"), attr.Shape(attr.None))
}

func (c *iacContainer) Atlantis(opts ...attr.Attribute) *node.Node {
	return node.New("atlantis", attr.AssetImage("assets/apps/iac/atlantis.png"), attr.Shape(attr.None))
}

func (c *iacContainer) Awx(opts ...attr.Attribute) *node.Node {
	return node.New("awx", attr.AssetImage("assets/apps/iac/awx.png"), attr.Shape(attr.None))
}

func (c *iacContainer) Terraform(opts ...attr.Attribute) *node.Node {
	return node.New("terraform", attr.AssetImage("assets/apps/iac/terraform.png"), attr.Shape(attr.None))
}

func init() {
	const namespace = "apps.iac"
	Register(namespace, "Ansible", Iac.Ansible)
	Register(namespace, "Atlantis", Iac.Atlantis)
	Register(namespace, "Awx", Iac.Awx)
	Register(namespace, "Terraform", Iac.Terraform)
}