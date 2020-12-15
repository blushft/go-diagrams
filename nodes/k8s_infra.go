package nodes

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type k8sInfraContainer struct {
	path  string
	attrs []attr.Attribute
}

var K8sInfra = &k8sInfraContainer{path: "assets/k8s/infra"}

func (c *k8sInfraContainer) Etcd(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/k8s/infra/etcd.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("etcd", opts...)
}

func (c *k8sInfraContainer) Master(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/k8s/infra/master.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("master", opts...)
}

func (c *k8sInfraContainer) Node(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/k8s/infra/node.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("node", opts...)
}

func init() {
  const namespace = "k8s.infra"
  Register(namespace, "Etcd", K8sInfra.Etcd)
  Register(namespace, "Master", K8sInfra.Master)
  Register(namespace, "Node", K8sInfra.Node)
}
