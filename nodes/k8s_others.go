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
	opts = append(opts, attr.AssetImage("assets/k8s/others/psp.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("psp", opts...)
}

func (c *K8sOthersContainer) Crd(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/k8s/others/crd.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("crd", opts...)
}

func init() {
  const namespace = "k8s.others"
  Register(namespace, "Psp", K8sOthers.Psp)
  Register(namespace, "Crd", K8sOthers.Crd)
}
