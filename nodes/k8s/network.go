package k8s

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type networkContainer struct {
	path  string
	attrs []attr.Attribute
}

var Network = &networkContainer{path: "assets/k8s/network"}

func (c *networkContainer) Svc(opts ...attr.Attribute) *node.Node {
	return node.New("svc", attr.AssetImage("assets/k8s/network/svc.png"), attr.Shape(attr.None))
}

func (c *networkContainer) Ep(opts ...attr.Attribute) *node.Node {
	return node.New("ep", attr.AssetImage("assets/k8s/network/ep.png"), attr.Shape(attr.None))
}

func (c *networkContainer) Ing(opts ...attr.Attribute) *node.Node {
	return node.New("ing", attr.AssetImage("assets/k8s/network/ing.png"), attr.Shape(attr.None))
}

func (c *networkContainer) Netpol(opts ...attr.Attribute) *node.Node {
	return node.New("netpol", attr.AssetImage("assets/k8s/network/netpol.png"), attr.Shape(attr.None))
}
