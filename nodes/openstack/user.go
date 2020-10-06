package openstack

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type userContainer struct {
	path  string
	attrs []attr.Attribute
}

var User = &userContainer{path: "assets/openstack/user"}

func (c *userContainer) Openstackclient(opts ...attr.Attribute) *node.Node {
	return node.New("openstackclient", attr.AssetImage("assets/openstack/user/openstackclient.png"), attr.Shape(attr.None))
}
