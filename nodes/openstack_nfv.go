package nodes

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type nfvContainer struct {
	path  string
	attrs []attr.Attribute
}

var OpenstackNfv =&nfvContainer{path: "assets/openstack/nfv"}

func (c *nfvContainer) Tacker(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/openstack/nfv/tacker.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("tacker", opts...)
}

func init() {
  const namespace = "openstack.nfv"
  Register(namespace, "Tacker", OpenstackNfv.Tacker)
}
