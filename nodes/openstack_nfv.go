package nodes

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type nfvContainer struct {
	path  string
	attrs []attr.Attribute
}

var Nfv = &nfvContainer{path: "assets/openstack/nfv"}

func (c *nfvContainer) Tacker(opts ...attr.Attribute) *node.Node {
	return node.New("tacker", attr.AssetImage("assets/openstack/nfv/tacker.png"), attr.Shape(attr.None))
}
