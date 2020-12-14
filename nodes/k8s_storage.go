package nodes

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type k8sStorageContainer struct {
	path  string
	attrs []attr.Attribute
}

var K8sStorage = &k8sStorageContainer{path: "assets/k8s/storage"}

func (c *k8sStorageContainer) Pv(opts ...attr.Attribute) *node.Node {
	return node.New("pv", attr.AssetImage("assets/k8s/storage/pv.png"), attr.Shape(attr.None))
}

func (c *k8sStorageContainer) Pvc(opts ...attr.Attribute) *node.Node {
	return node.New("pvc", attr.AssetImage("assets/k8s/storage/pvc.png"), attr.Shape(attr.None))
}

func (c *k8sStorageContainer) Sc(opts ...attr.Attribute) *node.Node {
	return node.New("sc", attr.AssetImage("assets/k8s/storage/sc.png"), attr.Shape(attr.None))
}

func (c *k8sStorageContainer) Vol(opts ...attr.Attribute) *node.Node {
	return node.New("vol", attr.AssetImage("assets/k8s/storage/vol.png"), attr.Shape(attr.None))
}

func init() {
  const namespace = "k8s.storage"
  Register(namespace, "Pv", K8sStorage.Pv)
  Register(namespace, "Pvc", K8sStorage.Pvc)
  Register(namespace, "Sc", K8sStorage.Sc)
  Register(namespace, "Vol", K8sStorage.Vol)
}
