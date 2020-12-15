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
	opts = append(opts, attr.AssetImage("assets/k8s/storage/pv.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("pv", opts...)
}

func (c *k8sStorageContainer) Pvc(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/k8s/storage/pvc.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("pvc", opts...)
}

func (c *k8sStorageContainer) Sc(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/k8s/storage/sc.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("sc", opts...)
}

func (c *k8sStorageContainer) Vol(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/k8s/storage/vol.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("vol", opts...)
}

func init() {
  const namespace = "k8s.storage"
  Register(namespace, "Pv", K8sStorage.Pv)
  Register(namespace, "Pvc", K8sStorage.Pvc)
  Register(namespace, "Sc", K8sStorage.Sc)
  Register(namespace, "Vol", K8sStorage.Vol)
}
