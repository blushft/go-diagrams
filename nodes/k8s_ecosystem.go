package nodes

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type k8sEcosystemContainer struct {
	path  string
	attrs []attr.Attribute
}

var K8sEcosystem = &k8sEcosystemContainer{path: "assets/k8s/ecosystem"}

func (c *k8sEcosystemContainer) Helm(opts ...attr.Attribute) *node.Node {
	return node.New("helm", attr.AssetImage("assets/k8s/ecosystem/helm.png"), attr.Shape(attr.None))
}

func (c *k8sEcosystemContainer) Krew(opts ...attr.Attribute) *node.Node {
	return node.New("krew", attr.AssetImage("assets/k8s/ecosystem/krew.png"), attr.Shape(attr.None))
}

func (c *k8sEcosystemContainer) Kustomize(opts ...attr.Attribute) *node.Node {
	return node.New("kustomize", attr.AssetImage("assets/k8s/ecosystem/kustomize.png"), attr.Shape(attr.None))
}

func init() {
  const namespace = "k8s.ecosystem"
  Register(namespace, "Helm", K8sEcosystem.Helm)
  Register(namespace, "Krew", K8sEcosystem.Krew)
  Register(namespace, "Kustomize", K8sEcosystem.Kustomize)
}
