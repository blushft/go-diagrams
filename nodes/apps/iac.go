package apps

import "github.com/blushft/go-diagrams/diagram"

type iacContainer struct {
	path string
	opts []diagram.NodeOption
}

var Iac = &iacContainer{
	opts: diagram.OptionSet{diagram.Provider("apps"), diagram.NodeShape("none")},
	path: "assets/apps/iac",
}

func (c *iacContainer) Awx(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/apps/iac/awx.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *iacContainer) Terraform(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/apps/iac/terraform.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *iacContainer) Ansible(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/apps/iac/ansible.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *iacContainer) Atlantis(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/apps/iac/atlantis.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
