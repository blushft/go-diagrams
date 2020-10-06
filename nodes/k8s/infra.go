package k8s

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type infraContainer struct {
	path  string
	attrs []attr.Attribute
}

var Infra = &infraContainer{path: "assets/k8s/infra"}

func (c *infraContainer) Etcd(opts ...attr.Attribute) *node.Node {
	return node.New("etcd", attr.AssetImage("assets/k8s/infra/etcd.png"), attr.Shape(attr.None))
}

func (c *infraContainer) Master(opts ...attr.Attribute) *node.Node {
	return node.New("master", attr.AssetImage("assets/k8s/infra/master.png"), attr.Shape(attr.None))
}

func (c *infraContainer) Node(opts ...attr.Attribute) *node.Node {
	return node.New("node", attr.AssetImage("assets/k8s/infra/node.png"), attr.Shape(attr.None))
}
