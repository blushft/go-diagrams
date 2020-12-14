package nodes

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type k8sGroupContainer struct {
	path  string
	attrs []attr.Attribute
}

var K8sGroup = &k8sGroupContainer{path: "assets/k8s/group"}

func (c *k8sGroupContainer) Ns(opts ...attr.Attribute) *node.Node {
	return node.New("ns", attr.AssetImage("assets/k8s/group/ns.png"), attr.Shape(attr.None))
}

func init() {
  const namespace = "k8s.group"
  Register(namespace, "Ns", K8sGroup.Ns)
}
