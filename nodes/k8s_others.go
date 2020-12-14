package nodes

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type K8sOthersContainer struct {
	path  string
	attrs []attr.Attribute
}

var K8sOthers = &K8sOthersContainer{path: "assets/k8s/others"}

func (c *K8sOthersContainer) Psp(opts ...attr.Attribute) *node.Node {
	return node.New("psp", attr.AssetImage("assets/k8s/others/psp.png"), attr.Shape(attr.None))
}

func (c *K8sOthersContainer) Crd(opts ...attr.Attribute) *node.Node {
	return node.New("crd", attr.AssetImage("assets/k8s/others/crd.png"), attr.Shape(attr.None))
}

func init() {
  const namespace = "k8s.others"
  Register(namespace, "Psp", K8sOthers.Psp)
  Register(namespace, "Crd", K8sOthers.Crd)
}
