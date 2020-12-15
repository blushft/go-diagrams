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
	opts = append(opts, attr.AssetImage("assets/apps/iac/ansible.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("ansible", opts...)
}

func (c *iacContainer) Atlantis(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/apps/iac/atlantis.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("atlantis", opts...)
}

func (c *iacContainer) Awx(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/apps/iac/awx.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("awx", opts...)
}

func (c *iacContainer) Terraform(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/apps/iac/terraform.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("terraform", opts...)
}

func init() {
	const namespace = "apps.iac"
	Register(namespace, "Ansible", Iac.Ansible)
	Register(namespace, "Atlantis", Iac.Atlantis)
	Register(namespace, "Awx", Iac.Awx)
	Register(namespace, "Terraform", Iac.Terraform)
}