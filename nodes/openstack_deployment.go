package nodes

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type deploymentContainer struct {
	path  string
	attrs []attr.Attribute
}

var OpenstackDeployment =&deploymentContainer{path: "assets/openstack/deployment"}

func (c *deploymentContainer) Tripleo(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/openstack/deployment/tripleo.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("tripleo", opts...)
}

func (c *deploymentContainer) Ansible(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/openstack/deployment/ansible.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("ansible", opts...)
}

func (c *deploymentContainer) Charms(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/openstack/deployment/charms.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("charms", opts...)
}

func (c *deploymentContainer) Chef(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/openstack/deployment/chef.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("chef", opts...)
}

func (c *deploymentContainer) Helm(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/openstack/deployment/helm.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("helm", opts...)
}

func (c *deploymentContainer) Kolla(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/openstack/deployment/kolla.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("kolla", opts...)
}

func init() {
  const namespace = "openstack.deployment"
  Register(namespace, "Tripleo", OpenstackDeployment.Tripleo)
  Register(namespace, "Ansible", OpenstackDeployment.Ansible)
  Register(namespace, "Charms", OpenstackDeployment.Charms)
  Register(namespace, "Chef", OpenstackDeployment.Chef)
  Register(namespace, "Helm", OpenstackDeployment.Helm)
  Register(namespace, "Kolla", OpenstackDeployment.Kolla)
}
