package nodes

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type frontendContainer struct {
	path  string
	attrs []attr.Attribute
}

var OpenstackFrontend =&frontendContainer{path: "assets/openstack/frontend"}

func (c *frontendContainer) Horizon(opts ...attr.Attribute) *node.Node {
	return node.New("horizon", attr.AssetImage("assets/openstack/frontend/horizon.png"), attr.Shape(attr.None))
}

func init() {
  const namespace = "openstack.frontend"
  Register(namespace, "Horizon", OpenstackFrontend.Horizon)
}
