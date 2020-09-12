package openstack

import "github.com/blushft/go-diagrams/node"

type deploymentContainer struct {
	path string
	opts []node.Option
}

var Deployment = &deploymentContainer{
	opts: node.OptionSet{node.Provider("openstack"), node.Shape("none")},
	path: "assets/openstack/deployment",
}

func (c *deploymentContainer) Ansible(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/openstack/lifecyclemanagement/deployment/ansible.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *deploymentContainer) Charms(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/openstack/lifecyclemanagement/deployment/charms.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *deploymentContainer) Chef(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/openstack/lifecyclemanagement/deployment/chef.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *deploymentContainer) Helm(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/openstack/lifecyclemanagement/deployment/helm.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *deploymentContainer) Kolla(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/openstack/lifecyclemanagement/deployment/kolla.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *deploymentContainer) Tripleo(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/openstack/lifecyclemanagement/deployment/tripleo.png")}, c.opts, opts)
	return node.New(nopts...)
}
