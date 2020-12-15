package nodes

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type multiregionContainer struct {
	path  string
	attrs []attr.Attribute
}

var OpenstackMultiregion =&multiregionContainer{path: "assets/openstack/multiregion"}

func (c *multiregionContainer) Tricircle(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/openstack/multiregion/tricircle.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("tricircle", opts...)
}

func init() {
  const namespace = "openstack.multiregion"
  Register(namespace, "Tricircle", OpenstackMultiregion.Tricircle)
}
