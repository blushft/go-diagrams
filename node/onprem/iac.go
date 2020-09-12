package onprem

import "github.com/blushft/go-diagrams/node"

type iacContainer struct {
	path string
	opts []node.Option
}

var Iac = &iacContainer{
	opts: node.OptionSet{node.Provider("onprem"), node.Shape("none")},
	path: "assets/onprem/iac",
}

func (c *iacContainer) Ansible(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/onprem/iac/ansible.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *iacContainer) Atlantis(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/onprem/iac/atlantis.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *iacContainer) Awx(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/onprem/iac/awx.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *iacContainer) Terraform(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/onprem/iac/terraform.png")}, c.opts, opts)
	return node.New(nopts...)
}
