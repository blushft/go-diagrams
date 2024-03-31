package openstack

import "github.com/blushft/go-diagrams/diagram"

type deploymentContainer struct {
	path string
	opts []diagram.NodeOption
}

var Deployment = &deploymentContainer{
	opts: diagram.OptionSet{diagram.Provider("openstack"), diagram.NodeShape("none")},
	path: "assets/openstack/deployment",
}

func (c *deploymentContainer) Ansible(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/openstack/lifecyclemanagement/deployment/ansible.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *deploymentContainer) Charms(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/openstack/lifecyclemanagement/deployment/charms.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *deploymentContainer) Chef(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/openstack/lifecyclemanagement/deployment/chef.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *deploymentContainer) Helm(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/openstack/lifecyclemanagement/deployment/helm.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *deploymentContainer) Kolla(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/openstack/lifecyclemanagement/deployment/kolla.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *deploymentContainer) Tripleo(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/openstack/lifecyclemanagement/deployment/tripleo.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
