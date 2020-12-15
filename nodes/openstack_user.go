package nodes

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type userContainer struct {
	path  string
	attrs []attr.Attribute
}

var OpenstackUser =&userContainer{path: "assets/openstack/user"}

func (c *userContainer) Openstackclient(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/openstack/user/openstackclient.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("openstackclient", opts...)
}

func init() {
  const namespace = "openstack.user"
  Register(namespace, "Openstackclient", OpenstackUser.Openstackclient)
}
