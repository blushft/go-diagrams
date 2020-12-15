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
	opts = append(opts, attr.AssetImage("assets/k8s/ecosystem/helm.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("helm", opts...)
}

func (c *k8sEcosystemContainer) Krew(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/k8s/ecosystem/krew.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("krew", opts...)
}

func (c *k8sEcosystemContainer) Kustomize(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/k8s/ecosystem/kustomize.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("kustomize", opts...)
}

func init() {
  const namespace = "k8s.ecosystem"
  Register(namespace, "Helm", K8sEcosystem.Helm)
  Register(namespace, "Krew", K8sEcosystem.Krew)
  Register(namespace, "Kustomize", K8sEcosystem.Kustomize)
}
