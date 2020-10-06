package k8s

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type ecosystemContainer struct {
	path  string
	attrs []attr.Attribute
}

var Ecosystem = &ecosystemContainer{path: "assets/k8s/ecosystem"}

func (c *ecosystemContainer) Helm(opts ...attr.Attribute) *node.Node {
	return node.New("helm", attr.AssetImage("assets/k8s/ecosystem/helm.png"), attr.Shape(attr.None))
}

func (c *ecosystemContainer) Krew(opts ...attr.Attribute) *node.Node {
	return node.New("krew", attr.AssetImage("assets/k8s/ecosystem/krew.png"), attr.Shape(attr.None))
}

func (c *ecosystemContainer) Kustomize(opts ...attr.Attribute) *node.Node {
	return node.New("kustomize", attr.AssetImage("assets/k8s/ecosystem/kustomize.png"), attr.Shape(attr.None))
}
